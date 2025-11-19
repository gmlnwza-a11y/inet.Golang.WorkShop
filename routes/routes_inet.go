package routes

import (
	c "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {

	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
	}))

	// /api/v1/
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v3 := api.Group("/v3")

	v1.Get("/", c.HelloTest)
	v2.Get("/", c.HelloTestv2)
	v1.Post("/", c.BodyParserTest)
	v1.Get("/user/:name", c.HelloName)
	v1.Post("/inet", c.QueryTest)
	v1.Post("/valid", c.ValidTest)
	v1.Get("/fact/:num", c.FactorialTest)
	v1.Post("/register", c.RegisterUser)
	v3.Get("/guy", c.GetGuys)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", c.GetDogs)
	dog.Get("/filter", c.GetDog)
	dog.Get("/json", c.GetDogsJson)
	dog.Get("/deleted", c.GetDeletedDogs)
	dog.Get("/range", c.GetDogsRange)
	dog.Post("/", c.AddDog)
	dog.Put("/:id", c.UpdateDog)
	dog.Delete("/:id", c.RemoveDog)

	//CRUD company
	company := v1.Group("/company")
	company.Get("", c.GetCompanies)
	company.Get("/filter", c.GetCompany)
	company.Get("/json", c.GetCompaniesJson)
	company.Post("/", c.AddCompany)
	company.Put("/:id", c.UpdateCompany)
	company.Delete("/:id", c.RemoveCompany)
}
