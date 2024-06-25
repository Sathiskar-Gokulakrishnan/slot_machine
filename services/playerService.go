package services

import (
	"context"
	"slot-machine-api/models"
	"slot-machine-api/utils"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePlayer(player models.Player) (*mongo.InsertOneResult, error) {
	collection := utils.GetMongoCollection("slotMachine", "players")
	return collection.InsertOne(context.TODO(), player)
}

func GetPlayerByID(id string) (models.Player, error) {
	var player models.Player
	redisClient := utils.GetRedisClient()
	cacheKey := "player:" + id

	cachedPlayer, err := redisClient.Get(utils.Ctx, cacheKey).Bytes()
	if err == redis.Nil {
		collection := utils.GetMongoCollection("slotMachine", "players")
		err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&player)
		if err != nil {
			return player, err
		}

		playerBytes, _ := utils.Marshal(player)
		redisClient.Set(utils.Ctx, cacheKey, playerBytes, 0)
	} else if err == nil {
		utils.Unmarshal(cachedPlayer, &player)
	} else {
		return player, err
	}

	return player, nil
}

func SuspendPlayer(id string) (*mongo.UpdateResult, error) {
	collection := utils.GetMongoCollection("slotMachine", "players")
	return collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": "suspended"}})
}
