package compA

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(api fiber.Router) fiber.Router {
	router := api.Group("/compA")

	router.Get("/:id", GetByIdController)
	router.Get("/:x/:y", SumController)

	return router
}

func GetByIdController(c *fiber.Ctx) error {
	// id, err := c.ParamsInt("id")
	// if err != nil {
	// 	return c.SendStatus(400)
	// }
	// return c.JSON(GetById(id))

	return c.SendString("Hello World")
}

func SumController(c *fiber.Ctx) error {
	x, xErr := c.ParamsInt("x")
	y, yErr := c.ParamsInt("y")
	if xErr != nil || yErr != nil {
		return c.SendStatus(400)
	}

	return c.JSON(fiber.Map{
		"result": Sum(x, y),
	})
}
