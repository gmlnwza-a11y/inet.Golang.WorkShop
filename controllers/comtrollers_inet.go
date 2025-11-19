package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloTestv2(c *fiber.Ctx) error {
	return c.SendString("Hello, World! V2")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(str)

}

func HelloName(c *fiber.Ctx) error {

	str := "hello ==> " + c.Params("name")
	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

// 5.2
func FactorialTest(c *fiber.Ctx) error {
	numStr := c.Params("num")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid number",
		})
	}
	result := factorial(num)
	return c.JSON(fiber.Map{
		"number":    num,
		"factorial": result,
	})
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 5.3 GetGuys controller to return ASCII values of Tax-ID query parameter
func GetGuys(c *fiber.Ctx) error {
	taxID := c.Query("Tax-ID")

	if taxID == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Tax-ID is required",
		})
	}

	asciiList := []int{}
	for _, ch := range taxID {
		fmt.Println(asciiList)
		asciiList = append(asciiList, int(ch))
	}

	return c.JSON(fiber.Map{
		"Tax-ID":    taxID,
		"ASCIIList": asciiList,
	})
}

// 6
func RegisterUser(c *fiber.Ctx) error {
	var user m.RegisterRequest

	// 1. รับ JSON
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var validate = validator.New()

	// create custom validation for BusinessType
	validate.RegisterValidation("business_allowed", BusinessAllowed)

	validate.RegisterValidation("username_valid", UsernameValid)

	validate.RegisterValidation("subdomain", SubDomain)
	// error handling
	if err := validate.Struct(user); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 3. ถ้าผ่าน → ทำงานต่อ เช่น save DB
	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

func UsernameValid(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	log.Println("Validate ", username)
	// Username must be alphanumeric
	checkUsername, _ := regexp.MatchString("^[A-Za-z0-9]*$", username)

	return checkUsername
}

func BusinessAllowed(fl validator.FieldLevel) bool {
	businessType := fl.Field().String()
	log.Println("Valiate ", businessType)
	allowedTypes := []string{"Retail", "Engineer", "Service", "It"}

	for _, t := range allowedTypes {
		if businessType == t {
			return true
		}
	}
	return false
}

func SubDomain(fl validator.FieldLevel) bool {
	webURL := fl.Field().String()
	log.Println("Validate ", webURL)
	// Subdomain must follow specific pattern
	checkSubdomain, _ := regexp.MatchString("^[a-z0-9-]{2,30}$", webURL)
	return checkSubdomain
}

// work shop dog controllers
func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var deletedDogs []m.Dogs

	// Query for soft deleted records using GORM's Unscoped() method
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&deletedDogs)

	return c.Status(200).JSON(fiber.Map{
		"message": "Deleted dogs retrieved successfully",
		"count":   len(deletedDogs),
		"data":    deletedDogs,
	})
}

func GetDogsRange(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	// dog_id เป็นตัวเลขหรือ string?
	// ถ้า dog_id เป็น INT ใช้แบบนี้ได้เลย:
	result := db.Where("dog_id > ? AND dog_id < ?", 50, 100).Find(&dogs)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	return c.Status(200).JSON(dogs)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //ดึงข้อมูล dogs ทั้งหมด

	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNoColor := 0

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""

		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNoColor++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}

		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:       dataResults,
		Name:       "golang-test",
		Count:      len(dataResults),
		SumRed:     sumRed,
		SumGreen:   sumGreen,
		SumPink:    sumPink,
		SumNoColor: sumNoColor,
	}

	return c.Status(200).JSON(r)
}

// work shop company
