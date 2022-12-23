package models

type (
	AgentInfo struct {
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:agents"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|phone|unique:agents"`
		OtherPhone     string `json:"otherPhone" validate:"string|phone:agents"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	AgentCredentials struct {
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:agents"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:Password"`
	}
	AgentGuarantor struct {
		GhanaCard      []byte `json:"ghanaCard"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|phone"`
		OtherPhone     string `json:"otherPhone" validate:"phone"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
		Relation       string `json:"relation" validate:"required|string"`
		Occupation     string `json:"occupation" validate:"required|string"`
	}
	AgentGuarantorUpdate struct {
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|phone"`
		OtherPhone     string `json:"otherPhone" validate:"phone"`
		Address        string `json:"address" validate:"required"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
		Relation       string `json:"relation" validate:"required|string"`
		Occupation     string `json:"occupation" validate:"required|string"`
	}
	AgentGuarantorModel struct {
		GhanaCard      []string `json:"ghanaCard"`
		LastName       string   `json:"lastName"`
		OtherName      string   `json:"otherName"`
		Phone          string   `json:"phone"`
		OtherPhone     string   `json:"otherPhone"`
		Address        string   `json:"address"`
		DigitalAddress string   `json:"digitalAddress"`
		Relation       string   `json:"relation"`
		Occupation     string   `json:"occupation"`
	}
	AgentCompliance struct {
		GhanaCard    []byte          `json:"ghanaCard"`
		PoliceReport []byte          `json:"policeReport"`
		Guarantor    *AgentGuarantor `json:"guarantor,omitempty"`
	}
	AgentComplianceRequest struct {
		GhanaCardPersonal []byte `json:"ghanaCard"`
		PoliceReport      []byte `json:"policeReport"`
		AgentGuarantor
	}
	AgentComplianceModel struct {
		GhanaCard    []string             `json:"ghanaCard"`
		PoliceReport string               `json:"policeReport"`
		Guarantor    *AgentGuarantorModel `json:"guarantor"`
	}
	AgentRequest struct {
		Info        AgentInfo
		Credentials AgentCredentials
	}
	Agent struct {
		AgentInfo
		AgentCredentials
	}
	AgentProfile struct {
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|phone"`
		OtherPhone     string `json:"otherPhone" validate:"string"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
		City           string `json:"city" validate:"required|string"`
		District       string `json:"district" validate:"required|string"`
		Region         string `json:"region" validate:"required|string"`
	}

	AgentMomoAccount struct {
		Name     string `json:"name"`
		Number   string `json:"number"`
		Provider string `json:"provider"`
	}
	AgentBankAccount struct {
		Name   string `json:"name"`
		Number string `json:"number"`
		Bank   string `json:"bank"`
		Branch string `json:"branch"`
	}

	AgentAccountDefaultRequest struct {
		AccountType string `json:"name" validate:"required|string"`
		Default     bool   `json:"default" validate:"required"`
	}

	AgentMomoAccountRequest struct {
		AccountName    string `json:"accountName" validate:"required|string"`
		PhoneNumber    string `json:"phoneNumber" validate:"required|phone"`
		Provider       string `json:"provider" validate:"required|string"`
		DefaultAccount bool   `json:"defaultAccount"`
	}
	AgentBankAccountRequest struct {
		AccountName    string `json:"accountName" validate:"required|string"`
		AccountNumber  string `json:"accountNumber" validate:"required"`
		Bank           string `json:"bank" validate:"required|string"`
		Branch         string `json:"branch,omitempty" validate:"string"`
		DefaultAccount bool   `json:"defaultAccount"`
	}
)
