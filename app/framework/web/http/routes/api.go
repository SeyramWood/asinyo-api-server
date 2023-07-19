package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/framework/database"
	handler "github.com/SeyramWood/app/framework/web/http/handlers/api"
	ware "github.com/SeyramWood/app/framework/web/http/middlewares"
	request "github.com/SeyramWood/app/framework/web/http/requests"
	"github.com/SeyramWood/pkg/app"
)

type ApiRouter struct {
	app        *app.Server
	middleware *ware.Middleware
	db         *database.Adapter
	noti       notification.NotificationService
	dbNoti     gateways.DBNotificationService
	storageSrv gateways.StorageService
	logis      gateways.LogisticService
	ms         gateways.MapService
}

func NewApiRouter(
	app *app.Server,
	db *database.Adapter,
	noti notification.NotificationService,
	dbNoti gateways.DBNotificationService,
	storageSrv gateways.StorageService,
	logis gateways.LogisticService,
	ms gateways.MapService,
) *ApiRouter {
	newWare := ware.NewMiddleware(db, noti, app.JWT, app.Cache)
	return &ApiRouter{
		app:        app,
		middleware: newWare,
		db:         db,
		noti:       noti,
		dbNoti:     dbNoti,
		storageSrv: storageSrv,
		logis:      logis,
		ms:         ms,
	}
}

func (ar *ApiRouter) Router(app *fiber.App) {
	r := app.Group("/api") // limiter.New()

	authRouter(r, ar)

	agentRouter(r, ar)

	merchantRouter(r, ar)

	retailMerchantRouter(r, ar)

	supplierMerchantRouter(r, ar)

	customerRouter(r, ar)

	productRouter(r, ar)

	paymentRouter(r, ar)

	orderRouter(r, ar)

	adminRouter(r, ar)

	addressRouter(r, ar)

	notificationRouter(r, ar)

	pricemodelRouter(r, ar)
}

func customerRouter(r fiber.Router, ar *ApiRouter) {
	h := handler.NewCustomerHandler(ar.db, ar.storageSrv)

	router := r.Group("/auth/customers")
	router.Route(
		"/", func(r fiber.Router) {
			r.Post("/sign-up", ar.middleware.IsAuthorized(), request.ValidateCustomer(), h.Create())
		}, "auth.customers.",
	)

	authRouter := r.Group("/customers")
	authRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", h.Fetch())
			r.Get("/:id", h.FetchByID())
			r.Get("/purchase-request/:id", h.FetchPurchaseRequestByID())
			r.Get("/:customer/purchase-request", h.FetchAllPurchaseRequestByCustomer())
			r.Post("/upload-logo", h.UpdateBusinessLogo())
			r.Post("/:customer/purchase-request", request.ValidateCustomerPurchaseRequest(), h.CreatePurchaseRequest())
			r.Put("/purchase-request/:id", request.ValidateCustomerPurchaseRequest(), h.UpdatePurchaseRequest())
			r.Put("/update/:id", request.ValidateCustomerUpdate(), h.Update())
			r.Delete("/:id", h.Delete())
		},
	)
}

func agentRouter(r fiber.Router, ar *ApiRouter) {
	h := handler.NewAgentHandler(ar.db, ar.storageSrv)

	router := r.Group("/auth/agents")

	authRouter := r.Group("/agents")

	router.Route(
		"/", func(r fiber.Router) {
			r.Post("/sign-up", ar.middleware.IsAuthorized(), request.ValidateAgent(), h.Create())
		},
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Get("/", h.Fetch())
			r.Get("/paginate", h.FetchWithPaginate())
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
			r.Put("/compliance/:agent/:status", h.ApproveAgent())
			r.Delete("/:id", h.Delete()).Name("delete")
		}, "agents.",
	)
}

func retailMerchantRouter(r fiber.Router, ar *ApiRouter) {
	h := handler.NewRetailMerchantHandler(ar.db, ar.noti)
	m := handler.NewMerchantHandler(ar.db, ar.noti, ar.storageSrv, ar.ms)

	router := r.Group("/auth/merchant/retailers")

	authRouter := r.Group("/merchant/retailers") // middleware.IsAuthorized()

	router.Route(
		"/", func(r fiber.Router) {
			r.Post("/sign-up", ar.middleware.IsAuthorized(), request.ValidateMerchant(), m.Create())
		},
	)
	router.Route(
		"/", func(r fiber.Router) {
			r.Post("/add-new-merchant/:agent", request.ValidateNewMerchant(), m.OnboardMerchant())
		},
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Get("/", h.Fetch()).Name("fetch")
			r.Delete("/:id", h.Delete()).Name("delete")
		},
	)
}

func supplierMerchantRouter(r fiber.Router, ar *ApiRouter) {
	h := handler.NewSupplierMerchantHandler(ar.db)
	m := handler.NewMerchantHandler(ar.db, ar.noti, ar.storageSrv, ar.ms)

	router := r.Group("/auth/merchant/suppliers")

	authRouter := r.Group("/merchant/suppliers") // middleware.IsAuthorized()

	router.Route(
		"/", func(r fiber.Router) {
			r.Post("/sign-up", ar.middleware.IsAuthorized(), request.ValidateMerchant(), m.Create())
		},
	)

	router.Route(
		"/", func(r fiber.Router) {
			r.Post("/add-new-merchant/:agent", request.ValidateNewMerchant(), m.OnboardMerchant())
		},
	)

	authRouter.Route(
		"/", func(r fiber.Router) {
			r.Get("/", h.Fetch()).Name("fetch")
			r.Delete("/:id", h.Delete()).Name("delete")
		},
	)
}

func merchantRouter(r fiber.Router, ar *ApiRouter) {
	mHandler := handler.NewMerchantHandler(ar.db, ar.noti, ar.storageSrv, ar.ms)
	msHandler := handler.NewMerchantStoreHandler(ar.db, ar.ms, ar.storageSrv)

	mRouter := r.Group("/merchants")
	msRouter := mRouter.Group("/storefront")

	mRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", mHandler.Fetch())
			r.Get("/:id", mHandler.FetchByID())
			r.Get("/:id/storefront", mHandler.FetchStorefrontByID())
			r.Put("/update/:id/profile", request.ValidateMerchantProfileUpdate(), mHandler.Update())
		},
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
		},
	)
}

func productRouter(r fiber.Router, ar *ApiRouter) {
	h := handler.NewProductHandler(ar.db, ar.storageSrv)
	majorH := handler.NewProductCatMajorHandler(ar.db, ar.storageSrv)
	minorH := handler.NewProductCatMinorHandler(ar.db, ar.storageSrv)

	router := r.Group("/products")

	authRouter := r.Group("/product")

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
		},
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

func paymentRouter(router fiber.Router, ar *ApiRouter) {
	ph := handler.NewPaystackHandler(ar.db, ar.logis, ar.noti)
	oh := handler.NewOrderHandler(ar.db, ar.logis, ar.noti)

	router.Route(
		"/payment", func(r fiber.Router) {
			r.Post("/initiate-transaction/paystack", ph.InitiateTransaction())
			r.Get("/verify-transaction/paystack/:reference", ph.VerifyTransaction())
			r.Post("/webhook/paystack", ph.SaveOrder())

			r.Post("/pay-on-delivery", oh.SaveOrder())
			r.Put("/pay-on-delivery/:id", oh.SaveOrderUpdate())
		}, "payment.",
	)
}

func orderRouter(router fiber.Router, ar *ApiRouter) {
	h := handler.NewOrderHandler(ar.db, ar.logis, ar.noti)

	router.Route(
		"/orders", func(r fiber.Router) {
			r.Get("", h.Fetch())
			r.Get("/:id", h.FetchById())

			r.Get("/:merchant/store", h.FetchAllByStore())
			r.Get("/agent/store/:agent", h.FetchAllByAgentStore())
			r.Get("/:merchant/store/:order", h.FetchByStore())

			r.Get("/:id/:userType", h.FetchByUser())

			r.Put("/update-order-details-status", h.UpdateOrderDetailStatus())
			r.Put("/update-approval-status/:order", h.UpdateOrderApprovalStatus())

			r.Post("/get-fare-estimate", h.FetchOrderFareEstimate())
			r.Post("/webhook/update-order-details-status", h.ListenTookanWebhook())
			r.Post("/test/order/creation", h.TestOrderCreation())
		}, "order.",
	)
}

func addressRouter(router fiber.Router, ar *ApiRouter) {
	h := handler.NewAddressHandler(ar.db, ar.ms)
	psh := handler.NewPickupStationHandler(ar.db)

	router.Route(
		"/address", func(r fiber.Router) {
			r.Get("/all/:userId/:userType", h.FetchAllByUser())

			r.Get("/pickup-station/all", psh.Fetch())

			r.Get("/default/:userId/:userType", h.FetchByUser())

			r.Post("/create/:userId/:userType", request.ValidateAddress(), h.Create())

			r.Put("/update/:addressId", request.ValidateAddress(), h.Update())

			r.Put("/set-default/:userId/:userType/:addressId", h.SaveDefaultAddress())

		}, "address.",
	)
}

func authRouter(router fiber.Router, ar *ApiRouter) {
	h := handler.NewAuthHandler(ar.db, ar.noti, ar.app.JWT, ar.app.Cache)
	router.Route(
		"/auth", func(r fiber.Router) {
			r.Get("/user", ar.middleware.IsAuthorized(), h.FetchAuthUser())
			r.Get("/refresh-token", ar.middleware.IsAuthorized(), h.RefreshToken())
			r.Post("/signout", ar.middleware.IsAuthorized(), h.Logout())
			r.Post("/signin", ar.middleware.IsAuthorized(), request.ValidateUser(), h.Login())
			r.Post(
				"/send-user-verification-code", ar.middleware.IsAuthorized(), request.ValidateUserName(true),
				h.SendVerificationCode(),
			)
			r.Post(
				"/send-password-reset-code", ar.middleware.IsAuthorized(), request.ValidateUserName(false),
				h.SendPasswordResetCode(),
			)
			r.Put(
				"/change-password/:user/:userType", ar.middleware.IsAuthorized(),
				request.ValidateChangePassword("update"), h.ChangePassword(),
			)
			r.Put(
				"/reset-password/:user/:userType", ar.middleware.IsAuthorized(),
				request.ValidateChangePassword("reset"), h.ResetPassword(),
			)
		},
	)
}

func adminRouter(r fiber.Router, ar *ApiRouter) {
	rph := handler.NewRoleAndPermissionHandler(ar.db)
	h := handler.NewAdminHandler(ar.db, ar.noti)
	router := r.Group("/admins") // middleware.IsAuthorized()

	// middleware.IsAuthorized()
	dashboardRouter := router.Group("/dashboard")
	dashboardRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", h.FetchCounts())
		},
	)
	confRouter := router.Group("/configurations")
	confRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", h.FetchConfigurations())
			r.Get("/:slug", h.FetchConfigurationByIdOrName())
			r.Put("/:configuration/:id", h.UpdateCurrentConfiguration())
		},
	)

	roleRouter := router.Group("/roles")
	roleRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", rph.Fetch())
			r.Get("/permissions", rph.FetchPermissions())
			r.Post("", request.ValidateRole(), rph.CreateRole())
			r.Put("/:id", request.ValidateRole(), rph.UpdateRole())
			r.Delete("/:id", rph.DeleteRole())
		},
	)

	acmRouter := router.Group("/account-manager")
	acmRouter.Route(
		"", func(r fiber.Router) {
			r.Get("/all", h.FetchAccountManagers())
			r.Get("/my-clients", h.FetchMyClients())
			r.Get("/my-client/orders", h.FetchMyClientOrders())
			r.Get("/my-client/purchase-requests", h.FetchMyClientsPurchaseRequest())
			r.Get("/my-client/products", h.FetchProducts())
			r.Post("/onboard", request.ValidateNewCustomer(), h.OnboardNewCustomer())
			r.Put("/assign", h.AssignAccountManager())
		},
	)

	orderRouter := router.Group("/orders")
	orderRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", h.FetchOrders())
			r.Delete("/:id", h.DeleteOrder())
		},
	)

	productRouter := router.Group("/products")
	productRouter.Route(
		"", func(r fiber.Router) {
			r.Get("", h.FetchAdminProducts())
			// r.Delete("/:id", h.DeleteOrder())
		},
	)

	router.Route(
		"", func(r fiber.Router) {
			r.Get("", h.Fetch())
			r.Get("/:id", h.FetchByID())
			r.Post("", request.ValidateAdmin(), h.Create())
			r.Put("/:id", request.ValidateAdmin(), h.Update())
			r.Delete("/:id", h.Delete())
		},
	)
}

func notificationRouter(router fiber.Router, ar *ApiRouter) {
	h := handler.NewNotificationHandler(ar.dbNoti)
	router.Route(
		"/notifications", func(r fiber.Router) {
			r.Get("/new/:userType/:userId", h.FetchNewNotifications())
			r.Get("/:userType/:userId", h.FetchAllByUser())
			r.Put("/:userType/:userId", h.MarkSelectedAsRead())
			r.Put("/:userType/:userId/:notificationId", h.MarkAsRead())
			r.Delete("/delete", h.RemoveSelected())
			r.Delete("/:id", h.RemoveByID())
		}, "notifications.",
	)
}
func pricemodelRouter(router fiber.Router, ar *ApiRouter) {
	h := handler.NewPriceModelHandler(ar.db)
	router.Route(
		"/price-models", func(r fiber.Router) {
			r.Get("", h.FetchAll())
			r.Get("/percentages", h.FetchAllPercentage())
			r.Get("/:model", h.Fetch())
			r.Post("", request.ValidatePriceModel(), h.Create())
			r.Put("/percentages/:category", request.ValidateCategoryPercentage(), h.CreatePercentage())
			r.Put("/:model", request.ValidatePriceModel(), h.Update())
			r.Delete("/:model", h.Remove())
			r.Delete("/percentages/:category", h.RemovePercentage())
		}, "price-models.",
	)
}
