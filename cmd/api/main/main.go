package main

import (
	v "github.com/BNIGang/MapLegacy/api/v1/nasabah"
	v1 "github.com/BNIGang/MapLegacy/api/v1/nasabah"
	u "github.com/BNIGang/MapLegacy/api/v1/user"
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/derpen/fastergoding"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var user *web.User
var username string
var secret []byte = login.Secret

func main() {

	fastergoding.Run("./cmd/api/main")

	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/web/", "./web/")

	// Login page
	app.Get("/", func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")
		if cookie != "" {
			return c.Redirect("/home")
		}
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", login.Handler(engine))

	// Handle CRUD for Nasabah
	app.Get("/home", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {

		username = c.Locals("username").(string)

		var err error

		user, err = web.GetUserByUsername(username)
		if err != nil {
			return c.Render("login", fiber.Map{"Error": err})
		}
		if user == nil {
			return c.Render("login", fiber.Map{"Error": err})
		}

		data_nasabah, err := v.GetNasabahDataByUser(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)
		if err != nil {
			return err
		}

		//Placeholder: This is to read afiliasi
		_, err2 := v.GetAfiliasiByUser(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)
		if err2 != nil {
			return err2
		}

		username = user.Name

		return c.Render("template", fiber.Map{
			"Name":         username,
			"Wilayah":      user.Wilayah_ID,
			"Cabang":       user.Cabang_ID,
			"Privilege":    user.User_Privileges,
			"data_nasabah": data_nasabah,
			"content":      "home",
		})
	})

	// Handle page for adding nasabah
	app.Get("/create", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"content":   "create",
		})
	})

	app.Get("/nasabah_detail/:nasabah_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		nasabah_id := c.Params("nasabah_id")

		data_nasabah, err := v.GetNasabahByID(nasabah_id)
		if err != nil {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":         username,
			"Wilayah":      user.Wilayah_ID,
			"Cabang":       user.Cabang_ID,
			"Privilege":    user.User_Privileges,
			"data_nasabah": data_nasabah,
			"content":      "nasabah_detail",
		})
	})

	app.Get("/edit/:nasabah_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		nasabah_id := c.Params("nasabah_id")

		data_nasabah, err := v.GetNasabahByID(nasabah_id)
		if err != nil {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":         username,
			"Wilayah":      user.Wilayah_ID,
			"Cabang":       user.Cabang_ID,
			"Privilege":    user.User_Privileges,
			"data_nasabah": data_nasabah,
			"content":      "edit",
		})
	})

	app.Get("/home/", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return c.Redirect("/home")
	})

	// Handle every endpoint to allow dynamic form
	app.Get("/get_bidang_usaha", web.JWTMiddleware(secret, engine), web.GetBidangUsahaHandler)
	app.Get("/get_produk_usaha/:bidang_id", web.JWTMiddleware(secret, engine), web.GetProdukUsahaHandler)

	app.Get("/get_cabang", web.JWTMiddleware(secret, engine), web.GetCabangHandler)
	app.Get("/get_kota_kabupaten/:cabang_id", web.JWTMiddleware(secret, engine), web.GetKotaKabupatenHandler)
	app.Get("/get_kcu_kcp_kk/:cabang_id", web.JWTMiddleware(secret, engine), web.GetKCPKCUKKHandler)
	// Handling dynamic column over

	// Add nasabah
	app.Post("/add", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return v.AddNasabahHandler(user.User_ID)(c)
	})

	app.Post("/add_afiliasi", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return v1.AddAfiliasi(user.User_ID)(c)
	})

	app.Get("/create_map_legacy/:nasabah_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		nasabah_id := c.Params("nasabah_id")
		data_nasabah, err := v.GetAfiliasiListById(nasabah_id)
		if err != nil {
			data_nasabah = &v.MergedRow{}
		}

		if len(data_nasabah.MergedAfiliasi) == 0 {
			data_nasabah2, err := v.GetNasabahByID(nasabah_id)
			if err != nil {
				return err
			}
			nama_pengusaha := data_nasabah2.Nama_pengusaha
			data_nasabah.NamaPengusaha = nama_pengusaha
		}

		afiliasiList, err := v.MapLegacyHandler(data_nasabah)
		if err != nil {
			return err
		}

		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":          username,
			"Wilayah":       user.Wilayah_ID,
			"Cabang":        user.Cabang_ID,
			"Privilege":     user.User_Privileges,
			"data":          data_nasabah,
			"afiliasi_list": afiliasiList,
			"content":       "map_legacy",
		})
	})

	// Delete nasabah
	// TODO: add confirmation before deleting
	app.Post("/delete/:nasabah_id", web.JWTMiddleware(secret, engine), v.DeleteNasabahData)

	// Update nasabah
	app.Post("/update/:nasabah_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return v.UpdateNasabahData(user.User_ID)(c)
	})

	app.Get("/search_nasabah/:query", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		// Check if user is nil or username is empty
		if user == nil || user.Username == "" {
			return c.Redirect("/home")
		}

		// Call SearchNasabah function with user information
		return v.SearchNasabah(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)(c)
	})

	// Now, do CRUD for afiliasi
	app.Get("/afiliasi", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		data_afiliasi, err := v.GetAfiliasiByUser(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)
		if err != nil {
			return err
		}

		return c.Render("template", fiber.Map{
			"Name":          username,
			"Wilayah":       user.Wilayah_ID,
			"Cabang":        user.Cabang_ID,
			"Privilege":     user.User_Privileges,
			"data_afiliasi": data_afiliasi,
			"content":       "afiliasi",
		})
	})

	app.Get("/create_afiliasi", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"content":   "create_afiliasi",
		})
	})

	// Delete nasabah
	// TODO: add confirmation before deleting
	app.Post("/delete_afiliasi/:afiliasi_id", web.JWTMiddleware(secret, engine), v.DeleteAfiliasiData)

	app.Get("/edit_afiliasi/:id_child", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		id_child := c.Params("id_child")

		data_afiliasi, err := v.GetAfiliasiById(id_child)
		if err != nil {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":          username,
			"Wilayah":       user.Wilayah_ID,
			"Cabang":        user.Cabang_ID,
			"Privilege":     user.User_Privileges,
			"data_afiliasi": data_afiliasi,
			"content":       "edit_afiliasi",
		})
	})

	app.Post("/update_afiliasi/:afiliasi_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return v.UpdateAfiliasi(user.User_ID)(c)
	})

	app.Get("/search_afiliasi/:query", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		// Check if user is nil or username is empty
		if user == nil || user.Username == "" {
			return c.Redirect("/home")
		}

		// Call SearchNasabah function with user information
		return v.SearchAfiliasi(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)(c)
	})

	// handle autofill
	app.Get("/get_suggestions/:nama_pengusaha", web.JWTMiddleware(secret, engine), web.AutoFillHandler)

	app.Get("/edit_password", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		alert := c.Query("alert") // Get the value of the "alert" query parameter

		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"Id":        user.User_ID,
			"Alert":     alert,
			"content":   "edit_password",
		})
	})

	app.Post("/edit_pass", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return u.EditPassword()(c)
	})

	app.Get("/user_page", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {

		privilege := c.Locals("privilege").(string)

		if privilege != "admin" {
			return c.Redirect("/home")
		}

		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		userslist, err := u.GetUsers()
		if err != nil {
			return nil
		}

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"UserList":  userslist,
			"content":   "user_page",
		})
	})

	app.Post("/add_users", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		privilege := c.Locals("privilege").(string)

		if privilege != "admin" {
			return c.Redirect("/home")
		}
		return u.AddUsersHandler()(c)
	})

	// Handle User account
	app.Get("/add_users", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {

		privilege := c.Locals("privilege").(string)
		alert := c.Query("alert") // Get the value of the "alert" query parameter

		if privilege != "admin" {
			return c.Redirect("/home")
		}

		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"Alert":     alert,
			"content":   "add_users",
		})
	})

	app.Post("/delete_user/:user_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		privilege := c.Locals("privilege").(string)

		if privilege != "admin" {
			return c.Redirect("/home")
		}
		return u.DeleteUser()(c)
	})

	app.Get("/edit_user/:user_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		privilege := c.Locals("privilege").(string)

		if privilege != "admin" {
			return c.Redirect("/home")
		}

		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		user_id := c.Params("user_id")

		data_user, err := u.GetUserByID(user_id)
		if err != nil {
			return err
		}

		alert := c.Query("alert") // Get the value of the "alert" query parameter

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"DataUser":  data_user,
			"Alert":     alert,
			"content":   "edit_user",
		})
	})

	app.Post("/update_user/:user_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		privilege := c.Locals("privilege").(string)

		if privilege != "admin" {
			return c.Redirect("/home")
		}

		user_id := c.Params("user_id")

		return u.UpdateUser(user_id)(c)
	})

	app.Get("/logout", login.LogoutHandler)

	port := ":8000"
	app.Listen(port)
}
