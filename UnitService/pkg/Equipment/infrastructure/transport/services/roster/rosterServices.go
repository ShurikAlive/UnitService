package roster

import (
	App "UnitService/pkg/equipment/app"
	"encoding/json"
	"github.com/streadway/amqp"
)

type RosterServices struct {
	channelRebbitMQ *amqp.Channel
	queueRebbitMQ amqp.Queue
}

type EventJSON struct {
	Essence string `json:"essence"`
	TypeEvent string `json:"typeEvent"`
	IdRecord string `json:"idRecord"`
}

func CreateRosterServices(channelRebbitMQ *amqp.Channel, queueRebbitMQ amqp.Queue) App.RosterRepository {
	return &RosterServices{channelRebbitMQ, queueRebbitMQ}
}

func (roster *RosterServices) SendEvent(typeEvent string, idRecord string) error {

	event := EventJSON{
		"EQUIPMENT",
		typeEvent,
		idRecord,
	}
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = roster.channelRebbitMQ.Publish(
		"",     // exchange
		roster.queueRebbitMQ.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	if err != nil {
		return err
	}
	return nil
}
