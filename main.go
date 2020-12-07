package main

import (
	"log"
	"net/http"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func initDB() *sqlx.DB {
	// db, err := sqlx.Open("mysql", "root@/practice")
	db, err := sqlx.Connect("mysql", "root@tcp(127.0.0.1:3306)/practice")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("USE practice")
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
	ID    int `db:"id" json:"money"`
	Money int `db:"money" json:"money"`
}

func main() {
	e := echo.New()
	db := initDB()
	// db.MustExec(schema)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	e.GET("/data", func(c echo.Context) error {
		tx := db.MustBegin()
		tx.NamedExec("INSERT INTO login_log (money) VALUES(:money)", LogData{0, 100})
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
