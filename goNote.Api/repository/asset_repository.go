package repository

import (
	"context"
	"log"
	"time"

	"github.com/buwud/goNote/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type assetRepository struct {
	collection *mongo.Collection
}

func NewAssetRepository(collection *mongo.Collection) *assetRepository {
	return &assetRepository{collection: collection}
}

func (t *assetRepository) CreateAsset(asset *domain.Asset) (*mongo.InsertOneResult, error) {
	asset.CreatedAt = time.Now()
	asset.UpdatedAt = time.Now()
	return t.collection.InsertOne(context.Background(), asset)
}
func (t *assetRepository) DeleteAsset(assetID string, c *fiber.Ctx) error {
	objectID, err := primitive.ObjectIDFromHex(assetID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid asset ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = t.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
func (t *assetRepository) UpdateAsset(assetID string, asset *domain.Asset) error {
	objectID, err := primitive.ObjectIDFromHex(assetID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":         asset.Name,
		"base_unit":    asset.BaseUnit,
		"value_in_try": asset.ValueInTRY,
		"updated_at":   time.Now(),
	}}
	_, err = t.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (t *assetRepository) GetAll() (*[]domain.Asset, error) {
	var assets []domain.Asset
	cursor, err := t.collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var asset domain.Asset
		cursor.Decode(&asset)
		assets = append(assets, asset)
	}
	return &assets, nil
}
