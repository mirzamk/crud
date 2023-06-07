package main

import (
	"crud/modules/user"
	"crud/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.New()

	//open connection db
	dbCrud := db.GormMysql()

	//check connection
	checkdb, err := dbCrud.DB()
	if err != nil {
		log.Fatal(err)
	}

	//ping to database
	errconn := checkdb.Ping()
	if err != nil {
		log.Fatal(errconn)
	}

	fmt.Println(("database connected"))

	handleruser := user.NewRouter(dbCrud)
	handleruser.Handle(router)

	errRouter := router.Run()
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}
}
