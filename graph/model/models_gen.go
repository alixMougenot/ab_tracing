// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Ingredient interface {
	IsIngredient()
}

type GatheringPlace struct {
	ID                 string  `json:"id"`
	Name               *string `json:"name,omitempty"`
	Notes              string  `json:"notes"`
	Address            string  `json:"address"`
	Country            string  `json:"country"`
	IsOrganicCompliant bool    `json:"isOrganicCompliant"`
}

type GatheringPlaceInput struct {
	Name               *string `json:"name,omitempty"`
	Notes              *string `json:"notes,omitempty"`
	Address            *string `json:"address,omitempty"`
	Country            *string `json:"country,omitempty"`
	IsOrganicCompliant *bool   `json:"isOrganicCompliant,omitempty"`
}

type GrowingMaterial struct {
	ID                     string         `json:"id"`
	Name                   *string        `json:"name,omitempty"`
	Notes                  string         `json:"notes"`
	Visibility             Visibility     `json:"visibility"`
	IsOrganicCompliant     bool           `json:"isOrganicCompliant"`
	Quantity               float64        `json:"quantity"`
	Unit                   string         `json:"unit"`
	CreationDate           time.Time      `json:"creationDate"`
	ProductionSteps        []string       `json:"productionSteps"`
	AquisitionType         AquisitionType `json:"aquisitionType"`
	AquisitionPlaces       []string       `json:"aquisitionPlaces"`
	AquisitionPurshaseInfo []string       `json:"aquisitionPurshaseInfo"`
}

func (GrowingMaterial) IsIngredient() {}

type GrowingMaterialInput struct {
	Name                   *string         `json:"name,omitempty"`
	Notes                  *string         `json:"notes,omitempty"`
	Visibility             *Visibility     `json:"visibility,omitempty"`
	IsOrganicCompliant     *bool           `json:"isOrganicCompliant,omitempty"`
	Quantity               *float64        `json:"quantity,omitempty"`
	Unit                   *string         `json:"unit,omitempty"`
	CreationDate           *time.Time      `json:"creationDate,omitempty"`
	ProductionSteps        []string        `json:"productionSteps,omitempty"`
	AquisitionType         *AquisitionType `json:"aquisitionType,omitempty"`
	AquisitionPlaces       []string        `json:"aquisitionPlaces,omitempty"`
	AquisitionPurshaseInfo []string        `json:"aquisitionPurshaseInfo,omitempty"`
}

type Mutation struct {
}

type Plant struct {
	ID                     string                     `json:"id"`
	Name                   *string                    `json:"name,omitempty"`
	LatinName              string                     `json:"latinName"`
	Notes                  string                     `json:"notes"`
	Visibility             Visibility                 `json:"visibility"`
	IsOrganic              bool                       `json:"isOrganic"`
	PlantingDate           time.Time                  `json:"plantingDate"`
	IsStockPlant           bool                       `json:"isStockPlant"`
	Quantity               float64                    `json:"quantity"`
	Unit                   string                     `json:"unit"`
	AquisitionType         AquisitionType             `json:"aquisitionType"`
	PlantingSource         *PlantReproductionMaterial `json:"plantingSource,omitempty"`
	AquisitionPlaces       []string                   `json:"aquisitionPlaces"`
	AquisitionPurshaseInfo []string                   `json:"aquisitionPurshaseInfo"`
	GraftingSteps          []string                   `json:"graftingSteps"`
	MaturationSteps        []string                   `json:"maturationSteps"`
	TreatmentSteps         []string                   `json:"treatmentSteps"`
}

func (Plant) IsIngredient() {}

type PlantInput struct {
	Name                   *string         `json:"name,omitempty"`
	LatinName              *string         `json:"latinName,omitempty"`
	Notes                  *string         `json:"notes,omitempty"`
	Visibility             *Visibility     `json:"visibility,omitempty"`
	IsOrganic              *bool           `json:"isOrganic,omitempty"`
	PlantingDate           *time.Time      `json:"plantingDate,omitempty"`
	IsStockPlant           *bool           `json:"isStockPlant,omitempty"`
	Quantity               *float64        `json:"quantity,omitempty"`
	Unit                   *string         `json:"unit,omitempty"`
	AquisitionType         *AquisitionType `json:"aquisitionType,omitempty"`
	PlantingSource         *string         `json:"plantingSource,omitempty"`
	AquisitionPlaces       []string        `json:"aquisitionPlaces,omitempty"`
	AquisitionPurshaseInfo []string        `json:"aquisitionPurshaseInfo,omitempty"`
	GraftingSteps          []string        `json:"graftingSteps,omitempty"`
	MaturationSteps        []string        `json:"maturationSteps,omitempty"`
	TreatmentSteps         []string        `json:"treatmentSteps,omitempty"`
}

type PlantReproductionMaterial struct {
	ID                     string                     `json:"id"`
	Name                   *string                    `json:"name,omitempty"`
	LatinName              string                     `json:"latinName"`
	Type                   ReproductiveMaterialType   `json:"type"`
	Visibility             Visibility                 `json:"visibility"`
	Notes                  string                     `json:"notes"`
	IsOrganic              bool                       `json:"isOrganic"`
	ProductionDate         time.Time                  `json:"productionDate"`
	Quantity               float64                    `json:"quantity"`
	Unit                   string                     `json:"unit"`
	AquisitionType         AquisitionType             `json:"aquisitionType"`
	GerminationSource      *PlantReproductionMaterial `json:"germinationSource,omitempty"`
	HarvestSource          []string                   `json:"harvestSource"`
	AquisitionPlaces       []string                   `json:"aquisitionPlaces"`
	AquisitionPurshaseInfo []string                   `json:"aquisitionPurshaseInfo"`
	TreatmentSteps         []string                   `json:"treatmentSteps"`
}

type PlantReproductionMaterialInput struct {
	Name                   *string                   `json:"name,omitempty"`
	LatinName              *string                   `json:"latinName,omitempty"`
	Type                   *ReproductiveMaterialType `json:"type,omitempty"`
	Visibility             *Visibility               `json:"visibility,omitempty"`
	Notes                  *string                   `json:"notes,omitempty"`
	IsOrganic              *bool                     `json:"isOrganic,omitempty"`
	ProductionDate         *time.Time                `json:"productionDate,omitempty"`
	Quantity               *float64                  `json:"quantity,omitempty"`
	Unit                   *string                   `json:"unit,omitempty"`
	AquisitionType         *AquisitionType           `json:"aquisitionType,omitempty"`
	GerminationSource      *string                   `json:"germinationSource,omitempty"`
	HarvestSource          []string                  `json:"harvestSource,omitempty"`
	AquisitionPlaces       []string                  `json:"aquisitionPlaces,omitempty"`
	AquisitionPurshaseInfo []string                  `json:"aquisitionPurshaseInfo,omitempty"`
	TreatmentSteps         []string                  `json:"treatmentSteps,omitempty"`
}

type PlantTreatment struct {
	ID                        string         `json:"id"`
	Name                      *string        `json:"name,omitempty"`
	Visibility                Visibility     `json:"visibility"`
	Notes                     string         `json:"notes"`
	IsOrganicCompliant        bool           `json:"isOrganicCompliant"`
	Quantity                  float64        `json:"quantity"`
	Unit                      string         `json:"unit"`
	CreationDate              time.Time      `json:"creationDate"`
	HomeProductionIngredients []string       `json:"homeProductionIngredients"`
	AquisitionType            AquisitionType `json:"aquisitionType"`
	AquisitionPlaces          []string       `json:"aquisitionPlaces"`
	AquisitionPurshaseInfo    []string       `json:"aquisitionPurshaseInfo"`
}

func (PlantTreatment) IsIngredient() {}

type PlantTreatmentInput struct {
	Name                      *string         `json:"name,omitempty"`
	Visibility                *Visibility     `json:"visibility,omitempty"`
	Notes                     *string         `json:"notes,omitempty"`
	IsOrganicCompliant        *bool           `json:"isOrganicCompliant,omitempty"`
	Quantity                  *float64        `json:"quantity,omitempty"`
	Unit                      *string         `json:"unit,omitempty"`
	CreationDate              *time.Time      `json:"creationDate,omitempty"`
	HomeProductionIngredients []string        `json:"homeProductionIngredients,omitempty"`
	AquisitionType            *AquisitionType `json:"aquisitionType,omitempty"`
	AquisitionPlaces          []string        `json:"aquisitionPlaces,omitempty"`
	AquisitionPurshaseInfo    []string        `json:"aquisitionPurshaseInfo,omitempty"`
}

type Query struct {
}

type SupplyInfo struct {
	ID         string     `json:"id"`
	Visibility Visibility `json:"visibility"`
	Name       *string    `json:"name,omitempty"`
	Supplier   string     `json:"supplier"`
	Bill       string     `json:"bill"`
	Notes      string     `json:"notes"`
	Country    *string    `json:"country,omitempty"`
}

type SupplyInfoInput struct {
	Visibility *Visibility `json:"visibility,omitempty"`
	Name       *string     `json:"name,omitempty"`
	Supplier   *string     `json:"supplier,omitempty"`
	Bill       *string     `json:"bill,omitempty"`
	Notes      *string     `json:"notes,omitempty"`
	Country    *string     `json:"country,omitempty"`
}

type AquisitionType string

const (
	AquisitionTypeGrown     AquisitionType = "GROWN"
	AquisitionTypePurchased AquisitionType = "PURCHASED"
	AquisitionTypeGathered  AquisitionType = "GATHERED"
	AquisitionTypeHomeMade  AquisitionType = "HOME_MADE"
)

var AllAquisitionType = []AquisitionType{
	AquisitionTypeGrown,
	AquisitionTypePurchased,
	AquisitionTypeGathered,
	AquisitionTypeHomeMade,
}

func (e AquisitionType) IsValid() bool {
	switch e {
	case AquisitionTypeGrown, AquisitionTypePurchased, AquisitionTypeGathered, AquisitionTypeHomeMade:
		return true
	}
	return false
}

func (e AquisitionType) String() string {
	return string(e)
}

func (e *AquisitionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AquisitionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AquisitionType", str)
	}
	return nil
}

func (e AquisitionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReproductiveMaterialType string

const (
	ReproductiveMaterialTypeSeed     ReproductiveMaterialType = "SEED"
	ReproductiveMaterialTypeCutting  ReproductiveMaterialType = "CUTTING"
	ReproductiveMaterialTypeGrafting ReproductiveMaterialType = "GRAFTING"
)

var AllReproductiveMaterialType = []ReproductiveMaterialType{
	ReproductiveMaterialTypeSeed,
	ReproductiveMaterialTypeCutting,
	ReproductiveMaterialTypeGrafting,
}

func (e ReproductiveMaterialType) IsValid() bool {
	switch e {
	case ReproductiveMaterialTypeSeed, ReproductiveMaterialTypeCutting, ReproductiveMaterialTypeGrafting:
		return true
	}
	return false
}

func (e ReproductiveMaterialType) String() string {
	return string(e)
}

func (e *ReproductiveMaterialType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReproductiveMaterialType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReproductiveMaterialType", str)
	}
	return nil
}

func (e ReproductiveMaterialType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Visibility string

const (
	VisibilityPublicFacing Visibility = "PUBLIC_FACING"
	VisibilityInternal     Visibility = "INTERNAL"
)

var AllVisibility = []Visibility{
	VisibilityPublicFacing,
	VisibilityInternal,
}

func (e Visibility) IsValid() bool {
	switch e {
	case VisibilityPublicFacing, VisibilityInternal:
		return true
	}
	return false
}

func (e Visibility) String() string {
	return string(e)
}

func (e *Visibility) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Visibility(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Visibility", str)
	}
	return nil
}

func (e Visibility) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
