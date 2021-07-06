package m

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	_client     *mongo.Client
	connectOnce sync.Once
)

func Connect(uri string) {
	connectOnce.Do(func() {
		c, err := mongo.NewClient(options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}
		_client = c
		err = _client.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	})
}

func GetClient() *mongo.Client {
	return _client
}

func Disconnect() {
	if _client != nil {
		if err := _client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}
}
