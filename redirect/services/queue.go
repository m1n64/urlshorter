package services

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
)

type Request struct {
	Slug string `json:"slug"`
}

type Response struct {
	Link string `json:"link"`
}

type Callback struct {
	Link string `json:"link"`
}

func PublishQueue(slug string) {
	// Подключение к RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"linksQueue", // Имя очереди
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	body := Request{Slug: slug}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		})
	if err != nil {
		log.Fatal(err)
	}
}

func ListenLink() string {
	// Подключение к RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"redirectQueue", // Имя очереди Redirect
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // Очередь Redirect
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() string {
		for d := range msgs {
			var response Response
			err := json.Unmarshal(d.Body, &response)
			if err != nil {
				log.Println("Ошибка при десериализации JSON:", err)
				continue
			}

			log.Printf("Получена ссылка: %s", response.Link)

			return response.Link
		}
		return ""
	}()
}
