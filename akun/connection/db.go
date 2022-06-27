package connection

import (
	"log"
	"os"
	"perpustakaan/akun/models"

	_ "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error Load Config Files !")
	}
}

func ConnectToDb() *gorm.DB {
	url := os.Getenv("URL_DB")
	port := os.Getenv("PORT_DB")
	DB := os.Getenv("DB_NAME")
	user := os.Getenv("USERNAME_DB")
	pw := os.Getenv("PW_DB")
	dsn := user + ":" + pw + "@tcp(" + url + ":" + port + ")/" + DB

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Pegawai{})

	return db
}
