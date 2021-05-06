package driver

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user     = "muthiaaarr"
	dbname   = "be-sharing-vision"
	password = "samsung"
	host     = "localhost"
	port     = 3306
)

func ConnectToDB() (*sql.DB, error) {
	sqlInfo := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`,
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	fmt.Println("Connect success!")
	return db, nil
}
