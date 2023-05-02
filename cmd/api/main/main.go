package main

import (
	v1 "github.com/BNIGang/MapLegacy/api/v1/nasabah"
	"github.com/BNIGang/MapLegacy/login"
	"github.com/BNIGang/MapLegacy/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var user *web.User
var username string
var secret []byte = login.Secret

func main() {

	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Login page
	app.Get("/", func(c *fiber.Ctx) error {
		cookie := c.Cookies("token")
		if cookie != "" {
			return c.Redirect("/home")
		}
		return c.Render("login", fiber.Map{})
	})

	app.Post("/login", login.Handler(engine))

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

	// Delete nasabah
	app.Post("/delete/:nasabah_id", web.JWTMiddleware(secret, engine), v1.DeleteNasabahData)

	// Update nasabah
	app.Post("/update/:nasabah_id", web.JWTMiddleware(secret, engine), v1.UpdateNasabahData)

	app.Get("/logout", login.LogoutHandler)

	port := ":8000"
	app.Listen(port)
}
