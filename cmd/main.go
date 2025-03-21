package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,HEAD,PUT,DELETE"}))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// db, err := gorm.Open(postgres.Open(psqlInfo))

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(
		&entities.User{}, &entities.Product{},
		&entities.Transaction{},
		&entities.Order{},
		&entities.OrderLine{}, &entities.TierList{},
		&entities.OrderLine{}, &entities.Supplier{},
		&entities.SupplierOrderList{},
		&entities.CreditCard{},
		&entities.SupplierOrderList{}, &entities.Chat{},
		&entities.Message{},
	)

	router.SetUpRouters(app, db)

	app.Listen(":8000")
}
