package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/nicelogic/config"
	"github.com/nicelogic/pubsub/contactevent"
)

type Pubsub struct {
	Client               pulsar.Client
	DefaultEventProducer pulsar.Producer
}

func (pubsub *Pubsub) Init(configFilePath string) (err error) {
	type ClientConfig struct {
		Url                 string
		Operation_timeout   int
		Connection_timeout  int
		Default_event_topic string
	}
	clientConfig := ClientConfig{}
	err = config.Init(configFilePath, &clientConfig)
	if err != nil {
		log.Println("config init err: ", err)
		return err
	}
	log.Printf("(%#v)\n", clientConfig)

	pubsub.Client, err = pulsar.NewClient(pulsar.ClientOptions{
		URL:               clientConfig.Url,
		OperationTimeout:  time.Duration(clientConfig.Operation_timeout) * time.Second,
		ConnectionTimeout: time.Duration(clientConfig.Connection_timeout) * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client(%v)", err)
	}

	pubsub.DefaultEventProducer, err = pubsub.Client.CreateProducer(pulsar.ProducerOptions{
		Topic: clientConfig.Default_event_topic,
	})
	if err != nil {
		log.Printf("pulsar create producer err(%v)", err)
		return err
	}
	return err
}

func (pubsub *Pubsub) SendAsync(ctx context.Context, payload any) error {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		log.Printf("json marshal err(%v), but still return true(ntf not affect whether success)\n", err)
		return err
	}
	pubsub.DefaultEventProducer.SendAsync(ctx,
		&pulsar.ProducerMessage{Payload: payloadJson},
		func(mi pulsar.MessageID, pm *pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Printf("pulsar send ntf(%v), err(%v)\n", payload, err)
			}
			log.Printf("pulsar send ntf(%v), success, msgId(%v)\n", payload, mi)
		})
	return nil
}

func (pubsub *Pubsub) AsyncPublishApplyAddContact(ctx context.Context, event *contactevent.Event) error {
	return pubsub.SendAsync(ctx, event)
}
