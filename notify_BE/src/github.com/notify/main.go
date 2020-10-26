package main

import (
	"github.com/joho/godotenv"
	"github.com/notify/disperser/db"
	routes "github.com/notify/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func server() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("port")

	// // Handle routes
	http.Handle("/", routes.Handlers())

	// // serve
	log.Println("INFO: [SV]: Server is Online @Port:" + port)
	log.Fatal(http.ListenAndServeTLS(":"+port,"./cert/localhost.crt","./cert/localhost.key", nil))
}

func requestDisperser(jobQueue *db.JobQueue) {

	log.Println("INFO: [RD]: Job Queue Is Activated")
	for {
		jobQueue.Disperse()
	}
}

func main() {
	log.Println("[NOTIFY]: Notification application.")

	// create job queue for run job
	jobQueue := db.GetJobsQueue()

	// go routing
	go server()
	go requestDisperser(jobQueue)

	// exit function
	exit()
}

func exit() {
	var endWaiter sync.WaitGroup
	endWaiter.Add(1)
	var signalChannel chan os.Signal
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		<-signalChannel
		log.Println("INFO: Exiting")
		endWaiter.Done()
	}()
	endWaiter.Wait()
}