package main

import (
	"log"
	"os"

	metrika "go-yandex-metrika"
)

func main() {

	log.Println("Start")

	token := os.Getenv("YANDEX_TOKEN")
	if token == "" {
		log.Fatal("empty yandex token")
	}

	metrika := metrika.NewMetrikaFromToken(token)
	metrika.SetDebug(true)

	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
