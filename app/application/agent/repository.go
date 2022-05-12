package agent

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"golang.org/x/crypto/bcrypt"
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
		return nil, fmt.Errorf("failed creating agent: %w", err)
	}

	return result, nil
}

func (r *repository) Read(id int) (*ent.Agent, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.Agent, error) {

	// b, err := r.db.User.Query().
	// 	All(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (a *repository) Update(i *models.Agent) (*models.Agent, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

//DeleteBook is a mongo repository that helps to delete books
func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
