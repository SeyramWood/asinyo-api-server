package agent

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
)

type repository struct {
	db *ent.Client
}

func NewAgentRepo(db *database.Adapter) gateways.AgentRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(agent *models.AgentRequest) (*ent.Agent, error) {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(agent.Credentials.Password), 16)

	result, err := r.db.Agent.Create().
		SetLastName(agent.Info.LastName).
		SetOtherName(agent.Info.OtherName).
		SetPhone(agent.Info.Phone).
		SetOtherPhone(agent.Info.OtherPhone).
		SetAddress(agent.Info.Address).
		SetDigitalAddress(agent.Info.DigitalAddress).
		SetGhanaCard(agent.Info.GhanaCard).
		SetUsername(agent.Credentials.Username).
		SetPassword(hashPassword).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed creating agent: %w", err)
	}

	return result, nil
}

func (r *repository) CreateCompliance(
	request *models.AgentComplianceRequest, id int, report string, personal []string, guarantor []string,
) (*ent.Agent, error) {
	ctx := context.Background()
	result, err := r.db.Agent.UpdateOneID(id).
		SetCompliance(
			&models.AgentComplianceModel{
				GhanaCard:    personal,
				PoliceReport: report,
				Guarantor: &models.AgentGuarantorModel{
					GhanaCard:      guarantor,
					LastName:       request.LastName,
					OtherName:      request.OtherName,
					Phone:          request.Phone,
					OtherPhone:     request.OtherPhone,
					Address:        request.Address,
					DigitalAddress: request.DigitalAddress,
					Relation:       request.Relation,
					Occupation:     request.Occupation,
				},
			},
		).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed update merchant momo account : %w", err)
	}
	return result, nil

}

func (r *repository) Read(id int) (*ent.Agent, error) {

	result, err := r.db.Agent.Query().Where(agent.ID(id)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) ReadAll(limit, offset int) (*presenters.PaginationResponse, error) {

	ctx := context.Background()
	query := r.db.Agent.Query()
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := query.
		Limit(limit).Offset(offset).
		Order(ent.Desc(agent.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: count,
		Data:  results,
	}, nil
}

func (r *repository) ReadAllMerchant(agentId int) ([]*ent.MerchantStore, error) {
	results, err := r.db.MerchantStore.Query().
		Where(merchantstore.HasAgentWith(agent.ID(agentId))).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.Select(merchant.FieldID, merchant.FieldType)
				mq.WithSupplier()
				mq.WithRetailer()
			},
		).
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed fetching merchant stores : %w", err)
	}

	return results, nil

}

func (r *repository) Update(id int, profile *models.AgentProfile) (*ent.Agent, error) {
	result, err := r.db.Agent.UpdateOneID(id).
		SetRegion(profile.Region).
		SetDistrict(profile.District).
		SetCity(profile.City).
		SetAddress(profile.Address).
		SetDigitalAddress(profile.DigitalAddress).
		SetLastName(profile.LastName).
		SetOtherName(profile.OtherName).
		SetPhone(profile.Phone).
		SetOtherPhone(profile.OtherPhone).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateGuarantor(id int, request *models.AgentGuarantorUpdate) (*ent.Agent, error) {
	ctx := context.Background()
	old := r.db.Agent.Query().Where(agent.ID(id)).OnlyX(ctx)
	result, err := old.Update().
		SetCompliance(
			&models.AgentComplianceModel{
				GhanaCard:    old.Compliance.GhanaCard,
				PoliceReport: old.Compliance.PoliceReport,
				Guarantor: &models.AgentGuarantorModel{
					GhanaCard:      old.Compliance.Guarantor.GhanaCard,
					LastName:       request.LastName,
					OtherName:      request.OtherName,
					Phone:          request.Phone,
					OtherPhone:     request.OtherPhone,
					Address:        request.Address,
					DigitalAddress: request.DigitalAddress,
					Relation:       request.Relation,
					Occupation:     request.Occupation,
				},
			},
		).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateAgentComplianceCard(agentId int, newPath, oldPath string) ([]string, error) {
	ctx := context.Background()
	old := r.db.Agent.Query().Where(agent.ID(agentId)).OnlyX(ctx)
	newCard := lo.Map[string](
		old.Compliance.GhanaCard, func(path string, index int) string {
			if path == oldPath {
				return newPath
			}
			return path
		},
	)
	_, err := old.Update().
		SetCompliance(
			&models.AgentComplianceModel{
				GhanaCard:    newCard,
				PoliceReport: old.Compliance.PoliceReport,
				Guarantor:    old.Compliance.Guarantor,
			},
		).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newCard, nil
}

func (r *repository) UpdateAgentPoliceReport(agentId int, filePath string) (string, error) {
	ctx := context.Background()
	old := r.db.Agent.Query().Where(agent.ID(agentId)).OnlyX(ctx)
	_, err := old.Update().
		SetCompliance(
			&models.AgentComplianceModel{
				GhanaCard:    old.Compliance.GhanaCard,
				PoliceReport: filePath,
				Guarantor:    old.Compliance.Guarantor,
			},
		).Save(ctx)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (r *repository) UpdateGuarantorComplianceCard(agentId int, newPath, oldPath string) ([]string, error) {
	ctx := context.Background()
	old := r.db.Agent.Query().Where(agent.ID(agentId)).OnlyX(ctx)
	newCard := lo.Map[string](
		old.Compliance.GhanaCard, func(path string, index int) string {
			if path == oldPath {
				return newPath
			}
			return path
		},
	)
	old.Compliance.Guarantor.GhanaCard = newCard
	_, err := old.Update().
		SetCompliance(
			&models.AgentComplianceModel{
				GhanaCard:    old.Compliance.GhanaCard,
				PoliceReport: old.Compliance.PoliceReport,
				Guarantor:    old.Compliance.Guarantor,
			},
		).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newCard, nil
}
func (r *repository) ApproveAgent(agentId int, complianceStatus bool) (*ent.Agent, error) {
	result, err := r.db.Agent.UpdateOneID(agentId).
		SetVerified(complianceStatus).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateAccount(account any, agentId int, accountType string) (*ent.Agent, error) {
	ctx := context.Background()
	if accountType == "bank" {
		request := account.(*models.AgentBankAccountRequest)
		if request.DefaultAccount {
			result, err := r.db.Agent.UpdateOneID(agentId).
				SetDefaultAccount("bank").
				SetBankAccount(
					&models.AgentBankAccount{
						Name:   request.AccountName,
						Number: request.AccountNumber,
						Bank:   request.Bank,
						Branch: request.Branch,
					},
				).
				Save(ctx)
			if err != nil {
				return nil, err
			}
			return result, nil
		}

		result, err := r.db.Agent.UpdateOneID(agentId).
			SetBankAccount(
				&models.AgentBankAccount{
					Name:   request.AccountName,
					Number: request.AccountNumber,
					Bank:   request.Bank,
					Branch: request.Branch,
				},
			).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	request := account.(*models.AgentMomoAccountRequest)
	if request.DefaultAccount {
		result, err := r.db.Agent.UpdateOneID(agentId).
			SetDefaultAccount("momo").
			SetMomoAccount(
				&models.AgentMomoAccount{
					Name:     request.AccountName,
					Number:   request.PhoneNumber,
					Provider: request.Provider,
				},
			).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	result, err := r.db.Agent.UpdateOneID(agentId).
		SetMomoAccount(
			&models.AgentMomoAccount{
				Name:     request.AccountName,
				Number:   request.PhoneNumber,
				Provider: request.Provider,
			},
		).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateDefaultAccount(agentId int, accountType string) (*ent.Agent, error) {
	result, err := r.db.Agent.UpdateOneID(agentId).
		SetDefaultAccount(agent.DefaultAccount(accountType)).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
