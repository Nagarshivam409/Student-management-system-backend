package repository

import (
	"backend/pkg/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func GetDB() *gorm.DB {
	return db
}
func Intialmigration() {
	err = godotenv.Load(".env")
	// Loading Environment Variables
	// dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//Database Connection String
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, dbPort)
	// DataBase connection
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	//defer db.Close()
	if err != nil {
		fmt.Println("Error Connecting to dataBase")
		panic(err)

	} else {
		fmt.Println("Succesfully Connected to database")
	}
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.Teacher{})
	db.AutoMigrate(&models.AttendanceStudent{})
	db.AutoMigrate(&models.AttendanceTeacher{})
}
