package price_model

import (
	"context"
	"log"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategoryminor"
)

type repository struct {
	db *ent.Client
}

func NewPriceModelRepo(db *database.Adapter) gateways.PriceModelRepo {
	return &repository{
		db: db.DB,
	}
}

func (r *repository) Insert(model *models.PriceModelRequest) (*ent.PriceModel, error) {
	result, err := r.db.PriceModel.Create().
		SetName(model.Name).
		SetInitials(model.Initials).
		SetFormula(model.Formula).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Read(id int) (*ent.PriceModel, error) {
	result, err := r.db.PriceModel.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) ReadAll(limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	modelQuery := r.db.PriceModel.Query()
	totalRecords, err := modelQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := modelQuery.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  results,
	}, nil
}

func (r *repository) ReadAllPercentage(limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	catQuery := r.db.ProductCategoryMinor.Query().Where(productcategoryminor.PercentageNotIn(0))
	totalRecords, err := catQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := catQuery.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  results,
	}, nil
}

func (r *repository) Update(id int, model *models.PriceModelRequest) (*ent.PriceModel, error) {
	result, err := r.db.PriceModel.UpdateOneID(id).
		SetName(model.Name).
		SetInitials(model.Initials).
		SetFormula(model.Formula).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdatePercentage(category, percentage int) (*ent.ProductCategoryMinor, error) {
	result, err := r.db.ProductCategoryMinor.UpdateOneID(category).
		SetPercentage(percentage).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.PriceModel.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeletePercentage(id int) error {
	rr, err := r.db.ProductCategoryMinor.UpdateOneID(id).
		SetPercentage(0).
		Save(context.Background())
	if err != nil {
		return err
	}
	log.Println(rr)
	return nil
}
