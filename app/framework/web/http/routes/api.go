package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
	handler "github.com/SeyramWood/app/framework/web/http/handlers/api"
	middleware "github.com/SeyramWood/app/framework/web/http/middlewares"
	request "github.com/SeyramWood/app/framework/web/http/requests"
)

type ApiRouter struct {
	db   *database.Adapter
	mail gateways.EmailService
}

func NewApiRouter(db *database.Adapter, mail gateways.EmailService) *ApiRouter {
	return &ApiRouter{db, mail}
}

func (h *ApiRouter) Router(app *fiber.App) {

	r := app.Group("/api") // limiter.New()

	authRouter(r, h.db, h.mail)

	agentRouter(r, h.db)

	merchantRouter(r, h.db)

	retailMerchantRouter(r, h.db, h.mail)

	supplierMerchantRouter(r, h.db, h.mail)

	customerRouter(r, h.db)

	productRouter(r, h.db)

	paymentRouter(r, h.db)

	orderRouter(r, h.db)

	adminRouter(r, h.db)

	addressRouter(r, h.db)

}

func customerRouter(r fiber.Router, db *database.Adapter) {

	h := handler.NewCustomerHandler(db)

	router := r.Group("/auth/customers")

	authRouter := r.Group("/customers", middleware.Auth())

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/sign-up", request.ValidateCustomer(), h.Create()).Name("register")

		}, "auth.customers.",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/", h.Fetch()).Name("fetch")
			r.Delete("/:id", h.Delete()).Name("delete")

		}, "customers.",
	)

}
func agentRouter(r fiber.Router, db *database.Adapter) {

	h := handler.NewAgentHandler(db)

	router := r.Group("/auth/agents")

	authRouter := r.Group("/agents") // /middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/sign-up", request.ValidateAgent(), h.Create()).Name("register")

		}, "auth.agents.",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Get("/", h.Fetch())
			r.Get("/my-merchants/:agent", h.FetchAllMerchant())
			r.Get("/compliance/:agent/get", h.FetchComplianceByID())
			r.Post("/add-compliance/:agent", request.ValidateAgentCompliance(), h.CreateCompliance())
			r.Delete("/:id", h.Delete()).Name("delete")

		}, "agents.",
	)

}

func retailMerchantRouter(r fiber.Router, db *database.Adapter, mail gateways.EmailService) {

	h := handler.NewRetailMerchantHandler(db, mail)
	m := handler.NewMerchantHandler(db, mail)

	router := r.Group("/auth/merchant/retailers")

	authRouter := r.Group("/merchant/retailers") // middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/sign-up", request.ValidateMerchant(), m.Create())

		}, "auth.merchant.retailers",
	)
	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/add-new-merchant/:agent", request.ValidateNewMerchant(), m.OnboardMerchant())

		}, "auth.merchant.retailers",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/", h.Fetch()).Name("fetch")
			r.Delete("/:id", h.Delete()).Name("delete")

		}, "merchant.retailers.",
	)

}

func supplierMerchantRouter(r fiber.Router, db *database.Adapter, mail gateways.EmailService) {

	h := handler.NewSupplierMerchantHandler(db)
	m := handler.NewMerchantHandler(db, mail)

	router := r.Group("/auth/merchant/suppliers")

	authRouter := r.Group("/merchant/suppliers") // middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/sign-up", request.ValidateMerchant(), m.Create())

		}, "auth.merchant.suppliers",
	)

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/add-new-merchant/:agent", request.ValidateNewMerchant(), m.OnboardMerchant())

		}, "auth.merchant.suppliers",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/", h.Fetch()).Name("fetch")
			r.Delete("/:id", h.Delete()).Name("delete")

		}, "merchant.suppliers.",
	)

}

func merchantRouter(r fiber.Router, db *database.Adapter) {

	// m := handlers.NewMerchantHandler(db)
	msHandler := handler.NewMerchantStoreHandler(db)

	mRouter := r.Group("/merchants")
	msRouter := mRouter.Group("/storefront")

	mRouter.Route(
		"/", func(r fiber.Router) {

			// r.Post("/store", request.ValidateMerchant(), m.Create()).Name("register")

		}, "merchants.",
	)

	msRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/:storeId", msHandler.FetchByID())

			r.Get("/:merchantType/all", msHandler.Fetch())

			r.Get("/:storeId/agent/my-agent", msHandler.FetchMerchantAgent())

			r.Get("/:merchantId/:merchant", msHandler.FetchByMerchantID())

			r.Post("/:merchantId/profile", request.ValidateMerchantStore(), msHandler.Create())

			r.Put("/:storeId/account/momo", request.ValidateMerchantMomoAccount(), msHandler.SaveMomoAccount())

			r.Put("/:storeId/account/bank", request.ValidateMerchantBankAccount(), msHandler.SaveBankAccount())

			r.Put("/:storeId/update-agent-permission/:permission", msHandler.SaveAgentPermission())

			r.Put("/:storeId/default-account/:type", msHandler.SaveDefaultAccount())

		}, "merchants.",
	)

}

func adminRouter(r fiber.Router, db *database.Adapter) {

	h := handler.NewAdminHandler(db)

	router := r.Group("/auth/admins")

	authRouter := r.Group("/admins") // middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/create", request.ValidateAdmin(), h.Create()).Name("register")

		}, "auth.admins.",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Post("/create", request.ValidateAdmin(), h.Create()).Name("register")
			r.Get("/", h.Fetch()).Name("fetch")
			r.Get("/:id", h.FetchByID()).Name("fetch.id")
			r.Delete("/:id", h.Delete()).Name("delete")

		}, "admins.",
	)

}

func productRouter(r fiber.Router, db *database.Adapter) {

	h := handler.NewProductHandler(db)
	majorH := handler.NewProductCatMajorHandler(db)
	minorH := handler.NewProductCatMinorHandler(db)

	router := r.Group("/products")

	authRouter := r.Group("/product") // middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Get("/", h.Fetch()).Name("all")

			r.Get("/category/major", majorH.Fetch()).Name("major")

			r.Get("/category/minor", minorH.Fetch()).Name("minor")

			r.Get("/best-seller/:merchantType", h.FetchBestSellerProducts()).Name("best.sellers")

			r.Get("/:merchant/:id", h.FetchMerchantProducts()).Name("merchant.all")

			r.Get("/category/major/:merchant", h.FetchAllMerchantCategoryMajorProducts())

			r.Get("/category/:merchant/:cat/:slug", h.FetchBySlugMerchantCategoryProducts())

			r.Get("/:merchant/:id/:slug", h.FetchByIDMerchantProduct())

			r.Post("/create/:merchant", request.ValidateProduct(), h.Create()).Name("create")

		}, "products.",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Post("/", request.ValidateProduct(), h.Create()).Name("create")

			// r.Get("/", h.Fetch()).Name("fetch")
			// r.Get("/:id", h.FetchByID()).Name("fetch.id")
			// r.Delete("/:id", h.Delete()).Name("delete")

			r.Post("/majors", request.ValidateProductCatMajor(), majorH.Create()).Name("majors.create")

			r.Post("/minors", request.ValidateProductCatMinor(), minorH.Create()).Name("minors.create")

		}, "product.",
	)

}

func paymentRouter(router fiber.Router, db *database.Adapter) {

	ph := handler.NewPaystackHandler(db)
	oh := handler.NewOrderHandler(db)

	router.Route(
		"/payment", func(r fiber.Router) {

			r.Post("/initiate-transaction/paystack", ph.InitiateTransaction())
			r.Get("/verify-transaction/paystack/:reference", ph.VerifyTransaction())
			r.Post("/webhook/paystack", ph.SaveOrder())

			r.Post("/pay-on-delivery", oh.SaveOrder())

		}, "payment.",
	)

}

func orderRouter(router fiber.Router, db *database.Adapter) {

	h := handler.NewOrderHandler(db)

	router.Route(
		"/orders", func(r fiber.Router) {

			r.Get("/:id", h.FetchById())
			r.Get("/:merchant/store", h.FetchAllByStore())
			r.Get("/agent/store/:agent", h.FetchAllByAgentStore())
			r.Get("/:merchant/store/:order", h.FetchByStore())

			r.Get("/:id/:userType", h.FetchByUser())

			r.Put("/update-order-details-status", h.UpdateOrderDetailStatus())

		}, "order.",
	)

}

func addressRouter(router fiber.Router, db *database.Adapter) {

	h := handler.NewAddressHandler(db)
	psh := handler.NewPickupStationHandler(db)

	router.Route(
		"/address", func(r fiber.Router) {

			r.Get("/all/:userId/:userType", h.FetchAllByUser())

			r.Get("/pickup-station/all", psh.Fetch())

			r.Get("/default/:userId/:userType", h.FetchByUser())

			r.Post("/create/:userId/:userType", request.ValidateAddress(), h.Create())

			r.Put("/update/:addressId", request.ValidateAddress(), h.Update())

			r.Put("/set-default/:userId/:userType/:addressId", h.SaveDefaultAddress())

		}, "payment.",
	)

}

func authRouter(router fiber.Router, db *database.Adapter, mail gateways.EmailService) {

	h := handler.NewAuthHandler(db, mail)

	router.Route(
		"/auth", func(r fiber.Router) {

			r.Get("/signout", middleware.Auth(), h.Logout()).Name("signout")
			r.Get("/user", middleware.Auth(), h.FetchAuthUser()).Name("user")

			r.Post("/signin", request.ValidateUser(), h.Login()).Name("signin")
			r.Post("/send-user-verification-code", request.ValidateUserName(), h.SendVerificationCode())

			r.Put("/change-password/:user/:userType", request.ValidateChangePassword(), h.ChangePassword())

		}, "auth.",
	)

}
