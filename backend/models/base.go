package models

import (
	"fmt"

	//blank import
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	//blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := "root"  //os.Getenv("db_user")
	password := ""      //os.Getenv("db_pass")
	dbName := "meetups" // os.Getenv("db_name")
	// dbHost := os.Getenv("db_host")
	dbType := "mysql"
	dbPort := "3306" //os.Getenv("db_port")

	dbURI := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?parseTime=true", username, password, dbPort, dbName) //Build connection string
	fmt.Println(dbURI)

	conn, err := gorm.Open(dbType, dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Meetup{})

}

// GetDB returns db instance.
func GetDB() *gorm.DB {
	return db
}
