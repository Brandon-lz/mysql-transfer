package main

import (
	"bytes"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"io/ioutil"
	"net/http"
)

func CreateKafkaConsumer() *kafka.Consumer {

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	return consumer
}

func SubscribeTopic(consumer *kafka.Consumer) {
	// 貌似一个观察者只能订阅一个topic
	// consumer.Subscribe("postgres.public.Product", nil)
	consumer.Subscribe("dbserver1.inventory.customers", nil)        // 监控dbserver1.inventory.customers表的变化
	// consumer.Subscribe("dbserver1.inventory.products_on_hand", nil)        // 监控dbserver1.inventory.geom

	fmt.Println("Subscribed to product topic")
}

func ReadTopicMessages(consumer *kafka.Consumer) {

	var message string
	for {
		fmt.Println("-------------------------------Waiting for message---------------------------------------")
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			message = message + string(msg.Value)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		consumer.CommitMessage(msg)
	}
}


func RegisterConnector() *http.Response {
	plan, _ := ioutil.ReadFile("./connectors/debezium-connector.json")
	response, err := http.Post("http://localhost:8083/connectors/", "application/json", bytes.NewBuffer(plan))

	if err != nil {
		panic(err)
	}

	return response
}

func CheckConnector() {
	response, err := http.Get("http://localhost:8083/connectors/product_connector")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		RegisterConnector()
	}

	// If you want to show response, uncomment following lines

	//body, _ := ioutil.ReadAll(response.Body)

	//fmt.Println(string(body))
}

func main() {
	consumer := CreateKafkaConsumer()
	defer consumer.Close()

	CheckConnector()

	SubscribeTopic(consumer)

	ReadTopicMessages(consumer)
}
