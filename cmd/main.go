package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		log.Println("Crone doing smth")
	})

	c.Start()

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
