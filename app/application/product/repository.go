package product

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategorymajor"
	"github.com/SeyramWood/ent/productcategoryminor"
)

type repository struct {
	db *ent.Client
}

func NewProductRepo(db *database.Adapter) gateways.ProductRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(prod *models.Product, imageUrl string) (*ent.Product, error) {
	ctx := context.Background()

	mq := r.db.Merchant.Query().Where(merchant.ID(prod.Merchant)).OnlyX(ctx)
	major := r.db.ProductCategoryMajor.Query().Where(productcategorymajor.ID(prod.CategoryMajor)).OnlyX(ctx)
	minor := r.db.ProductCategoryMinor.Query().Where(productcategoryminor.ID(prod.CategoryMinor)).OnlyX(ctx)

	result, err := r.db.Product.Create().SetMerchant(mq).SetMajor(major).SetMinor(minor).
		SetQuantity(uint32(prod.Quantity)).
		SetWeight(uint32(prod.Weight)).
		SetUnit(prod.Unit).
		SetName(prod.Name).
		SetPrice(prod.Price).
		SetPromoPrice(prod.PromoPrice).
		SetDescription(prod.Description).
		SetImage(imageUrl).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating product: %w", err)
	}
	return result, nil

}

func (r *repository) Read(id int) (*ent.Product, error) {

	result, err := r.db.Product.Query().Where(product.ID(id)).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().WithMinor().
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return result, nil

}
func (r *repository) ReadBySupplierMerchant(id int) (*ent.Product, error) {

	result, err := r.db.Product.Query().Where(product.ID(id)).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return result, nil

}
func (r *repository) ReadByRetailMerchant(id int) (*ent.Product, error) {

	result, err := r.db.Product.Query().Where(product.ID(id)).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return result, nil

}

func (r *repository) ReadAll() ([]*ent.Product, error) {
	products, err := r.db.Product.Query().
		Order(ent.Desc(product.FieldCreatedAt)).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return products, nil
}
func (r *repository) ReadBySlugRetailMerchantCategoryMinor(slug string) ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		Where(productcategoryminor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (r *repository) ReadBySlugRetailMerchantCategoryMajor(slug string) ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		Where(productcategorymajor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) ReadBySlugSupplierMerchantCategoryMinor(slug string) ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		Where(productcategoryminor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) ReadBySlugSupplierMerchantCategoryMajor(slug string) ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		Where(productcategorymajor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return products, nil

}

func (r *repository) ReadAllRetailMerchantCategoryMinor() ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) ReadAllRetailMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) ReadAllSupplierMerchantCategoryMinor() ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) ReadAllSupplierMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name")
								},
							)
						},
					)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) ReadAllBySupplierMerchant(merchantId int) ([]*ent.Product, error) {

	products, err := r.db.Product.Query().
		Order(ent.Desc(product.FieldCreatedAt)).
		Where(product.HasMerchantWith(merchant.ID(merchantId))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil

}

func (r *repository) ReadAllByRetailMerchant(merchantId int) ([]*ent.Product, error) {

	products, err := r.db.Product.Query().
		Order(ent.Desc(product.FieldCreatedAt)).
		Where(product.HasMerchantWith(merchant.ID(merchantId))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil

}

func (r *repository) ReadBestSellerBySupplierMerchant(limit, offset int) ([]*ent.Product, error) {
	products, err := r.db.Product.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		Where(product.HasMerchantWith(merchant.Type("supplier"))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return products, nil

}

func (r *repository) ReadBestSellerRetailMerchant() ([]*ent.Product, error) {

	products, err := r.db.Product.Query().
		Order(ent.Desc(product.FieldCreatedAt)).
		Where(product.HasMerchantWith(merchant.Type("retailer"))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name")
					},
				)
			},
		).
		WithMajor().
		WithMinor().
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return products, nil

}

func (r *repository) Update(i *models.Product) (*models.Product, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

func (r *repository) Delete(id string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
