package delivery

import (
	"log"
	"sahih-be/internal/database"
	"sahih-be/internal/entity"

	"github.com/gofiber/fiber/v2"
)

func GetAllDoa(c *fiber.Ctx) error {
	userInfo := c.Locals("userInfo")
	log.Println(userInfo)

	var doaCollection []entity.DoaCollection
	err := database.DB.Find(&doaCollection).Error
	if err != nil {
		log.Println(err)
	}
	return c.Status(200).JSON(fiber.Map{
		"message":       "success",
		"doaCollection": doaCollection,
	})
}
