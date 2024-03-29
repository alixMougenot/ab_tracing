package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"

	"github.com/alixMougenot/ab_tracing/db"
	"github.com/alixMougenot/ab_tracing/graph/model"
)

// CreatePlant is the resolver for the createPlant field.
func (r *mutationResolver) CreatePlant(ctx context.Context, plant model.PlantInput) (string, error) {
	pool := r.DBPool
	return db.CreatePlant(plant, ctx, pool)
}

// UpdatePlant is the resolver for the updatePlant field.
func (r *mutationResolver) UpdatePlant(ctx context.Context, id string, update *model.PlantInput) (bool, error) {
	pool := r.DBPool
	var input = model.PlantInput{}
	if update != nil {
		input = *update
	}
	result := db.UpdatePlant(id, input, ctx, pool)
	return result == nil, result
}

// DeletePlant is the resolver for the deletePlant field.
func (r *mutationResolver) DeletePlant(ctx context.Context, id string) (bool, error) {
	pool := r.DBPool
	result := db.DeletePlant(id, ctx, pool)
	return result == nil, result
}

// CreatePlantReproductionMaterial is the resolver for the createPlantReproductionMaterial field.
func (r *mutationResolver) CreatePlantReproductionMaterial(ctx context.Context, plantReproductionMaterial model.PlantReproductionMaterialInput) (string, error) {
	pool := r.DBPool
	return db.CreatePlantReproductionMaterial(plantReproductionMaterial, ctx, pool)
}

// UpdatePlantReproductionMaterial is the resolver for the updatePlantReproductionMaterial field.
func (r *mutationResolver) UpdatePlantReproductionMaterial(ctx context.Context, id string, update model.PlantReproductionMaterialInput) (bool, error) {
	pool := r.DBPool
	result := db.UpdatePlantReproductionMaterial(id, update, ctx, pool)
	return result == nil, result
}

// DeletePlantReproductionMaterial is the resolver for the deletePlantReproductionMaterial field.
func (r *mutationResolver) DeletePlantReproductionMaterial(ctx context.Context, id string) (bool, error) {
	pool := r.DBPool
	result := db.DeletePlantReproductionMaterial(id, ctx, pool)
	return result == nil, result
}

// CreatePlantTreatment is the resolver for the createPlantTreatment field.
func (r *mutationResolver) CreatePlantTreatment(ctx context.Context, plantTreatment model.PlantTreatmentInput) (string, error) {
	pool := r.DBPool
	return db.CreatePlantTreatment(plantTreatment, ctx, pool)
}

// UpdatePlantTreatment is the resolver for the updatePlantTreatment field.
func (r *mutationResolver) UpdatePlantTreatment(ctx context.Context, id string, update model.PlantTreatmentInput) (bool, error) {
	pool := r.DBPool
	result := db.UpdatePlantTreatment(id, update, ctx, pool)
	return result == nil, result
}

// DeletePlantTreatment is the resolver for the deletePlantTreatment field.
func (r *mutationResolver) DeletePlantTreatment(ctx context.Context, id string) (bool, error) {
	pool := r.DBPool
	result := db.DeletePlantTreatment(id, ctx, pool)
	return result == nil, result
}

// CreateGrowingMaterial is the resolver for the createGrowingMaterial field.
func (r *mutationResolver) CreateGrowingMaterial(ctx context.Context, growingMaterial model.GrowingMaterialInput) (string, error) {
	pool := r.DBPool
	return db.CreateGrowingMaterial(growingMaterial, ctx, pool)
}

// UpdateGrowingMaterial is the resolver for the updateGrowingMaterial field.
func (r *mutationResolver) UpdateGrowingMaterial(ctx context.Context, id string, update model.GrowingMaterialInput) (bool, error) {
	pool := r.DBPool
	result := db.UpdateGrowingMaterial(id, update, ctx, pool)
	return result == nil, result
}

// DeleteGrowingMaterial is the resolver for the deleteGrowingMaterial field.
func (r *mutationResolver) DeleteGrowingMaterial(ctx context.Context, id string) (bool, error) {
	pool := r.DBPool
	result := db.DeleteGrowingMaterial(id, ctx, pool)
	return result == nil, result
}

// CreateGatheringPlace is the resolver for the createGatheringPlace field.
func (r *mutationResolver) CreateGatheringPlace(ctx context.Context, gatheringPlace model.GatheringPlaceInput) (string, error) {
	pool := r.DBPool
	return db.CreateGatheringPlace(gatheringPlace, ctx, pool)
}

// UpdateGatheringPlace is the resolver for the updateGatheringPlace field.
func (r *mutationResolver) UpdateGatheringPlace(ctx context.Context, id string, update *model.GatheringPlaceInput) (bool, error) {
	pool := r.DBPool
	var input = model.GatheringPlaceInput{}
	if update != nil {
		input = *update
	}
	result := db.UpdateGatheringPlace(id, input, ctx, pool)
	return result == nil, result
}

// DeleteGatheringPlace is the resolver for the deleteGatheringPlace field.
func (r *mutationResolver) DeleteGatheringPlace(ctx context.Context, id string) (bool, error) {
	pool := r.DBPool
	result := db.DeleteGatheringPlace(id, ctx, pool)
	return result == nil, result
}

// CreateSupplyInfo is the resolver for the createSupplyInfo field.
func (r *mutationResolver) CreateSupplyInfo(ctx context.Context, supplyInfo model.SupplyInfoInput) (string, error) {
	pool := r.DBPool
	return db.CreateSupplyInfo(supplyInfo, ctx, pool)
}

// UpdateSupplyInfo is the resolver for the updateSupplyInfo field.
func (r *mutationResolver) UpdateSupplyInfo(ctx context.Context, id string, udpate model.SupplyInfoInput) (bool, error) {
	pool := r.DBPool
	result := db.UpdateSupplyInfo(id, udpate, ctx, pool)
	return result == nil, result
}

// DeleteSupplyInfo is the resolver for the deleteSupplyInfo field.
func (r *mutationResolver) DeleteSupplyInfo(ctx context.Context, id string) (bool, error) {
	pool := r.DBPool
	result := db.DeleteSupplyInfo(id, ctx, pool)
	return result == nil, result
}

// Plants is the resolver for the plants field.
func (r *queryResolver) Plants(ctx context.Context) ([]*model.Plant, error) {
	pool := r.DBPool
	return db.ListPlant(ctx, pool)
}

// Plant is the resolver for the plant field.
func (r *queryResolver) Plant(ctx context.Context, id string) (*model.Plant, error) {
	pool := r.DBPool
	return db.GetPlant(id, ctx, pool)
}

// PlantReproductionMaterials is the resolver for the plantReproductionMaterials field.
func (r *queryResolver) PlantReproductionMaterials(ctx context.Context) ([]*model.PlantReproductionMaterial, error) {
	pool := r.DBPool
	return db.ListPlantReproductionMaterials(ctx, pool)
}

// PlantReproductionMaterial is the resolver for the plantReproductionMaterial field.
func (r *queryResolver) PlantReproductionMaterial(ctx context.Context, id string) (*model.PlantReproductionMaterial, error) {
	pool := r.DBPool
	return db.GetPlantReproductionMaterial(id, ctx, pool)
}

// PlantTreatments is the resolver for the plantTreatments field.
func (r *queryResolver) PlantTreatments(ctx context.Context) ([]*model.PlantTreatment, error) {
	pool := r.DBPool
	return db.ListPlantTreatments(ctx, pool)
}

// PlantTreatment is the resolver for the plantTreatment field.
func (r *queryResolver) PlantTreatment(ctx context.Context, id string) (*model.PlantTreatment, error) {
	pool := r.DBPool
	return db.GetPlantTreatment(id, ctx, pool)
}

// GrowingMaterials is the resolver for the growingMaterials field.
func (r *queryResolver) GrowingMaterials(ctx context.Context) ([]*model.GrowingMaterial, error) {
	pool := r.DBPool
	return db.ListGrowingMaterials(ctx, pool)
}

// GrowingMaterial is the resolver for the growingMaterial field.
func (r *queryResolver) GrowingMaterial(ctx context.Context, id string) (*model.GrowingMaterial, error) {
	pool := r.DBPool
	return db.GetGrowingMaterial(id, ctx, pool)
}

// GatheringPlaces is the resolver for the gatheringPlaces field.
func (r *queryResolver) GatheringPlaces(ctx context.Context) ([]*model.GatheringPlace, error) {
	pool := r.DBPool
	return db.ListGatheringPlaces(ctx, pool)
}

// GatheringPlace is the resolver for the gatheringPlace field.
func (r *queryResolver) GatheringPlace(ctx context.Context, id string) (*model.GatheringPlace, error) {
	pool := r.DBPool
	return db.GetGatheringPlace(id, ctx, pool)
}

// SupplyInfos is the resolver for the supplyInfos field.
func (r *queryResolver) SupplyInfos(ctx context.Context) ([]*model.SupplyInfo, error) {
	pool := r.DBPool
	return db.ListSupplyInfos(ctx, pool)
}

// SupplyInfo is the resolver for the supplyInfo field.
func (r *queryResolver) SupplyInfo(ctx context.Context, id string) (*model.SupplyInfo, error) {
	pool := r.DBPool
	return db.GetSupplyInfo(id, ctx, pool)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
