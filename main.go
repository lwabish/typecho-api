package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lwabish/typecho-api/handlers/content"
	"github.com/lwabish/typecho-api/utils"
	"github.com/lwabish/typecho-api/utils/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	dbType   string
	host     string
	port     int
	user     string
	password string
	dbname   string
)

func init() {
	flag.StringVar(&dbType, "dbType", "mysql", "database type")
	flag.StringVar(&host, "host", "localhost", "database host")
	flag.IntVar(&port, "port", 3307, "database port")
	flag.StringVar(&user, "user", "root", "database user")
	flag.StringVar(&password, "password", "root", "database password")
	flag.StringVar(&dbname, "dbname", "typecho", "database name")
}

func main() {
	flag.Parse()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port, dbname),
	}), &gorm.Config{})
	utils.Must(err)

	l := log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)

	content.Hdl.Setup(db, l)

	r := gin.Default()
	routes.RegisterRoutes(r)
	utils.Must(r.Run(":8080"))
}
