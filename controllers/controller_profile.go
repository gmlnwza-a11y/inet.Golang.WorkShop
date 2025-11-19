package controllers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUserProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var userprofile []m.UserProfile
	db.Find(&userprofile)

	return c.Status(200).JSON(userprofile)
}

// this is company controller
func AddUserProfile(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var userprofile m.UserProfile

	if err := c.BodyParser(&userprofile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&userprofile)
	return c.Status(201).JSON(userprofile)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var userprofile m.UserProfile
	id := c.Params("id")

	if err := c.BodyParser(&userprofile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&userprofile)
	return c.Status(200).JSON(userprofile)
}

func RemoveUserProfile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var userprofile m.UserProfile

	db.Delete(&userprofile, id)

	return c.SendStatus(200)
}

func GetUserProfileJson(c *fiber.Ctx) error {
	db := database.DBConn
	var userprofiles []m.UserProfile

	db.Find(&userprofiles) //ดึงข้อมูล userprofiles ทั้งหมด

	GenZ := 0
	GenY := 0
	GenX := 0
	BabyBoomer := 0
	GIGeneration := 0

	var dataResults []m.UserProfileRes
	for _, a := range userprofiles {
		typeStr := ""

		if a.Age < 24 {
			typeStr = "GenZ"
			GenZ++
		} else if a.Age >= 24 && a.Age <= 41 {
			typeStr = "GenY"
			GenY++
		} else if a.Age >= 42 && a.Age <= 56 {
			typeStr = "GenX"
			GenX++
		} else if a.Age >= 57 && a.Age <= 75 {
			typeStr = "BabyBoomer"
			BabyBoomer++
		} else {
			typeStr = "no color"
			GIGeneration++
		}

		t := m.UserProfileRes{
			Name: a.Name,
			Age:  a.Age,
			Type: typeStr,
		}

		dataResults = append(dataResults, t)
	}

	r := m.UserProfileResultData{
		Data:         dataResults,
		Name:         "golang-test",
		Count:        len(dataResults),
		GenZ:         GenZ,
		GenY:         GenY,
		GenX:         GenX,
		BabyBoomer:   BabyBoomer,
		GIGeneration: GIGeneration,
	}

	return c.Status(200).JSON(r)
}

func GetUserProfileSearch(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var UserProfile []m.UserProfile

	result := db.Where(
		"employee_id = ? OR name LIKE ? OR last_name LIKE ?",
		search, "%"+search+"%", "%"+search+"%").Find(&UserProfile)
	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&UserProfile)
}
