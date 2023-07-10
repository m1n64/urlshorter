package services

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"urlshorter/models"
	"urlshorter/utils"
)

type Request struct {
	Slug string `json:"slug"`
}

type Response struct {
	Link string `json:"link"`
}

func ListenQueue() {
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

	msgs, err := ch.Consume(
		q.Name, // очередь
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

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var msg Request
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Println("Ошибка при десериализации JSON:", err)
				continue
			}

			log.Printf("Получено сообщение: %s", msg.Slug)

			db := utils.GetDBConnection()
			defer db.Close()

			linkModel := &models.LinkModel{DB: db}

			link, err := linkModel.GetLinkBySlug(msg.Slug)

			fmt.Println(err)

			err = PublishLink(ch, link.Link)
			if err != nil {
				log.Println("Ошибка при публикации ссылки:", err)
			}

		}
	}()
	<-forever
}

func PublishLink(ch *amqp.Channel, link string) error {
	q, err := ch.QueueDeclare(
		"redirectQueue", // Имя очереди Redirect
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(Response{Link: link})
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
