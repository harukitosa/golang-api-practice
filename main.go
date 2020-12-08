package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

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

const schema = `
CREATE TABLE login_log(
	id INT AUTO_INCREMENT,
	money INT,
	PRIMARY KEY (id)
);
`

// LogData is db data
type LogData struct {
	ID    int `db:"id" json:"id"`
	Money int `db:"money" json:"money"`
}

func main() {
	e := echo.New()
	db := initDB()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	e.GET("/data", func(c echo.Context) error {
		tx := db.MustBegin()
		tx.NamedExec("INSERT INTO login_log (money) VALUES(:money)", LogData{0, 130})
		tx.Commit()
		return c.JSON(http.StatusOK, map[string]string{"hello": "ok"})
	})

	e.GET("/all", func(c echo.Context) error {
		data := []LogData{}
		db.Select(&data, "SELECT * FROM login_log ORDER BY money ASC")
		_, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		log.Println(data)
		return c.JSON(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":1313"))
}
