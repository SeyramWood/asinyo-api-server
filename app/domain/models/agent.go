package models

type (
	AgentInfo struct {
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card|unique:agents"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:agents"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:agents"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	AgentCredentials struct {
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:agents"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}
	AgentGuarantor struct {
		GhanaCard      []byte `json:"ghanaCard"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string"`
		OtherPhone     string `json:"otherPhone" validate:"string"`
		Address        string `json:"address" validate:"required|string"`
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
		Region       string          `json:"region" validate:"required|string"`
		District     string          `json:"district" validate:"required|string"`
		City         string          `json:"city" validate:"required|string"`
		Guarantor    *AgentGuarantor `json:"guarantor"`
	}
	AgentComplianceRequest struct {
		GhanaCardPersonal []byte `json:"ghanaCard"`
		PoliceReport      []byte `json:"policeReport"`
		Region            string `json:"region" validate:"required|string"`
		District          string `json:"district" validate:"required|string"`
		City              string `json:"city" validate:"required|string"`
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
)
