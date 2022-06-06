package product_cat_minor

import (
	"context"
	"fmt"

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

func (r *repository) Insert(cat *models.ProductCategoryMinor) (*ent.ProductCategoryMinor, error) {

	major := r.db.ProductCategoryMajor.Query().Where(productcategorymajor.ID(cat.CategoryMajor)).OnlyX(context.Background())
	category, err := r.db.ProductCategoryMinor.Create().SetMajor(major).
		SetCategory(cat.Category).Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating product category minor: %w", err)
	}
	return category, nil

}

func (r *repository) Read(id int) (*ent.ProductCategoryMinor, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.ProductCategoryMinor, error) {

	cats, err := r.db.ProductCategoryMinor.Query().WithMajor().All(context.Background())

	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (a *repository) Update(i *models.ProductCategoryMinor) (*models.ProductCategoryMinor, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

//DeleteBook is a mongo repository that helps to delete books
func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
