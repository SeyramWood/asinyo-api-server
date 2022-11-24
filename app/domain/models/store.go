package models

type (
	MerchantMomoAccount struct {
		Name     string `json:"name"`
		Number   string `json:"number"`
		Provider string `json:"provider"`
	}
	MerchantBankAccount struct {
		Name   string `json:"name"`
		Number string `json:"number"`
		Bank   string `json:"bank"`
		Branch string `json:"branch"`
	}

	MerchantAccountDefaultRequest struct {
		AccountType string `json:"name" validate:"required|string"`
		Default     bool   `json:"default" validate:"required"`
	}

	MerchantMomoAccountRequest struct {
		AccountName    string `json:"accountName" validate:"required|string"`
		PhoneNumber    string `json:"phoneNumber" validate:"required"`
		Provider       string `json:"provider" validate:"required|string"`
		DefaultAccount bool   `json:"defaultAccount"`
	}
	MerchantBankAccountRequest struct {
		AccountName    string `json:"accountName" validate:"required|string"`
		AccountNumber  string `json:"accountNumber" validate:"required"`
		Bank           string `json:"bank" validate:"required|string"`
		Branch         string `json:"branch" validate:"required|string"`
		DefaultAccount bool   `json:"defaultAccount"`
	}
	MerchantStore struct {
		BusinessName   string `json:"businessName" validate:"required|string"`
		BusinessSlogan string `json:"businessLogan" validate:"required|string"`
		About          string `json:"about" validate:"required"`
		Description    string `json:"description" validate:"required"`
		Banner         []byte `json:"banner"`
		OtherImages    []byte `json:"otherImages"`
		MerchantType   string `json:"merchantType" validate:"required|string"`
	}
	// StoreAddress struct {
	// 	Address      string `json:"postalAddress" validate:"required|string"`
	// 	Country      string `json:"country,omitempty" validate:"required|string"`
	// 	Region       string `json:"region" validate:"required|string"`
	// 	City         string `json:"city" validate:"required|string"`
	// 	District     string `json:"district" validate:"required|string"`
	// 	StreetName   string `json:"streetName" validate:"required|string"`
	// 	StreetNumber string `json:"streetNumber" validate:"string"`
	// }
	MerchantStoreRequest struct {
		Info    *MerchantStore        `json:"postalAddress"`
		Address *MerchantStoreAddress `json:"address"`
	}
)
