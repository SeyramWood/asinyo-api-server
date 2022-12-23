package models

type (
	RetailMerchantRequestInfo struct {
		MerchantType string `json:"merchantType" validate:"required|string"`
		GhanaCard    string `json:"ghanaCard" validate:"required|id_card|unique:retail_merchants"`
		LastName     string `json:"lastName" validate:"required|string"`
		OtherName    string `json:"otherName" validate:"required|string"`
		Phone        string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone   string `json:"otherPhone" validate:"string|unique:retail_merchants"`
	}
	SupplierMerchantRequestInfo struct {
		MerchantType string `json:"merchantType" validate:"required|string"`
		GhanaCard    string `json:"ghanaCard" validate:"required|id_card|unique:supplier_merchants"`
		LastName     string `json:"lastName" validate:"required|string"`
		OtherName    string `json:"otherName" validate:"required|string"`
		Phone        string `json:"phone" validate:"required|string|unique:supplier_merchants"`
		OtherPhone   string `json:"otherPhone" validate:"string|unique:supplier_merchants"`
	}
	MerchantRequestCredentials struct {
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:merchants"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:Password"`
	}

	MerchantRequest struct {
		Info        RetailMerchantRequestInfo
		Credentials MerchantRequestCredentials
	}

	Merchant struct {
		RetailMerchantRequestInfo
		MerchantRequestCredentials
	}

	SupplierStorePersonalInfo struct {
		MerchantType string `json:"merchantType" validate:"required|string"`
		GhanaCard    string `json:"ghanaCard" validate:"required|id_card|unique:supplier_merchants"`
		LastName     string `json:"lastName" validate:"required|string"`
		OtherName    string `json:"otherName" validate:"required|string"`
		Phone        string `json:"phone" validate:"required|phone|unique:supplier_merchants"`
		OtherPhone   string `json:"otherPhone" validate:"phone|unique:supplier_merchants"`
		Username     string `json:"username" validate:"required|email_phone|unique:merchants"`
	}
	RetailerStorePersonalInfo struct {
		MerchantType string `json:"merchantType" validate:"required|string"`
		GhanaCard    string `json:"ghanaCard" validate:"required|id_card|unique:retail_merchants"`
		LastName     string `json:"lastName" validate:"required|string"`
		OtherName    string `json:"otherName" validate:"required|string"`
		Phone        string `json:"phone" validate:"required|phone|unique:retail_merchants"`
		OtherPhone   string `json:"otherPhone" validate:"phone|unique:retail_merchants"`
		Username     string `json:"username" validate:"required|email_phone|unique:merchants"`
	}
	MerchantStoreAddress struct {
		Country        string `json:"country,omitempty" validate:"string"`
		Region         string `json:"region" validate:"required|string"`
		City           string `json:"city" validate:"required|string"`
		District       string `json:"district" validate:"required|string"`
		StreetName     string `json:"streetName" validate:"required"`
		StreetNumber   string `json:"streetNumber" validate:"string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|digital_address"`
	}
	MerchantStoreInfo struct {
		BusinessName   string `json:"businessName" validate:"required|string"`
		BusinessSlogan string `json:"businessLogan" validate:"required|string"`
		About          string `json:"about" validate:"required"`
		Description    string `json:"description" validate:"required"`
		Banner         []byte `json:"banner"`
		OtherImages    []byte `json:"otherImages"`
	}
	OnboardMerchantFullRequest struct {
		PersonalInfo *RetailerStorePersonalInfo `json:"personalInfo"`
		Address      *MerchantStoreAddress      `json:"address"`
		StoreInfo    *MerchantStoreInfo         `json:"storeInfo"`
	}
	SupplierStoreFinalRequest struct {
		Info  SupplierStorePersonalInfo
		Store MerchantStore
	}
	RetailerStoreFinalRequest struct {
		Info  RetailerStorePersonalInfo
		Store MerchantStore
	}
	SupplierProfileUpdate struct {
		LastName   string `json:"lastName" validate:"required|string"`
		OtherName  string `json:"otherName" validate:"required|string"`
		Phone      string `json:"phone" validate:"required|phone"`
		OtherPhone string `json:"otherPhone" validate:"phone"`
	}
	RetailerProfileUpdate struct {
		LastName   string `json:"lastName" validate:"required|string"`
		OtherName  string `json:"otherName" validate:"required|string"`
		Phone      string `json:"phone" validate:"required|phone"`
		OtherPhone string `json:"otherPhone" validate:"phone"`
	}
)
