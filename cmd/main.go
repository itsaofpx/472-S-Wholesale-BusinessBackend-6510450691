package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
		&entities.Transaction{}, &entities.Shipment{},
		&entities.Order{}, &entities.Package{},
		&entities.OrderLine{})
	// db.Create(&entities.TierList{Tier: 1, DiscountPercent: 10})
	// db.Create(&entities.TierList{Tier: 2, DiscountPercent: 20})
	// db.Create(&entities.TierList{Tier: 3, DiscountPercent: 30})

	router.SetUpRouters(app, db)

	app.Listen(":8000")
}
