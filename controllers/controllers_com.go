package controllers

import (
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"

	"github.com/gofiber/fiber/v2"
)

func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Company
	db.Find(&companies)

	return c.Status(200).JSON(companies)
}

// this is company controller
func AddCompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var company []m.Company

	result := db.Find(&company, "id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&company)
}

func GetCompaniesJson(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Company

	db.Find(&companies)

	// Create CompanyRes struct for response

	var dataResults []m.CompanyRes
	for _, v := range companies {
		typeStr := ""
		// Determine type based on employee count
		if v.Employee <= 50 {
			typeStr = "small"
		} else if v.Employee <= 200 {
			typeStr = "medium"
		} else {
			typeStr = "large"
		}

		d := m.CompanyRes{
			Name:        v.Name,
			CEO:         v.CEO,
			Employee:    v.Employee,
			Location:    v.Location,
			Established: v.Established,
			Type:        typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.CompanyResultData{
		Data:  dataResults,
		Name:  "company-api",
		Count: len(companies),
	}
	return c.Status(200).JSON(r)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	db.Delete(&company, id)

	return c.SendStatus(200)
}
