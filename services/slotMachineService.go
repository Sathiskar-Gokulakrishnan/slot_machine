package services

import (
	"context"
	"math/rand"
	"slot-machine-api/models"
	"slot-machine-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	RTP_THRESHOLD = 0.975
)

type PlayResult struct {
	WinAmount int    `json:"win_amount"`
	Outcome   string `json:"outcome"`
}

func Play(playerID string) (PlayResult, error) {
	redisClient := utils.GetRedisClient()
	statsKey := "stats:rtp"
	totalPlaysKey := "stats:total_plays"

	result := PlayResult{}

	// Simulate slot machine play
	rand.Seed(time.Now().UnixNano())
	outcome := rand.Intn(100)

	// Determine win/lose
	if outcome < 50 {
		result.WinAmount = rand.Intn(10) + 1 // Win small
		result.Outcome = "win_small"
	} else if outcome < 90 {
		result.WinAmount = -(rand.Intn(10) + 1) // Lose small
		result.Outcome = "lose_small"
	} else if outcome < 98 {
		result.WinAmount = rand.Intn(100) + 50 // Win big
		result.Outcome = "win_big"
	} else {
		result.WinAmount = -(rand.Intn(100) + 50) // Lose big
		result.Outcome = "lose_big"
	}

	// Store game result in MongoDB
	game := models.Game{
		PlayerID:  playerID,
		WinAmount: result.WinAmount,
		Outcome:   result.Outcome,
		PlayedAt:  time.Now(),
	}
	collection := utils.GetMongoCollection("slotMachine", "games")
	_, err := collection.InsertOne(context.TODO(), game)
	if err != nil {
		return result, err
	}

	// Update Redis statistics
	rtp, _ := redisClient.Get(utils.Ctx, statsKey).Float64()
	totalPlays, _ := redisClient.Get(utils.Ctx, totalPlaysKey).Int64()

	rtp = ((rtp * float64(totalPlays)) + float64(result.WinAmount)) / float64(totalPlays+1)
	totalPlays++

	redisClient.Set(utils.Ctx, statsKey, rtp, 0)
	redisClient.Set(utils.Ctx, totalPlaysKey, totalPlays, 0)

	return result, nil
}

func GetGamesByPlayerID(playerID string) ([]models.Game, error) {
	var games []models.Game
	collection := utils.GetMongoCollection("slotMachine", "games")
	cursor, err := collection.Find(context.TODO(), bson.M{"player_id": playerID})
	if err != nil {
		return games, err
	}
	if err = cursor.All(context.TODO(), &games); err != nil {
		return games, err
	}
	return games, nil
}
