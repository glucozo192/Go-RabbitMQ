package main

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	conn *amqp.Connection
}

func NewProducer(conn *amqp.Connection) *Producer {
	return &Producer{conn: conn}
}

func (p *Producer) TaskSend(ctx context.Context, payload string) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("verify_email_message", false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(payload),
	})
	if err != nil {
		return err
	}

	return nil

}
