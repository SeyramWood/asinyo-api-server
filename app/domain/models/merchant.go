package models

type (
	RetailMerchantRequestInfo struct {
		MerchantType   string `json:"merchantType" validate:"required|string"`
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:retail_merchants"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:retail_merchants"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	SupplierMerchantRequestInfo struct {
		MerchantType   string `json:"merchantType" validate:"required|string"`
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:supplier_merchants"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:supplier_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:supplier_merchants"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	MerchantRequestCredentials struct {
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:merchants"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
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
		MerchantType   string `json:"merchantType" validate:"required|string"`
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:supplier_merchants"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:supplier_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:supplier_merchants"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
		Username       string `json:"username" validate:"required|email_phone|unique:merchants"`
	}
	SupplierStoreFinalRequest struct {
		Info  SupplierStorePersonalInfo
		Store MerchantStore
	}
	RetailerStorePersonalInfo struct {
		MerchantType   string `json:"merchantType" validate:"required|string"`
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:retail_merchants"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:retail_merchants"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
		Username       string `json:"username" validate:"required|email_phone|unique:merchants"`
	}
	RetailerStoreFinalRequest struct {
		Info  RetailerStorePersonalInfo
		Store MerchantStore
	}
	StoreFinalRequest struct {
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:supplier_merchants"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:supplier_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:supplier_merchants"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
		Username       string `json:"username" validate:"required|email_phone|unique:merchants"`
		MerchantStore
	}
)
