package database

import (
	"log"
	"os"

	"com.wangzhumo.distribute/database/mysql"
	"gorm.io/gen"
)

func InitDB() {
	db := mysql.InstanceDB
	if db != nil {
		log.Println("mysql create succeed")
	}
}

func InitDir() {
	_, err := os.Stat(".data")
	if os.IsNotExist(err) {
		os.MkdirAll(".data", 0755)
	}

	_, err = os.Stat(".data/icons")
	if os.IsNotExist(err) {
		os.MkdirAll(".data/icons", 0755)
	}

	_, err = os.Stat(".data/apks")
	if os.IsNotExist(err) {
		os.MkdirAll(".data/apks", 0755)
	}
}

func GenerateModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "models/dal",
		Mode:    gen.WithDefaultQuery,
	})
	g.UseDB(mysql.InstanceDB())

	g.ApplyBasic(g.GenerateModel("apkinfo"), g.GenerateModel("versions"))

	g.Execute()
}
