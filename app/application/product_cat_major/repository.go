package product_cat_major

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewProductCatMajorRepo(db *database.Adapter) gateways.ProductCatMajorRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(cat *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error) {

	category, err := r.db.ProductCategoryMajor.Create().SetCategory(cat.Category).Save(context.Background())

	if err != nil {
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

	cats, err := r.db.ProductCategoryMajor.Query().WithMinors().All(context.Background())
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (a *repository) Update(i *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error) {
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
