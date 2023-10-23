package main

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-3/config"
	"homework-3/internal/infrastructure/kafka"
	"homework-3/internal/pkg/app"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/models"
	"homework-3/internal/pkg/repository/postgresql"
	"homework-3/internal/pkg/routers"
	"log"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
)

const port = ":9000"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := config.NewConfig()
	database, err := db.NewDB(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	authorRepo := postgresql.NewAuthors(database)
	bookRepo := postgresql.NewBooks(database)
	brokers := []string{"localhost:9091"}
	producer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	sender := kafka.NewKafkaSender(producer, "library")

	a := app.NewApp(authorRepo, bookRepo, sender)

	go consumerStart(brokers, a)

	router := mux.NewRouter()
	routers.CreateAuthorRouter(router, a)
	routers.CreateBookSubRouter(router, a)
	http.Handle("/", router)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func consumerStart(brokers []string, a *app.App) {
	topic := "library"
	kafkaConsumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	// обработчики по каждому из топиков
	handlers := map[string]kafka.HandleFunc{
		topic: func(message *sarama.ConsumerMessage) {
			mes := models.HandlerMessage{}
			err = json.Unmarshal(message.Value, &mes)
			if err != nil {
				fmt.Println("Consumer error", err)
			}

			fmt.Println("Received Key: ", string(message.Key), " Value: ", mes)
		},
	}

	a.MessageReceiver = kafka.NewReceiver(kafkaConsumer, handlers)

	// При условии одного инстанса подходит идеально
	// payments.StartConsume("payments")
	err = a.MessageReceiver.Subscribe(topic)
	if err != nil {
		fmt.Println("Subscribe error ", err)
	}

	<-context.TODO().Done()
}
