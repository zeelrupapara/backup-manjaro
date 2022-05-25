package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/shareed2k/goth_fiber"
)

// LoadENV
func ConnectENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("ENV FILE NOT LOADED PROPERLY")
	}
}

func main() {
	ConnectENV()
	app := fiber.New()
	goth.UseProviders(
		google.New(os.ExpandEnv("${OAUTH_KEY}"), os.ExpandEnv("${OAUTH_SECRET}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${PORT}/auth/google/callback"), "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/plus.me"),
	)
	app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(user)
	})
	app.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
		if err := goth_fiber.Logout(ctx); err != nil {
			log.Fatal(err)
		}
		return ctx.SendString("SUCCESSFULLY LOGOUT")
	})
	if err := app.Listen(os.ExpandEnv(":${PORT}")); err != nil {
		log.Fatal(err)
	}
}
