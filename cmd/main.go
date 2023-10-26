package main

import (
	"context"
	"fmt"
	"homework-3/config"
	"homework-3/internal/infrastructure/kafka"
	receiverhandlers "homework-3/internal/infrastructure/kafka/receiver_handlers"
	"homework-3/internal/pkg/app"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository/postgresql"
	"homework-3/internal/pkg/routers"
	"log"
	"net/http"

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

	topic := "library"
	sender := kafka.NewKafkaSender(producer, topic)

	a := app.NewApp(authorRepo, bookRepo)

	go kafkaReceiverStart(brokers, topic)

	router := mux.NewRouter()
	routers.CreateAuthorRouter(router, a, sender)
	routers.CreateBookSubRouter(router, a)
	http.Handle("/", router)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func kafkaReceiverStart(brokers []string, topic string) {
	kafkaConsumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		fmt.Println(err)
	}

	// обработчики по каждому из топиков
	handlers := map[string]kafka.KafkaReceiverHandler{
		topic: receiverhandlers.NewStdoutKafkaReceiverHandler(),
	}

	messageReceiver := kafka.NewReceiver(kafkaConsumer, handlers)

	// При условии одного инстанса подходит идеально
	err = messageReceiver.Subscribe(topic)
	if err != nil {
		fmt.Println("Subscribe error ", err)
	}

	<-context.TODO().Done()
}
