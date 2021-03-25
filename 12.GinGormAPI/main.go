package main

import (
	"github.com/Lexa-san/spc-go2/12.GinGormAPI/models"
	"github.com/Lexa-san/spc-go2/12.GinGormAPI/routers"
	"github.com/Lexa-san/spc-go2/12.GinGormAPI/storage"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var err error

func main() {
	storage.DB, err = gorm.Open("postgres",
		"host=localhost port=5432 sslmode=disable dbname=alexeysolovev user=postgres password=postgres",
	)
	if err != nil {
		log.Println("error while accessing database:", err)
	}
	defer storage.DB.Close()

	storage.DB.AutoMigrate(&models.Article{})

	//r - gin router
	r := routers.SetupRouter()
	r.Run()
}
