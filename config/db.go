package config

import (
	"os"

	"github.com/alifwildanaz/FP2-MSIB5-Hacktiv8/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host		= os.Getenv("PGHOST")
	user		= os.Getenv("PGUSER")
	password	= os.Getenv("PGPASSWORD")
	port		= os.Getenv("PGPORT")
	dbname		= os.Getenv("PGDATABASE")
)

func InitDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}
