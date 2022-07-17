package gonsq

import (
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
)

const (
	nsqLookup string = "10.1.26.150:4161"
)

func handler(message *nsq.Message) error {
	log.Println(string(message.Body))
	return nil
}

func TestProduct(t *testing.T) {
	Init(nsqLookup, 10)
	PublishAsync(TestNSQ, "Saber_Test", []byte("wtq"), nil)
}

func TestConsumer(t *testing.T) {
	Init(nsqLookup, 10)
	NewConsumer(TestNSQ, "Saber_Test", "wtq-channl", nsq.HandlerFunc(handler))
	select {}
}

func TestConsumer2(t *testing.T) {
	Init(nsqLookup, 10)
	NewConsumer(TestNSQ, "Saber_Test", "wtq-channl2", nsq.HandlerFunc(handler))
	select {}
}
