package product_cat_minor

import (
	"context"
	"fmt"
	"strings"

	"github.com/SeyramWood/ent/productcategoryminor"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/productcategorymajor"
)

type repository struct {
	db *ent.Client
}

func NewProductCatMinorRepo(db *database.Adapter) gateways.ProductCatMinorRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(cat *models.ProductCategoryMinor, image string) (*ent.ProductCategoryMinor, error) {

	major := r.db.ProductCategoryMajor.Query().Where(productcategorymajor.ID(cat.CategoryMajor)).OnlyX(context.Background())
	category, err := r.db.ProductCategoryMinor.Create().SetMajor(major).
		SetCategory(cat.Category).
		SetImage(image).
		SetSlug(strings.ToLower(strings.Replace(cat.Category, " ", "-", -1))).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating product category minor: %w", err)
	}
	res, err := r.Read(category.ID)
	if err != nil {
		return nil, fmt.Errorf("failed reading product category minor: %w", err)
	}
	return res, nil

}

func (r *repository) Read(id int) (*ent.ProductCategoryMinor, error) {
	cat, err := r.db.ProductCategoryMinor.Query().Where(productcategoryminor.ID(id)).WithMajor().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (r *repository) ReadAll(limit, offset int) ([]*ent.ProductCategoryMinor, error) {

	cats, err := r.db.ProductCategoryMinor.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(productcategoryminor.FieldCreatedAt)).
		WithMajor().All(context.Background())

	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *repository) Update(id int, request *models.ProductCategoryMinorUpdate) (*ent.ProductCategoryMinor, error) {
	ctx := context.Background()
	result, err := r.db.ProductCategoryMinor.UpdateOneID(id).SetCategory(request.Category).Save(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r *repository) UpdateImage(id int, imagePath string) (string, error) {
	ctx := context.Background()
	_, err := r.db.ProductCategoryMinor.UpdateOneID(id).SetImage(imagePath).Save(ctx)
	if err != nil {
		return "", err
	}
	return imagePath, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.ProductCategoryMinor.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
