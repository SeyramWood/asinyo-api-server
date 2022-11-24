package product

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

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
						query.Select("id", "name", "coordinate")
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
						query.Select("id", "name", "coordinate")
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
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAll(limit, offset int) ([]*ent.Product, error) {
	products, err := r.db.Product.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllBySlugCategoryMajor(merchantType, slug string, limit, offset int) (
	[]*ent.Product, error,
) {
	products, err := r.db.ProductCategoryMajor.Query().
		Where(productcategorymajor.Slug(slug)).QueryProducts().
		Where(product.HasMerchantWith(merchant.Type(merchantType))).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		WithMajor().
		WithMinor().
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllBySlugCategoryMinor(merchantType, slug string, limit, offset int) (
	[]*ent.Product, error,
) {
	products, err := r.db.ProductCategoryMinor.Query().
		Where(productcategoryminor.Slug(slug)).QueryProducts().
		Where(product.HasMerchantWith(merchant.Type(merchantType))).
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		WithMajor().
		WithMinor().
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadBySlugRetailMerchantCategoryMinor(slug string, limit, offset int) (
	[]*ent.ProductCategoryMinor, error,
) {
	products, err := r.db.ProductCategoryMinor.Query().
		Where(productcategoryminor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadBySlugRetailMerchantCategoryMajor(slug string, limit, offset int) (
	[]*ent.ProductCategoryMajor, error,
) {
	products, err := r.db.ProductCategoryMajor.Query().
		Where(productcategorymajor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadBySlugSupplierMerchantCategoryMinor(
	slug string, limit, offset int,
) ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		Where(productcategoryminor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadBySlugSupplierMerchantCategoryMajor(
	slug string, limit, offset int,
) ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		Where(productcategorymajor.Slug(slug)).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllRetailMerchantCategoryMinor(limit, offset int) ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllRetailMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("retailer"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithRetailer()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllSupplierMerchantCategoryMinor(limit, offset int) ([]*ent.ProductCategoryMinor, error) {
	products, err := r.db.ProductCategoryMinor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllSupplierMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error) {
	products, err := r.db.ProductCategoryMajor.Query().
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.Where(
					product.HasMerchantWith(
						merchant.Type("supplier"),
					),
				).
					Limit(limit).
					Offset(offset).
					Order(ent.Desc(product.FieldCreatedAt)).
					WithMajor().
					WithMinor().
					WithMerchant(
						func(mq *ent.MerchantQuery) {
							mq.WithSupplier()
							mq.WithStore(
								func(query *ent.MerchantStoreQuery) {
									query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllBySupplierMerchant(merchantId, limit, offset int) ([]*ent.Product, error) {

	products, err := r.db.Product.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		Where(product.HasMerchantWith(merchant.ID(merchantId))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadAllByRetailMerchant(merchantId, limit, offset int) ([]*ent.Product, error) {

	products, err := r.db.Product.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		Where(product.HasMerchantWith(merchant.ID(merchantId))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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
		Order(ent.Desc(product.FieldBestDeal)).
		Where(product.HasMerchantWith(merchant.Type("supplier"))).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.GTE(product.FieldBestDeal, 0))
			},
		).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadBestSellerRetailMerchant(limit, offset int) ([]*ent.Product, error) {

	products, err := r.db.Product.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldBestDeal)).
		Where(product.HasMerchantWith(merchant.Type("retailer"))).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.GTE(product.FieldBestDeal, 0))
			},
		).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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

func (r *repository) ReadBestSellerByMerchant(id, limit, offset int) ([]*ent.Product, error) {
	products, err := r.db.Product.Query().
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldBestDeal)).
		Where(product.HasMerchantWith(merchant.ID(id))).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.GTE(product.FieldBestDeal, 0))
			},
		).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
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
	fmt.Println(products)
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
