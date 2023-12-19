package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMiddleware(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("gormDB", db)
		return ctx.Next()
	}
}

func InitGormDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Test Database Connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("successfully create connection to GORM DB")
	return db, nil
}
