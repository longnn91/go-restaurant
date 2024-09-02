package main

import (
	authModal "gogo/modules/auth/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	//Connect to database
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	//Migrate the schema

	db.AutoMigrate(&authModal.Users{})
	// db.Migrator().CreateTable(&authModal.Users{})
	// db.Migrator().DropTable(&authModal.Users{})

	//Config API use gin
	r := SetupRouter(db)

	server_err := r.Run(":" + os.Getenv("PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if server_err != nil {
		log.Fatal("Error loading .env file", server_err)
		panic(server_err)
	}
}
