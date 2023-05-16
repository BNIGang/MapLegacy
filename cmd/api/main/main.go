package main

import (
	v1 "github.com/BNIGang/MapLegacy/api/v1/nasabah"
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

	//TODO
	//Probably Remove this Later
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

		data_nasabah, err := v1.GetNasabahDataByUser(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)
		if err != nil {
			return err
		}

		return c.Render("template", fiber.Map{
			"Name":         username,
			"Wilayah":      user.Wilayah_ID,
			"Cabang":       user.Cabang_ID,
			"Privilege":    user.User_Privileges,
			"data_nasabah": data_nasabah,
			"content":      "home",
		})
	})

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

	app.Get("/edit/:nasabah_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		nasabah_id := c.Params("nasabah_id")

		data_nasabah, err := v1.GetNasabahByID(nasabah_id)
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
		return v1.AddNasabahHandler(user.User_ID)(c)
	})

	app.Get("/create_map_legacy/:nasabah_id", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		nasabah_id := c.Params("nasabah_id")
		data_nasabah, err := v1.GetNasabahByID(nasabah_id)
		// afiliasiList, err := v1.MapLegacyHandler(nasabah_id)

		if err != nil {
			// Handle the error appropriately
			return err
		}

		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		counter := 1

		return c.Render("template", fiber.Map{
			"Name":      username,
			"Wilayah":   user.Wilayah_ID,
			"Cabang":    user.Cabang_ID,
			"Privilege": user.User_Privileges,
			"data":      data_nasabah,
			"counter":   counter,
			"content":   "map_legacy",
		})
	})

	// Delete nasabah
	// TODO: add confirmation before deleting
	app.Post("/delete/:nasabah_id", web.JWTMiddleware(secret, engine), v1.DeleteNasabahData)

	// Update nasabah
	app.Post("/update/:nasabah_id", web.JWTMiddleware(secret, engine), v1.UpdateNasabahData)

	// Now, do CRUD for afiliasi
	app.Get("/afiliasi", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		if user == nil || username == "" {
			return c.Redirect("/home")
		}

		data_afiliasi, err := v1.GetAfiliasiByUser(user.User_ID, user.Wilayah_ID, user.Cabang_ID, user.User_Privileges)
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
	app.Post("/delete_afiliasi/:afiliasi_id", web.JWTMiddleware(secret, engine), v1.DeleteAfiliasiData)

	app.Get("/edit_afiliasi/:id_child", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		id_child := c.Params("id_child")

		data_afiliasi, err := v1.GetAfiliasiById(id_child)
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

	// handle autofill
	app.Get("/get_suggestions/:nama_pengusaha", web.JWTMiddleware(secret, engine), web.AutoFillHandler)

	app.Get("/logout", login.LogoutHandler)

	app.Post("/add_afiliasi", web.JWTMiddleware(secret, engine), func(c *fiber.Ctx) error {
		return v1.AddAfiliasi(user.User_ID)(c)
	})

	port := ":8000"
	app.Listen(port)
}
