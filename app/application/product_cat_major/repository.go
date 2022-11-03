package product_cat_major

import (
	"context"
	"fmt"
	"strings"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/product"
)

type repository struct {
	db *ent.Client
}

func NewProductCatMajorRepo(db *database.Adapter) gateways.ProductCatMajorRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(cat *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error) {
	category, err := r.db.ProductCategoryMajor.Create().
		SetCategory(cat.Category).
		SetSlug(strings.ToLower(strings.Replace(cat.Category, " ", "-", -1))).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)

		return nil, fmt.Errorf("failed creating product category major: %w", err)
	}
	return category, nil

}

func (r *repository) Read(id int) (*ent.ProductCategoryMajor, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.ProductCategoryMajor, error) {
	ctx := context.Background()
	cats, err := r.db.ProductCategoryMajor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Select(product.FieldID)
			},
		).
		WithMinors().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *repository) Update(i *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
