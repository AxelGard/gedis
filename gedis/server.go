package main

import "github.com/gofiber/fiber/v2"

// db should be map of string to interface{} so that we can handle both having a string as a value and array of strings as a value
var db = map[string]interface{}{
	"foo": "bar",
	"bar": []string{"baz", "qux"},
}

func main() {
	app := fiber.New()

	app.Get("/", get_all)

	app.Get("/:key", get_value)
	app.Post("/:key", set_value_body)
	app.Post("/:key/:value", set_value)

	app.Listen(":3000")
}

func get_all(c *fiber.Ctx) error {
	return c.JSON(db)
}

func get_value(c *fiber.Ctx) error {
	key := c.Params("key")
	value := db[key]
	return c.JSON(value)

}

func set_value(c *fiber.Ctx) error {
	key := c.Params("key")
	value := c.Params("value")
	db[key] = value
	return c.JSON(db[key])
}

func set_value_body(c *fiber.Ctx) error {
	key := c.Params("key")
	value := c.Body()
	db[key] = value
	return c.JSON(db[key])
}
