package main

import (
	"fmt"
	"go-api-server/wire"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func initDB() *sqlx.DB {
	err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}

	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	dns := fmt.Sprintf(user + ":" + pass + "@tcp(127.0.0.1:3306)/" + name)
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	e := echo.New()
	db := initDB()
	userAPI := wire.InitUserAPI(db)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"ping": "pong"})
	})
	e.GET("/data", userAPI.CreateUser())
	e.GET("/get", userAPI.GetAllUser())
	e.Logger.Fatal(e.Start(":3000"))
}
