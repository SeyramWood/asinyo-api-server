package presenters

import (
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

type AddressResponse struct {
	ID           int       `json:"id"`
	LastName     string    `json:"lastName"`
	OtherName    string    `json:"otherName"`
	Phone        string    `json:"phone"`
	OtherPhone   string    `json:"otherPhone"`
	Address      string    `json:"address"`
	OtherAddress string    `json:"otherAddress"`
	Region       string    `json:"region"`
	City         string    `json:"city"`
	Default      bool      `json:"default"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func AddressSuccessResponse(data *ent.Address) *fiber.Map {
	//if data.ID == true {
	//	return successResponse(&AddressResponse{
	//		ID:           data.ID,
	//		LastName:     data.LastName,
	//		OtherName:    data.OtherName,
	//		Phone:        data.Phone,
	//		OtherPhone:   *data.OtherPhone,
	//		Address:      data.Address,
	//		OtherAddress: *data.OtherInformation,
	//		Region:       data.Region,
	//		City:         data.City,
	//		Default:      data.Default,
	//		CreatedAt:    data.CreatedAt,
	//		UpdatedAt:    data.UpdatedAt,
	//	})
	//}
	//return nil
	return successResponse(&AddressResponse{
		ID:           data.ID,
		LastName:     data.LastName,
		OtherName:    data.OtherName,
		Phone:        data.Phone,
		OtherPhone:   *data.OtherPhone,
		Address:      data.Address,
		OtherAddress: *data.OtherInformation,
		Region:       data.Region,
		City:         data.City,
		Default:      data.Default,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	})
}

func AddressSuccessResponses(data []*ent.Address) *fiber.Map {
	var response []*AddressResponse
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Address) {
			defer wg.Done()
			response = append(response, &AddressResponse{
				ID:           v.ID,
				LastName:     v.LastName,
				OtherName:    v.OtherName,
				Phone:        v.Phone,
				OtherPhone:   *v.OtherPhone,
				Address:      v.Address,
				OtherAddress: *v.OtherInformation,
				Region:       v.Region,
				City:         v.City,
				Default:      v.Default,
				CreatedAt:    v.CreatedAt,
				UpdatedAt:    v.UpdatedAt,
			})
		}(v)
	}
	wg.Wait()

	return successResponse(response)
}

func AddressErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
