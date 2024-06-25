package jobs

import (
	"context"
	"log"
	"slot-machine-api/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitIndexes() {
	collection := utils.GetMongoCollection("slotMachine", "players")
	indexes := collection.Indexes()

	_, err := indexes.CreateOne(context.TODO(), mongo.IndexModel{
		Keys:    map[string]interface{}{"name": 1},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		log.Fatal(err)
	}
}
