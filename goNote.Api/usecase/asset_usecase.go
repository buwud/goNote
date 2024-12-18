package usecase

import (
	"context"
	"time"

	"github.com/buwud/goNote/domain"
	"github.com/buwud/goNote/domain/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type assetUseCase struct {
	assetRepo domain.AssetRepository
}

func NewAssetUseCase(assetRepo domain.AssetRepository) (domain.AssetUseCase, error) {
	return &assetUseCase{assetRepo: assetRepo}, nil
}

func (t *assetUseCase) CreateAsset(asset *domain.Asset) (*mongo.InsertOneResult, error) {
	return t.assetRepo.CreateAsset(asset)
}
func (t *assetUseCase) DeleteAsset(assetID string, c *fiber.Ctx) error {
	return t.assetRepo.DeleteAsset(assetID, c)
}
func (t *assetUseCase) UpdateAsset(assetID string, asset *domain.Asset) error {
	return t.assetRepo.UpdateAsset(assetID, asset)
}
func (t *assetUseCase) GetAll() (*[]domain.Asset, error) {
	return t.assetRepo.GetAll()
}
func (t *assetUseCase) CreateUserAsset(userAsset *domain.UserAsset) (*mongo.InsertOneResult, error) {
	return t.assetRepo.CreateUserAsset(userAsset)
}
func (t *assetUseCase) GetUserAssetHistory(userID primitive.ObjectID, startDate time.Time, endDate time.Time, page int, pageSize int, c context.Context) (map[string]interface{}, error) {
	return t.assetRepo.GetUserAssetHistory(userID, startDate, endDate, page, pageSize, c)
}
func (t *assetUseCase) UpdateUserAsset(userAssetID primitive.ObjectID, userAsset *models.UpdateUserAsset) error {
	return t.assetRepo.UpdateUserAsset(userAssetID, userAsset)
}
func (t *assetUseCase) DeleteUserAsset(userAssetID primitive.ObjectID) error {
	return t.assetRepo.DeleteUserAsset(userAssetID)
}
