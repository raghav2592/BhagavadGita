package main

import (
	"BhagavadGita/controller"
	"fmt"
	"log"
	"os"
	"time"
)

const PeriodicInterval = 5
var Loglocation = "C:\\BhagavadGita\\app.log"

func main() {

	//Create a log file and store logs
	logfile, err := os.Create(Loglocation)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	// Handle Panic
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic occurred in the code: %s \n", err)
		}
	}()

	//Calling the function after every 5 sec
	count := 1
	ticker := time.NewTicker(PeriodicInterval * time.Second)
	go func(){
		for t := range ticker.C {
			//Call the periodic function here.
			fmt.Printf("starting execution %d at  %v \n", count, time.Now())
			err := controller.GetChapters(t)
			count++
			if err != nil {
				log.Printf("error while getting chapters: %v \n", err.Error())
				os.Exit(1)
			}
		}
	}()

	quit := make(chan bool, 1)
	<-quit
}
