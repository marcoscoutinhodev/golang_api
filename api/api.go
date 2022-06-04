package api

import (
	"fmt"
	"log"
	"net/http"

	database "golang_api/api/database/sql"
)

func Server() {
	srv := &http.Server{
		Addr:    ":4000",
		Handler: Routes(),
	}

	databaseHelperDev()
	fmt.Println("tbTask was created in mysql instance")

	fmt.Println("Starting sever...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func databaseHelperDev() {
	db, err := database.Connection()

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("MYSQL connection error")
	}

	defer db.Close()

	db.Exec("CREATE DATABASE IF NOT EXISTS golang_db")
	db.Exec("USE golang_db")
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tbTask (
		id 					INT AUTO_INCREMENT,
		title				VARCHAR(40) NOT NULL,
		description VARCHAR(150) NOT NULL,
		done 				TINYINT NOT NULL,

		PRIMARY KEY (id)
	)`)

	if err != nil {
		log.Fatal(err.Error())
	}
}
