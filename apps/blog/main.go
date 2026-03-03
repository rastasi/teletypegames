package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"os"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:   os.Getenv("DB_HOST"),
				Port:   os.Getenv("DB_PORT"),
				User:   os.Getenv("DB_USER"),
				Pwd:    os.Getenv("DB_PASSWORD"),
				Name:   os.Getenv("DB_NAME"),
				Driver: config.DriverMysql,
			},
		},
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}

	if err := eng.AddConfig(&cfg).Use(r); err != nil {
		panic(err)
	}

	if err := r.Run(":9033"); err != nil {
		panic(err)
	}
}
