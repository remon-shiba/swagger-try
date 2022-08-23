package main

import (
	"log"
	"net/http"
	"os"
	"swagger-try/model"
	"swagger-try/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type DBRepository struct {
	DB *gorm.DB
}

// UserList godoc
// @Summary      List of users
// @Description  get all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.entities
// @Failure      400  {object}  httputil.HTTPError
// @Router       /get_users [get]
func (dbRepo *DBRepository) GetUsers(c *fiber.Ctx) error {
	usersModel := &[]model.Users{}

	if findErr := dbRepo.DB.Find(usersModel).Error; findErr != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the users"})
		return findErr
	}

	if len(*usersModel) > 0 {
		c.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "record found", "data": usersModel})
	} else {
		c.Status(http.StatusOK).JSON(
			&fiber.Map{"message": "no record found"})
	}
	return nil
}

// API Endpoint
func (dbRepo *DBRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/")
	api.Get("/get_users", dbRepo.GetUsers)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @license.name Apache 2.0
// @host localhost:8000
// @BasePath /
func main() {

	if envErr := godotenv.Load(".env"); envErr != nil {
		log.Fatal(envErr)
	}

	//Database Configurations
	config := &storage.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	db, connErr := storage.NewConnection(config)

	if connErr != nil {
		log.Fatal("could not load the database")
	}
	if migratErr := model.MigrateStruct(db); migratErr != nil {
		log.Fatal("could not migrate the struct")

	}

	dbRepo := DBRepository{
		DB: db,
	}

	app := fiber.New()
	dbRepo.SetupRoutes(app)
	app.Listen(":8000")
}
