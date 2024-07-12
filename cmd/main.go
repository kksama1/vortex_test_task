package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"vortex/config"
	postgres "vortex/internal/db/postgre"
	"vortex/internal/handlers"
	"vortex/internal/pod_placeholder"
)

func main() {

	cfg, err := config.LoadConfig[config.DatabaseConfig]()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
	pool := postgres.CreateConnection(cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password, cfg.Sslmode)
	service := handlers.NewService(pool)
	defer func() {
		err = service.DB.CloseConnection()
		if err != nil {
			log.Println(err)
		}
	}()

	service.DB.SetUpDB()
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
