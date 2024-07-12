package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	postgres "vortex/internal/db/postgre"
	"vortex/internal/handlers"
)

func main() {
	//db := postgres.createConnection()
	defer postgres.Db.Close()
	postgres.DropAll()
	postgres.GetTables()
	postgres.SetUpDB()
	postgres.GetTables()

	//postgres.GetTables()

	//log.Println()
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		//log.Println("Crone doing smth")
	})

	c.Start()

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/AddClient", handlers.AddClient)
	router.HandleFunc("/UpdateClient", handlers.UpdateClient)
	router.HandleFunc("/DeleteClient", handlers.DeleteClient)
	router.HandleFunc("/UpdateAlgorithmStatus", handlers.UpdateAlgorithmStatus)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
