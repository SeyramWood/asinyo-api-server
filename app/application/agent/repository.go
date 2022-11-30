package agent

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
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

func (r *repository) Read(id int) (*ent.Agent, error) {

	result, err := r.db.Agent.Query().Where(agent.ID(id)).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) ReadAll() ([]*ent.Agent, error) {

	results, err := r.db.Agent.Query().
		Order(ent.Desc(agent.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
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
func (r *repository) Update(i *models.Agent) (*models.Agent, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}

func (r *repository) CreateCompliance(
	request *models.AgentComplianceRequest, id int, report string, personal []string, guarantor []string,
) (*ent.Agent, error) {
	ctx := context.Background()
	result, err := r.db.Agent.UpdateOneID(id).
		SetRegion(request.Region).
		SetDistrict(request.District).
		SetCity(request.City).
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
