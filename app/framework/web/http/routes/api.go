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
	db         *database.Adapter
	mail       gateways.EmailService
	logis      gateways.LogisticService
	maps       gateways.MapService
	storageSrv gateways.StorageService
}

func NewApiRouter(
	db *database.Adapter, mail gateways.EmailService, logis gateways.LogisticService, maps gateways.MapService,
	storageSrv gateways.StorageService,
) *ApiRouter {
	return &ApiRouter{db, mail, logis, maps, storageSrv}
}

func (h *ApiRouter) Router(app *fiber.App) {

	r := app.Group("/api") // limiter.New()

	authRouter(r, h.db, h.mail)

	agentRouter(r, h.db, h.storageSrv)

	merchantRouter(r, h.db, h.mail, h.maps, h.storageSrv)

	retailMerchantRouter(r, h.db, h.mail, h.maps)

	supplierMerchantRouter(r, h.db, h.mail, h.maps)

	customerRouter(r, h.db, h.storageSrv)

	productRouter(r, h.db, h.storageSrv)

	paymentRouter(r, h.db, h.logis, h.mail)

	orderRouter(r, h.db, h.logis, h.mail)

	adminRouter(r, h.db)

	addressRouter(r, h.db, h.maps)

}

func customerRouter(r fiber.Router, db *database.Adapter, storageSrv gateways.StorageService) {

	h := handler.NewCustomerHandler(db, storageSrv)

	router := r.Group("/auth/customers")

	authRouter := r.Group("/customers") // middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Post("/sign-up", request.ValidateCustomer(), h.Create()).Name("register")

		}, "auth.customers.",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/", h.Fetch())
			r.Get("/:id", h.FetchByID())
			r.Post("/upload-logo", h.UpdateBusinessLogo())
			r.Put("/update/:id", request.ValidateCustomerUpdate(), h.Update())
			r.Delete("/:id", h.Delete())

		}, "customers.",
	)

}
func agentRouter(r fiber.Router, db *database.Adapter, storageSrv gateways.StorageService) {

	h := handler.NewAgentHandler(db, storageSrv)

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
			r.Get("/:id", h.FetchByID())
			r.Get("/my-merchants/:agent", h.FetchAllMerchant())
			r.Get("/compliance/:agent/get", h.FetchComplianceByID())
			r.Post("/add-compliance/:agent", request.ValidateAgentCompliance(), h.CreateCompliance())
			r.Post("/update/agent/compliance/card", h.UpdateAgentComplianceCard())
			r.Post("/update/agent/compliance/police-report", h.UpdateAgentPoliceReport())
			r.Post("/update/guarantor/compliance/card", h.UpdateGuarantorComplianceCard())
			r.Put("/update/:id/profile", request.ValidateAgentUpdate(), h.Update())
			r.Put("/update-compliance-guarantor/:id", request.ValidateAgentGuarantor(), h.UpdateGuarantor())
			r.Put("/update-account/:id/:accountType", request.ValidateAgentAccount(), h.SaveAccount())
			r.Put("/update-account/:id/default-account/:accountType", h.SaveDefaultAccount())
			r.Delete("/:id", h.Delete()).Name("delete")

		}, "agents.",
	)

}

func retailMerchantRouter(r fiber.Router, db *database.Adapter, mail gateways.EmailService, maps gateways.MapService) {

	h := handler.NewRetailMerchantHandler(db, mail)
	m := handler.NewMerchantHandler(db, mail, maps)

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

func supplierMerchantRouter(
	r fiber.Router, db *database.Adapter, mail gateways.EmailService, maps gateways.MapService,
) {

	h := handler.NewSupplierMerchantHandler(db)
	m := handler.NewMerchantHandler(db, mail, maps)

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

func merchantRouter(
	r fiber.Router, db *database.Adapter, mail gateways.EmailService, maps gateways.MapService,
	storageSrv gateways.StorageService,
) {

	mHandler := handler.NewMerchantHandler(db, mail, maps)
	msHandler := handler.NewMerchantStoreHandler(db, maps, storageSrv)

	mRouter := r.Group("/merchants")
	msRouter := mRouter.Group("/storefront")

	mRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/:id", mHandler.FetchByID())
			r.Put("/update/:id/profile", request.ValidateMerchantProfileUpdate(), mHandler.Update())

		}, "merchants.",
	)

	msRouter.Route(
		"/", func(r fiber.Router) {

			r.Get("/:storeId", msHandler.FetchByID())

			r.Get("/:merchantType/all", msHandler.Fetch())

			r.Get("/:storeId/agent/my-agent", msHandler.FetchMerchantAgent())

			r.Get("/:merchantId/:merchant", msHandler.FetchByMerchantID())

			r.Post("/:merchantId/profile", request.ValidateMerchantStore(), msHandler.Create())
			r.Post("/update/business/banner", msHandler.UpdateBusinessBanner())
			r.Post("/update/business/image", msHandler.UpdateBusinessImages())

			r.Put("/:storeId/profile", request.ValidateMerchantStoreUpdate(), msHandler.Update())

			r.Put("/update-account/:id/:accountType", request.ValidateMerchantAccount(), msHandler.SaveAccount())

			r.Put("/:storeId/update-agent-permission/:permission", msHandler.SaveAgentPermission())

			r.Put("/:storeId/default-account/:type", msHandler.SaveDefaultAccount())

		}, "merchants.",
	)

}

func productRouter(r fiber.Router, db *database.Adapter, storageSrv gateways.StorageService) {

	h := handler.NewProductHandler(db, storageSrv)
	majorH := handler.NewProductCatMajorHandler(db, storageSrv)
	minorH := handler.NewProductCatMinorHandler(db, storageSrv)

	router := r.Group("/products")

	authRouter := r.Group("/product") // middleware.Auth()

	router.Route(
		"/", func(r fiber.Router) {

			r.Get("/", h.Fetch()).Name("all")

			r.Get("/category/major", majorH.Fetch())

			r.Get("/category/minor", minorH.Fetch())

			r.Get("/best-seller/:merchantType", h.FetchBestSellerProducts())

			r.Get("/merchant-best-seller/:id", h.FetchMerchantBestSellerProducts())

			r.Get("/:merchant/:id", h.FetchMerchantProducts())

			r.Get("/category/major/:merchant", h.FetchAllMerchantCategoryMajorProducts())

			r.Get("/category/:merchant/:cat/:slug", h.FetchBySlugMerchantCategoryProducts())

			r.Get("/category/:merchant/:cat/:slug/all", h.FetchAllBySlugMerchantCategoryProducts())

			r.Get("/:merchant/:id/:slug", h.FetchByIDMerchantProduct())

			r.Post("/create/:merchant", request.ValidateProduct(), h.Create())
			r.Post("/update/product-image", h.UpdateImage())
			r.Put("/update/:id", request.ValidateProductUpdate(), h.Update())
			r.Delete("/delete/:id", h.Delete())

		}, "products.",
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Post("/", request.ValidateProduct(), h.Create())
			r.Post("/majors", request.ValidateProductCatMajor(), majorH.Create())
			r.Post("/minors", request.ValidateProductCatMinor(), minorH.Create())
			r.Put("/majors/:id", request.ValidateProductCatMajor(), majorH.Update())
			r.Put("/minors/:id", request.ValidateProductCatMinorUpdate(), minorH.Update())
			r.Post("/minors/update/category-image", minorH.UpdateImage())
			r.Delete("/majors/:id", majorH.Delete())
			r.Delete("/minors/:id", minorH.Delete())

		}, "product.",
	)

}

func paymentRouter(
	router fiber.Router, db *database.Adapter, logis gateways.LogisticService, mail gateways.EmailService,
) {

	ph := handler.NewPaystackHandler(db, logis, mail)
	oh := handler.NewOrderHandler(db, logis, mail)

	router.Route(
		"/payment", func(r fiber.Router) {

			r.Post("/initiate-transaction/paystack", ph.InitiateTransaction())
			r.Get("/verify-transaction/paystack/:reference", ph.VerifyTransaction())
			r.Post("/webhook/paystack", ph.SaveOrder())

			r.Post("/pay-on-delivery", oh.SaveOrder())

		}, "payment.",
	)

}

func orderRouter(
	router fiber.Router, db *database.Adapter, logis gateways.LogisticService, mail gateways.EmailService,
) {

	h := handler.NewOrderHandler(db, logis, mail)

	router.Route(
		"/orders", func(r fiber.Router) {

			r.Get("/", h.Fetch())
			r.Get("/:id", h.FetchById())

			r.Get("/:merchant/store", h.FetchAllByStore())
			r.Get("/agent/store/:agent", h.FetchAllByAgentStore())
			r.Get("/:merchant/store/:order", h.FetchByStore())

			r.Get("/:id/:userType", h.FetchByUser())

			r.Put("/update-order-details-status", h.UpdateOrderDetailStatus())

			r.Post("/get-fare-estimate", h.FetchOrderFareEstimate())
			r.Post("/webhook/update-order-details-status", h.ListenTookanWebhook())
			r.Post("/test/order/creation", h.TestOrderCreation())

		}, "order.",
	)

}

func addressRouter(router fiber.Router, db *database.Adapter, maps gateways.MapService) {

	h := handler.NewAddressHandler(db, maps)
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

			r.Get("/signout", middleware.Auth(), h.Logout())
			r.Get("/user", middleware.Auth(), h.FetchAuthUser())

			r.Post("/signin", request.ValidateUser(), h.Login())
			r.Post("/send-user-verification-code", request.ValidateUserName(true), h.SendVerificationCode())
			r.Post("/send-password-reset-code", request.ValidateUserName(false), h.SendPasswordResetCode())

			r.Put("/change-password/:user/:userType", request.ValidateChangePassword("update"), h.ChangePassword())

			r.Put("/reset-password/:user/:userType", request.ValidateChangePassword("reset"), h.ResetPassword())

		},
	)

}

func adminRouter(r fiber.Router, db *database.Adapter) {

	rph := handler.NewRoleAndPermissionHandler(db)
	router := r.Group("/admins") // middleware.Auth()

	router.Route(
		"/roles", func(r fiber.Router) {
			r.Get("/", rph.Fetch())
			r.Get("/permissions", rph.FetchPermissions())
			r.Post("/", request.ValidateRole(), rph.CreateRole())
			r.Put("/:id", request.ValidateRole(), rph.UpdateRole())
			r.Delete("/:id", rph.DeleteRole())
		},
	)

	h := handler.NewAdminHandler(db)
	router.Route(
		"/", func(r fiber.Router) {
			r.Get("/", h.Fetch())
			r.Get("/:id", h.FetchByID())
			r.Post("/", request.ValidateAdmin(), h.Create())
			r.Put("/:id", request.ValidateAdmin(), h.Update())
			r.Delete("/:id", h.Delete())
		},
	)

}
