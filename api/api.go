package api

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	srv := &http.Server{
		Addr: ":4000",
	}

	fmt.Println("Starting sever...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
