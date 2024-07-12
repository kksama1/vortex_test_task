package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
	"strconv"
	"vortex/internal/handlers"
	"vortex/internal/pod_placeholder"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DATABASE_HOST")
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic(err)
	}
	database := os.Getenv("DATABASE_NAME")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")

	service := handlers.Service{}
	service.DB.CreateConnection(host, port, database, username, password)
	defer func() {
		err := service.DB.CloseConnection()
		if err != nil {
			log.Println(err)
		}
	}()
	//service.DB.DropAll()
	//service.DB.GetTables()
	service.DB.SetUpDB()
	service.DB.GetTables()

	podList := pod_placeholder.PodList{Pods: make([]pod_placeholder.PodPlaceholder, 0)}

	c := cron.New()
	_, err = c.AddFunc("@every 30s", func() {
		activeAlgorithms, err := service.DB.GetActiveAlgorithms()
		if err != nil {
			log.Println(err)
			return
		}
		inActiveAlgorithms, err := service.DB.GetInActiveAlgorithms()
		if err != nil {
			log.Println(err)
			return
		}

		activePods, err := podList.GetPodList()
		if err != nil {
			log.Println(err)
		}
		for i := range activeAlgorithms {
			var exists = false
			podName := fmt.Sprintf("Pod for algorithm №%d", activeAlgorithms[i].ID)
			for j := range activePods {
				if podName == activePods[j] {
					exists = true
				}
			}
			if !exists {
				err = podList.CreatePod(podName)
				if err != nil {
					log.Println(err)
				}
			}
		}

		for i := range inActiveAlgorithms {
			if len(activePods) == 0 {
				break
			}
			podName := fmt.Sprintf("Pod for algorithm №%d", inActiveAlgorithms[i].ID)
			for j := range activePods {
				if podName == activePods[j] {
					err = podList.DeletePod(podName)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}

		activePods, err = podList.GetPodList()
		if err != nil {
			log.Println(err)
		}
		log.Println("Active pods: ", activePods)
	})
	if err != nil {
		panic(fmt.Errorf("creation of cron func failed: %v", err))
	}

	c.Start()

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/AddClient", service.AddClient)
	router.HandleFunc("/UpdateClient", service.UpdateClient)
	router.HandleFunc("/DeleteClient", service.DeleteClient)
	router.HandleFunc("/UpdateAlgorithmStatus", service.UpdateAlgorithmStatus)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
