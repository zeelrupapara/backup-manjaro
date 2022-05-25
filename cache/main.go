package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jellydator/ttlcache/v2"
)

type Todo struct {
	Userid    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var cache ttlcache.SimpleCache = ttlcache.NewCache()

func verifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := cache.Get(id)
	if err != ttlcache.ErrNotFound {
		return c.JSON(fiber.Map{"Cached": val})
	}
	return c.Next()
}

func main() {
	app := fiber.New()
	cache.SetTTL(time.Duration(24 * time.Hour))
	app.Get("/:id", verifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + id)
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		todo := Todo{}
		parseErr := json.Unmarshal(body, &todo)
		if parseErr != nil {
			return parseErr
		}

		cache.Set(id, todo)
		return c.JSON(fiber.Map{"Data": todo})
	})

	app.Listen(":3000")
}
