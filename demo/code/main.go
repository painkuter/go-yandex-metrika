package main

import (
	"log"
	"os"

	metrika "go-yandex-metrika"
)

func main() {

	log.Println("Start")
	code := os.Getenv("YANDEX_CODE")
	if code == "" {
		log.Fatal("empty yandex code")
	}
	clientID := os.Getenv("YANDEX_CLIENT_ID")
	if clientID == "" {
		log.Fatal("empty yandex client id")
	}
	clientSecret := os.Getenv("YANDEX_CLIENT_SECRET")
	if clientID == "" {
		log.Fatal("empty yandex client secret")
	}

	metrika := metrika.NewMetrikaFromCode(code, clientID, clientSecret)
	metrika.SetDebug(true)

	if err := metrika.Authorize(); err != nil {
		log.Fatal(err.Error())
	}

	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
