package main

import (
	"fmt"
	"log"
	"github.com/IBM/sarama"
)

func main(){
	fmt.Println("test")
	producer , err := initProducer()
	if err != nil{
		log.Fatalln("error initialize producer:", err)
	}

	for i := 0; i < 100; i++{
		message := &sarama.ProducerMessage{
			Topic: "test-topic",
			Value: sarama.StringEncoder("testMessage" + string(i)),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil{
			fmt.Println(err)
			log.Printf("failed to send message!!:", err)
		} else {
			log.Printf("partition: %d, offset: %d", partition, offset)
		}
		fmt.Printf("ok! %d", i)
	}

}

func initProducer()(sarama.SyncProducer, error){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{"localhost:29092"}
	producer, err := sarama.NewSyncProducer(brokers, config)

	return producer, err
}