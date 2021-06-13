package roster

import (
	App "UnitService/pkg/unit/app"
	DB "UnitService/pkg/unit/infrastructure/db"
	"encoding/json"
	uuid "github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RosterServices struct {
	channelRebbitMQ *amqp.Channel
	queueRebbitMQ amqp.Queue

	db DB.EventSQLDB
}

type EventJSON struct {
	Essence string `json:"essence"`
	TypeEvent string `json:"typeEvent"`
	IdRecord string `json:"idRecord"`
}

func CreateRosterServices(channelRebbitMQ *amqp.Channel, queueRebbitMQ amqp.Queue, db DB.EventSQLDB) App.RosterRepository {
	roster := RosterServices{channelRebbitMQ, queueRebbitMQ, db}
	roster.startEventPost()
	return &roster
}

func (roster *RosterServices) generateId() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id := u.String()
	return id, nil
}

func (roster *RosterServices) SendEvent(typeEvent string, idRecord string) error {
	id, err := roster.generateId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	eventDB := DB.EventDB {
		id,
		typeEvent,
		idRecord,
	}
	err = roster.db.InsertEvent(eventDB)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (roster *RosterServices) startEventPost() {
	go func() {
		for roster.db.Connection.Db != nil {
			event, err := roster.db.GetEventWithMinimalDate()
			if err != nil {
				log.Fatal(err)
				continue
			}
			if event.IdEvent == "" {
				continue
			}

			eventJson := EventJSON{
				"UNIT",
				event.TypeEvent,
				event.IdRecord,
			}

			b, err := json.Marshal(eventJson)
			if err != nil {
				log.Fatal(err)
				continue
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
				log.Fatal(err)
				continue
			}
			log.Print("Send Event: UNIT " + event.TypeEvent + " " + event.IdRecord)
			log.Print(b)

			err = roster.db.DeleteEvent(event.IdEvent)
			if err != nil {
				log.Fatal(err)
				continue
			}
		}
	}()
}