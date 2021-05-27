package database

import (
	"context"
	"log"
	"time"

	"github.com/jailtonjunior94/go-challenge/partner/infrastructure/environments"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type IMongoConnection interface {
	Disconnect()
	CreateIndex() error
	GetCollection(dbName, collectionName string) (*mongo.Collection, error)
}

type MongoConnection struct {
	Client *mongo.Client
}

func NewConnection() IMongoConnection {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(environments.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConnection{Client: client}
}

func (m *MongoConnection) GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	collection := m.Client.Database(dbName).Collection(collectionName)
	return collection, nil
}

func (m *MongoConnection) Disconnect() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func (m *MongoConnection) CreateIndex() error {
	indexOptions := options.CreateIndexes().SetMaxTime(time.Second * 10)
	pointIndex := mongo.IndexModel{
		Options: options.Index().SetBackground(true),
		Keys:    bsonx.MDoc{"address.location": bsonx.String("2dsphere")},
	}

	addressIndexes := m.Client.Database(environments.Database).Collection(environments.Collection).Indexes()
	_, err := addressIndexes.CreateOne(context.Background(), pointIndex, indexOptions)
	if err != nil {
		return err
	}

	return nil
}
