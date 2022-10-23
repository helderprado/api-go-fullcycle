package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-api/application/usecase"
	"go-api/infra/kafka"
	"go-api/infra/repository"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/fullcycle")
	if err != nil {
		log.Fatalln(err)
	}
	repository := repository.CourseMySQLRepository{Db: db}
	usecase := usecase.CreateCourse{Repository: repository}

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
	}
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var input usecase.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)
		output, err := usecase.Execute(input)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println(output)
		}
	}

}
