package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"kafka/consumer/data"

	"github.com/IBM/sarama"
)

func main(){
	db := data.NewData()
	defer db.CloseDb()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	brokers := []string{"localhost:29092"}
	topic := "test-topic"

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalln("Error creating consumer", err)
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("Error creating partition consumer", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	consumed := 0

	ConsumerLoop:
		for {
			select{
			case msg := <-partitionConsumer.Messages():
				err := db.InsertQuery(consumed ,string(msg.Value))
				if err != nil{
					log.Fatalln(err)
				}
				consumed++
				fmt.Println(consumed)
			case <-signals:
				break ConsumerLoop
			}
		}
		log.Printf("Consumed: %d\n", consumed)

		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln("Failed to close partition consumer", err)
		}
	
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln("Failed to close consumer", err)
		}
	
}

