package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	requestId := make(chan int)

	concurrency := 2

	for i := 1; i <= concurrency; i++ {
		go worker(requestId, i)
	}

	for i := 0; i < 20; i++ {
		requestId <- i
	}
}

func worker(requestId chan int, worker int) {
	for r := range requestId {
		//res, err := http.Get("http://localhost:8585/product")
		//if err != nil {
		//	log.Fatal("Erro")
		//}
		//defer res.Body.Close()
		//content, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("Worker %d. RequestId: %d. ", worker, r)
		r := rand.Intn(2)
		time.Sleep((time.Duration(r * int(time.Second))))
	}
}
