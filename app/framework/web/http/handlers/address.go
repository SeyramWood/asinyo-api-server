package handlers

import (
	"errors"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/address"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	service gateways.AddressService
}

func NewAddressHandler(db *database.Adapter) *AddressHandler {
	repo := address.NewAddressRepo(db)
	service := address.NewAddressService(repo)
	return &AddressHandler{
		service: service,
	}
}

func (h *AddressHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//storeId, _ := c.ParamsInt("storeId")
		//
		//result, err := h.service.Fetch(storeId)
		//
		//if err != nil {
		//	return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantStoreErrorResponse(err))
		//}
		//return c.JSON(presenters.MerchantStoreSuccessResponse(result))
		return nil

	}
}
func (h *AddressHandler) FetchByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		result, err := h.service.FetchByUser(userId, userType)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AddressErrorResponse(err))
		}
		return c.JSON(presenters.AddressSuccessResponse(result))

	}
}
func (h *AddressHandler) FetchAllByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		results, err := h.service.FetchAllByUser(userId, userType)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AddressErrorResponse(err))
		}
		return c.JSON(presenters.AddressSuccessResponses(results))

	}
}

func (h *AddressHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//merchantType := c.Params("merchantType")
		//
		//result, err := h.service.FetchAllByMerchant(merchantType)
		//
		//if err != nil {
		//
		//	return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		//}
		//return c.JSON(presenters.MerchantStorefrontsSuccessResponse(result))
		return nil
	}
}

func (h *AddressHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.Address
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		result, err := h.service.Create(&request, userId, userType)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AddressErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.AddressSuccessResponse(result))

	}
}

func (h *AddressHandler) SaveDefaultAddress() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		addressId, _ := c.ParamsInt("addressId")

		results, err := h.service.SaveDefaultAddress(userId, addressId, userType)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AddressErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.AddressSuccessResponses(results))

	}
}

func (h *AddressHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Address

		addressId, _ := c.ParamsInt("addressId")

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		result, err := h.service.Update(addressId, &request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AddressErrorResponse(errors.New("error updating address")))
		}
		return c.JSON(presenters.AddressSuccessResponse(result))
	}

}
func (h *AddressHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//if err := h.service.Remove(c.Params("id")); err != nil {
		//	return c.Status(fiber.StatusNotFound).JSON(presenters.MerchantErrorResponse(err))
		//}
		//return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		//	"status": true,
		//	"error":  nil,
		//})
		return nil
	}
}
