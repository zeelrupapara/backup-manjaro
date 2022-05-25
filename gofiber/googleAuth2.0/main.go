// For Google Authentication:
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	gf "github.com/shareed2k/goth_fiber"
)

func main() {
	goth.UseProviders(
		google.New(os.ExpandEnv("OAUTH_KEY"), os.ExpandEnv("OUTH_SECRET"), "http://localhost:3000/auth/google/callback"),
	)

	app := fiber.New()

	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		ctx.JSON(user)
		return nil
	})

	app.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		ctx.Redirect("http://localhost:41159/")
		return nil
	})

	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
			ctx.JSON(gothUser)
		} else {
			gf.BeginAuthHandler(ctx)
		}
		return nil
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Redirect("/auth/google/")
		return nil
	})

	log.Fatal(app.Listen(":3000"))

}
