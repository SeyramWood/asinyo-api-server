// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/SeyramWood/ent/admin"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategorymajor"
	"github.com/SeyramWood/ent/productcategoryminor"
	"github.com/SeyramWood/ent/retailmerchant"
	"github.com/SeyramWood/ent/schema"
	"github.com/SeyramWood/ent/suppliermerchant"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreatedAt is the schema descriptor for created_at field.
	adminDescCreatedAt := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreatedAt holds the default value on creation for the created_at field.
	admin.DefaultCreatedAt = adminDescCreatedAt.Default.(func() time.Time)
	// adminDescUpdatedAt is the schema descriptor for updated_at field.
	adminDescUpdatedAt := adminMixinFields0[1].Descriptor()
	// admin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	admin.DefaultUpdatedAt = adminDescUpdatedAt.Default.(func() time.Time)
	// admin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	admin.UpdateDefaultUpdatedAt = adminDescUpdatedAt.UpdateDefault.(func() time.Time)
	// adminDescUsername is the schema descriptor for username field.
	adminDescUsername := adminFields[0].Descriptor()
	// admin.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	admin.UsernameValidator = adminDescUsername.Validators[0].(func(string) error)
	// adminDescPassword is the schema descriptor for password field.
	adminDescPassword := adminFields[1].Descriptor()
	// admin.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	admin.PasswordValidator = adminDescPassword.Validators[0].(func([]byte) error)
	agentMixin := schema.Agent{}.Mixin()
	agentMixinFields0 := agentMixin[0].Fields()
	_ = agentMixinFields0
	agentFields := schema.Agent{}.Fields()
	_ = agentFields
	// agentDescCreatedAt is the schema descriptor for created_at field.
	agentDescCreatedAt := agentMixinFields0[0].Descriptor()
	// agent.DefaultCreatedAt holds the default value on creation for the created_at field.
	agent.DefaultCreatedAt = agentDescCreatedAt.Default.(func() time.Time)
	// agentDescUpdatedAt is the schema descriptor for updated_at field.
	agentDescUpdatedAt := agentMixinFields0[1].Descriptor()
	// agent.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	agent.DefaultUpdatedAt = agentDescUpdatedAt.Default.(func() time.Time)
	// agent.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	agent.UpdateDefaultUpdatedAt = agentDescUpdatedAt.UpdateDefault.(func() time.Time)
	// agentDescUsername is the schema descriptor for username field.
	agentDescUsername := agentFields[0].Descriptor()
	// agent.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	agent.UsernameValidator = agentDescUsername.Validators[0].(func(string) error)
	// agentDescPassword is the schema descriptor for password field.
	agentDescPassword := agentFields[1].Descriptor()
	// agent.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	agent.PasswordValidator = agentDescPassword.Validators[0].(func([]byte) error)
	// agentDescGhanaCard is the schema descriptor for ghana_card field.
	agentDescGhanaCard := agentFields[2].Descriptor()
	// agent.GhanaCardValidator is a validator for the "ghana_card" field. It is called by the builders before save.
	agent.GhanaCardValidator = agentDescGhanaCard.Validators[0].(func(string) error)
	// agentDescLastName is the schema descriptor for last_name field.
	agentDescLastName := agentFields[3].Descriptor()
	// agent.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	agent.LastNameValidator = agentDescLastName.Validators[0].(func(string) error)
	// agentDescOtherName is the schema descriptor for other_name field.
	agentDescOtherName := agentFields[4].Descriptor()
	// agent.OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	agent.OtherNameValidator = agentDescOtherName.Validators[0].(func(string) error)
	// agentDescPhone is the schema descriptor for phone field.
	agentDescPhone := agentFields[5].Descriptor()
	// agent.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	agent.PhoneValidator = agentDescPhone.Validators[0].(func(string) error)
	// agentDescAddress is the schema descriptor for address field.
	agentDescAddress := agentFields[7].Descriptor()
	// agent.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	agent.AddressValidator = agentDescAddress.Validators[0].(func(string) error)
	// agentDescDigitalAddress is the schema descriptor for digital_address field.
	agentDescDigitalAddress := agentFields[8].Descriptor()
	// agent.DigitalAddressValidator is a validator for the "digital_address" field. It is called by the builders before save.
	agent.DigitalAddressValidator = agentDescDigitalAddress.Validators[0].(func(string) error)
	customerMixin := schema.Customer{}.Mixin()
	customerMixinFields0 := customerMixin[0].Fields()
	_ = customerMixinFields0
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCreatedAt is the schema descriptor for created_at field.
	customerDescCreatedAt := customerMixinFields0[0].Descriptor()
	// customer.DefaultCreatedAt holds the default value on creation for the created_at field.
	customer.DefaultCreatedAt = customerDescCreatedAt.Default.(func() time.Time)
	// customerDescUpdatedAt is the schema descriptor for updated_at field.
	customerDescUpdatedAt := customerMixinFields0[1].Descriptor()
	// customer.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	customer.DefaultUpdatedAt = customerDescUpdatedAt.Default.(func() time.Time)
	// customer.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	customer.UpdateDefaultUpdatedAt = customerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// customerDescUsername is the schema descriptor for username field.
	customerDescUsername := customerFields[0].Descriptor()
	// customer.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	customer.UsernameValidator = customerDescUsername.Validators[0].(func(string) error)
	// customerDescPassword is the schema descriptor for password field.
	customerDescPassword := customerFields[1].Descriptor()
	// customer.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	customer.PasswordValidator = customerDescPassword.Validators[0].(func([]byte) error)
	// customerDescFirstName is the schema descriptor for first_name field.
	customerDescFirstName := customerFields[2].Descriptor()
	// customer.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	customer.FirstNameValidator = customerDescFirstName.Validators[0].(func(string) error)
	// customerDescLastName is the schema descriptor for last_name field.
	customerDescLastName := customerFields[3].Descriptor()
	// customer.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	customer.LastNameValidator = customerDescLastName.Validators[0].(func(string) error)
	// customerDescPhone is the schema descriptor for phone field.
	customerDescPhone := customerFields[4].Descriptor()
	// customer.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	customer.PhoneValidator = customerDescPhone.Validators[0].(func(string) error)
	merchantMixin := schema.Merchant{}.Mixin()
	merchantMixinFields0 := merchantMixin[0].Fields()
	_ = merchantMixinFields0
	merchantFields := schema.Merchant{}.Fields()
	_ = merchantFields
	// merchantDescCreatedAt is the schema descriptor for created_at field.
	merchantDescCreatedAt := merchantMixinFields0[0].Descriptor()
	// merchant.DefaultCreatedAt holds the default value on creation for the created_at field.
	merchant.DefaultCreatedAt = merchantDescCreatedAt.Default.(func() time.Time)
	// merchantDescUpdatedAt is the schema descriptor for updated_at field.
	merchantDescUpdatedAt := merchantMixinFields0[1].Descriptor()
	// merchant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	merchant.DefaultUpdatedAt = merchantDescUpdatedAt.Default.(func() time.Time)
	// merchant.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	merchant.UpdateDefaultUpdatedAt = merchantDescUpdatedAt.UpdateDefault.(func() time.Time)
	// merchantDescUsername is the schema descriptor for username field.
	merchantDescUsername := merchantFields[0].Descriptor()
	// merchant.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	merchant.UsernameValidator = merchantDescUsername.Validators[0].(func(string) error)
	// merchantDescPassword is the schema descriptor for password field.
	merchantDescPassword := merchantFields[1].Descriptor()
	// merchant.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	merchant.PasswordValidator = merchantDescPassword.Validators[0].(func([]byte) error)
	productMixin := schema.Product{}.Mixin()
	productMixinFields0 := productMixin[0].Fields()
	_ = productMixinFields0
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productMixinFields0[0].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productMixinFields0[1].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productDescName is the schema descriptor for Name field.
	productDescName := productFields[0].Descriptor()
	// product.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescPrice is the schema descriptor for Price field.
	productDescPrice := productFields[1].Descriptor()
	// product.DefaultPrice holds the default value on creation for the Price field.
	product.DefaultPrice = productDescPrice.Default.(float64)
	// productDescPromoPrice is the schema descriptor for PromoPrice field.
	productDescPromoPrice := productFields[2].Descriptor()
	// product.DefaultPromoPrice holds the default value on creation for the PromoPrice field.
	product.DefaultPromoPrice = productDescPromoPrice.Default.(float64)
	// productDescDescription is the schema descriptor for Description field.
	productDescDescription := productFields[3].Descriptor()
	// product.DescriptionValidator is a validator for the "Description" field. It is called by the builders before save.
	product.DescriptionValidator = productDescDescription.Validators[0].(func(string) error)
	// productDescImage is the schema descriptor for Image field.
	productDescImage := productFields[4].Descriptor()
	// product.ImageValidator is a validator for the "Image" field. It is called by the builders before save.
	product.ImageValidator = productDescImage.Validators[0].(func(string) error)
	productcategorymajorMixin := schema.ProductCategoryMajor{}.Mixin()
	productcategorymajorMixinFields0 := productcategorymajorMixin[0].Fields()
	_ = productcategorymajorMixinFields0
	productcategorymajorFields := schema.ProductCategoryMajor{}.Fields()
	_ = productcategorymajorFields
	// productcategorymajorDescCreatedAt is the schema descriptor for created_at field.
	productcategorymajorDescCreatedAt := productcategorymajorMixinFields0[0].Descriptor()
	// productcategorymajor.DefaultCreatedAt holds the default value on creation for the created_at field.
	productcategorymajor.DefaultCreatedAt = productcategorymajorDescCreatedAt.Default.(func() time.Time)
	// productcategorymajorDescUpdatedAt is the schema descriptor for updated_at field.
	productcategorymajorDescUpdatedAt := productcategorymajorMixinFields0[1].Descriptor()
	// productcategorymajor.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productcategorymajor.DefaultUpdatedAt = productcategorymajorDescUpdatedAt.Default.(func() time.Time)
	// productcategorymajor.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productcategorymajor.UpdateDefaultUpdatedAt = productcategorymajorDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productcategorymajorDescCategory is the schema descriptor for category field.
	productcategorymajorDescCategory := productcategorymajorFields[0].Descriptor()
	// productcategorymajor.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	productcategorymajor.CategoryValidator = productcategorymajorDescCategory.Validators[0].(func(string) error)
	productcategoryminorMixin := schema.ProductCategoryMinor{}.Mixin()
	productcategoryminorMixinFields0 := productcategoryminorMixin[0].Fields()
	_ = productcategoryminorMixinFields0
	productcategoryminorFields := schema.ProductCategoryMinor{}.Fields()
	_ = productcategoryminorFields
	// productcategoryminorDescCreatedAt is the schema descriptor for created_at field.
	productcategoryminorDescCreatedAt := productcategoryminorMixinFields0[0].Descriptor()
	// productcategoryminor.DefaultCreatedAt holds the default value on creation for the created_at field.
	productcategoryminor.DefaultCreatedAt = productcategoryminorDescCreatedAt.Default.(func() time.Time)
	// productcategoryminorDescUpdatedAt is the schema descriptor for updated_at field.
	productcategoryminorDescUpdatedAt := productcategoryminorMixinFields0[1].Descriptor()
	// productcategoryminor.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	productcategoryminor.DefaultUpdatedAt = productcategoryminorDescUpdatedAt.Default.(func() time.Time)
	// productcategoryminor.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	productcategoryminor.UpdateDefaultUpdatedAt = productcategoryminorDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productcategoryminorDescCategory is the schema descriptor for category field.
	productcategoryminorDescCategory := productcategoryminorFields[0].Descriptor()
	// productcategoryminor.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	productcategoryminor.CategoryValidator = productcategoryminorDescCategory.Validators[0].(func(string) error)
	retailmerchantMixin := schema.RetailMerchant{}.Mixin()
	retailmerchantMixinFields0 := retailmerchantMixin[0].Fields()
	_ = retailmerchantMixinFields0
	retailmerchantFields := schema.RetailMerchant{}.Fields()
	_ = retailmerchantFields
	// retailmerchantDescCreatedAt is the schema descriptor for created_at field.
	retailmerchantDescCreatedAt := retailmerchantMixinFields0[0].Descriptor()
	// retailmerchant.DefaultCreatedAt holds the default value on creation for the created_at field.
	retailmerchant.DefaultCreatedAt = retailmerchantDescCreatedAt.Default.(func() time.Time)
	// retailmerchantDescUpdatedAt is the schema descriptor for updated_at field.
	retailmerchantDescUpdatedAt := retailmerchantMixinFields0[1].Descriptor()
	// retailmerchant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	retailmerchant.DefaultUpdatedAt = retailmerchantDescUpdatedAt.Default.(func() time.Time)
	// retailmerchant.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	retailmerchant.UpdateDefaultUpdatedAt = retailmerchantDescUpdatedAt.UpdateDefault.(func() time.Time)
	// retailmerchantDescGhanaCard is the schema descriptor for ghana_card field.
	retailmerchantDescGhanaCard := retailmerchantFields[0].Descriptor()
	// retailmerchant.GhanaCardValidator is a validator for the "ghana_card" field. It is called by the builders before save.
	retailmerchant.GhanaCardValidator = retailmerchantDescGhanaCard.Validators[0].(func(string) error)
	// retailmerchantDescLastName is the schema descriptor for last_name field.
	retailmerchantDescLastName := retailmerchantFields[1].Descriptor()
	// retailmerchant.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	retailmerchant.LastNameValidator = retailmerchantDescLastName.Validators[0].(func(string) error)
	// retailmerchantDescOtherName is the schema descriptor for other_name field.
	retailmerchantDescOtherName := retailmerchantFields[2].Descriptor()
	// retailmerchant.OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	retailmerchant.OtherNameValidator = retailmerchantDescOtherName.Validators[0].(func(string) error)
	// retailmerchantDescPhone is the schema descriptor for phone field.
	retailmerchantDescPhone := retailmerchantFields[3].Descriptor()
	// retailmerchant.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	retailmerchant.PhoneValidator = retailmerchantDescPhone.Validators[0].(func(string) error)
	// retailmerchantDescAddress is the schema descriptor for address field.
	retailmerchantDescAddress := retailmerchantFields[5].Descriptor()
	// retailmerchant.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	retailmerchant.AddressValidator = retailmerchantDescAddress.Validators[0].(func(string) error)
	// retailmerchantDescDigitalAddress is the schema descriptor for digital_address field.
	retailmerchantDescDigitalAddress := retailmerchantFields[6].Descriptor()
	// retailmerchant.DigitalAddressValidator is a validator for the "digital_address" field. It is called by the builders before save.
	retailmerchant.DigitalAddressValidator = retailmerchantDescDigitalAddress.Validators[0].(func(string) error)
	suppliermerchantMixin := schema.SupplierMerchant{}.Mixin()
	suppliermerchantMixinFields0 := suppliermerchantMixin[0].Fields()
	_ = suppliermerchantMixinFields0
	suppliermerchantFields := schema.SupplierMerchant{}.Fields()
	_ = suppliermerchantFields
	// suppliermerchantDescCreatedAt is the schema descriptor for created_at field.
	suppliermerchantDescCreatedAt := suppliermerchantMixinFields0[0].Descriptor()
	// suppliermerchant.DefaultCreatedAt holds the default value on creation for the created_at field.
	suppliermerchant.DefaultCreatedAt = suppliermerchantDescCreatedAt.Default.(func() time.Time)
	// suppliermerchantDescUpdatedAt is the schema descriptor for updated_at field.
	suppliermerchantDescUpdatedAt := suppliermerchantMixinFields0[1].Descriptor()
	// suppliermerchant.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	suppliermerchant.DefaultUpdatedAt = suppliermerchantDescUpdatedAt.Default.(func() time.Time)
	// suppliermerchant.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	suppliermerchant.UpdateDefaultUpdatedAt = suppliermerchantDescUpdatedAt.UpdateDefault.(func() time.Time)
	// suppliermerchantDescGhanaCard is the schema descriptor for ghana_card field.
	suppliermerchantDescGhanaCard := suppliermerchantFields[0].Descriptor()
	// suppliermerchant.GhanaCardValidator is a validator for the "ghana_card" field. It is called by the builders before save.
	suppliermerchant.GhanaCardValidator = suppliermerchantDescGhanaCard.Validators[0].(func(string) error)
	// suppliermerchantDescLastName is the schema descriptor for last_name field.
	suppliermerchantDescLastName := suppliermerchantFields[1].Descriptor()
	// suppliermerchant.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	suppliermerchant.LastNameValidator = suppliermerchantDescLastName.Validators[0].(func(string) error)
	// suppliermerchantDescOtherName is the schema descriptor for other_name field.
	suppliermerchantDescOtherName := suppliermerchantFields[2].Descriptor()
	// suppliermerchant.OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	suppliermerchant.OtherNameValidator = suppliermerchantDescOtherName.Validators[0].(func(string) error)
	// suppliermerchantDescPhone is the schema descriptor for phone field.
	suppliermerchantDescPhone := suppliermerchantFields[3].Descriptor()
	// suppliermerchant.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	suppliermerchant.PhoneValidator = suppliermerchantDescPhone.Validators[0].(func(string) error)
	// suppliermerchantDescAddress is the schema descriptor for address field.
	suppliermerchantDescAddress := suppliermerchantFields[5].Descriptor()
	// suppliermerchant.AddressValidator is a validator for the "address" field. It is called by the builders before save.
	suppliermerchant.AddressValidator = suppliermerchantDescAddress.Validators[0].(func(string) error)
	// suppliermerchantDescDigitalAddress is the schema descriptor for digital_address field.
	suppliermerchantDescDigitalAddress := suppliermerchantFields[6].Descriptor()
	// suppliermerchant.DigitalAddressValidator is a validator for the "digital_address" field. It is called by the builders before save.
	suppliermerchant.DigitalAddressValidator = suppliermerchantDescDigitalAddress.Validators[0].(func(string) error)
}