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
	result, err := r.db.Product.Create().
		SetMerchantID(prod.Merchant).
		SetMajorID(prod.CategoryMajor).
		SetMinorID(prod.CategoryMinor).
		SetPriceModelID(prod.Model).
		SetQuantity(uint32(prod.Quantity)).
		SetWeight(uint32(prod.Weight)).
		SetUnit(prod.Unit).
		SetName(prod.Name).
		SetPrice(prod.Price).
		SetPromoPrice(prod.PromoPrice).
		SetDescription(prod.Description).
		SetImage(imageUrl).
		Save(context.Background())

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
		WithMajor().
		WithMinor().
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
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
					WithPriceModel().
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
					WithPriceModel().
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
					WithPriceModel().
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
					WithPriceModel().
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
					WithPriceModel().
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
					WithPriceModel().
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
					WithPriceModel().
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

	categories := r.db.ProductCategoryMajor.Query().AllX(context.Background())
	responses := make([]*ent.ProductCategoryMajor, 0, len(categories))
	for _, category := range categories {
		prod, err := r.db.ProductCategoryMajor.Query().
			Where(productcategorymajor.SlugContains(category.Slug)).
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
						WithPriceModel().
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
			Only(context.Background())
		if err != nil {
			return nil, err
		}
		responses = append(responses, prod)
	}
	return responses, nil
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
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
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
		WithPriceModel().
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) Update(id int, request *models.ProductUpdate) (*ent.Product, error) {
	_, err := r.db.Product.UpdateOneID(id).
		SetName(request.Name).
		SetQuantity(uint32(request.Quantity)).
		SetWeight(uint32(request.Weight)).
		SetUnit(request.Unit).
		SetPrice(request.Price).
		SetPromoPrice(request.PromoPrice).
		SetDescription(request.Description).
		SetMajorID(request.CategoryMajor).
		SetMinorID(request.CategoryMinor).
		SetPriceModelID(request.Model).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

func (r *repository) UpdateImage(id int, imagePath string) (string, error) {
	ctx := context.Background()
	_, err := r.db.Product.UpdateOneID(id).SetImage(imagePath).Save(ctx)
	if err != nil {
		return "", err
	}
	return imagePath, nil
}

func (r *repository) Delete(id int) error {
	err := r.db.Product.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
