package routes

import (
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/handlers"
	middleware "github.com/SeyramWood/app/framework/web/http/middlewares"
	request "github.com/SeyramWood/app/framework/web/http/requests"
	"github.com/gofiber/fiber/v2"
)

type ApiRouter struct {
	db *database.Adapter
}

func NewApiRouter(db *database.Adapter) *ApiRouter {
	return &ApiRouter{db}
}

func (h *ApiRouter) Router(app *fiber.App) {

	r := app.Group("/api") //limiter.New()

	authRouter(r, h.db)

	agentRouter(r, h.db)

	merchantRouter(r, h.db)

	retailMerchantRouter(r, h.db)

	supplierMerchantRouter(r, h.db)

	customerRouter(r, h.db)

	productRouter(r, h.db)

	adminRouter(r, h.db)

}

func customerRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewCustomerHandler(db)

	router := r.Group("/auth/customers")

	authRouter := r.Group("/customers", middleware.Auth())

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateCustomer(), h.Create()).Name("register")

	}, "auth.customers.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "customers.")

}
func agentRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewAgentHandler(db)

	router := r.Group("/auth/agents")

	authRouter := r.Group("/agents", middleware.Auth())

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateAgent(), h.Create()).Name("register")

	}, "auth.agents.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "merchant.retailers.")

}

func retailMerchantRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewRetailMerchantHandler(db)
	m := handlers.NewMerchantHandler(db)

	router := r.Group("/auth/merchant/retailers")

	authRouter := r.Group("/merchant/retailers") //middleware.Auth()

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateMerchant(), m.Create()).Name("register")

	}, "auth.merchant.retailers.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "merchant.retailers.")

}
func supplierMerchantRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewSupplierMerchantHandler(db)
	m := handlers.NewMerchantHandler(db)

	router := r.Group("/auth/merchant/suppliers")

	authRouter := r.Group("/merchant/suppliers") //middleware.Auth()

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateMerchant(), m.Create()).Name("register")

	}, "auth.merchant.suppliers.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "merchant.suppliers.")

}

func merchantRouter(r fiber.Router, db *database.Adapter) {

	//m := handlers.NewMerchantHandler(db)
	msHandler := handlers.NewMerchantStoreHandler(db)

	mRouter := r.Group("/merchants")
	msRouter := mRouter.Group("/storefront")

	mRouter.Route("/", func(r fiber.Router) {

		//r.Post("/store", request.ValidateMerchant(), m.Create()).Name("register")

	}, "merchants.")

	msRouter.Route("/", func(r fiber.Router) {

		r.Get("/:merchantId", msHandler.FetchByID())

		r.Post("/:merchantId/profile", request.ValidateMerchantStore(), msHandler.Create())

		r.Put("/:storeId/account/momo", request.ValidateMerchantMomoAccount(), msHandler.SaveMomoAccount())

		r.Put("/:storeId/account/bank", request.ValidateMerchantBankAccount(), msHandler.SaveBankAccount())

	}, "merchants.")

}

func adminRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewAdminHandler(db)

	router := r.Group("/auth/admins")

	authRouter := r.Group("/admins") //middleware.Auth()

	router.Route("/", func(r fiber.Router) {

		r.Post("/create", request.ValidateAdmin(), h.Create()).Name("register")

	}, "auth.admins.")

	authRouter.Route("/", func(r fiber.Router) {
		r.Post("/create", request.ValidateAdmin(), h.Create()).Name("register")
		r.Get("/", h.Fetch()).Name("fetch")
		r.Get("/:id", h.FetchByID()).Name("fetch.id")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "admins.")

}

func productRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewProductHandler(db)
	majorH := handlers.NewProductCatMajorHandler(db)
	minorH := handlers.NewProductCatMinorHandler(db)

	router := r.Group("/products")

	authRouter := r.Group("/product") //middleware.Auth()

	router.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("all")

		r.Get("/category/major", majorH.Fetch()).Name("major")

		r.Get("/category/minor", minorH.Fetch()).Name("minor")

		r.Get("/best-seller/:merchant", h.FetchBestSellerProducts()).Name("best.sellers")

		r.Get("/:merchant/:id", h.FetchMerchantProducts()).Name("merchant.all")

		r.Get("/category/major/:merchant", h.FetchAllMerchantCategoryMajorProducts())

		r.Get("/category/:merchant/:cat/:slug", h.FetchBySlugMerchantCategoryProducts())

		r.Get("/:merchant/:id/:slug", h.FetchByIDMerchantProduct())

		r.Post("/create/:merchant", request.ValidateProduct(), h.Create()).Name("create")

	}, "products.")

	authRouter.Route("/", func(r fiber.Router) {
		r.Post("/", request.ValidateProduct(), h.Create()).Name("create")

		// r.Get("/", h.Fetch()).Name("fetch")
		// r.Get("/:id", h.FetchByID()).Name("fetch.id")
		// r.Delete("/:id", h.Delete()).Name("delete")

		r.Post("/majors", request.ValidateProductCatMajor(), majorH.Create()).Name("majors.create")

		r.Post("/minors", request.ValidateProductCatMinor(), minorH.Create()).Name("minors.create")

	}, "product.")

}

func authRouter(router fiber.Router, db *database.Adapter) {

	h := handlers.NewAuthHandler(db)

	router.Route("/auth", func(r fiber.Router) {

		r.Post("/signin", request.ValidateUser(), h.Login()).Name("signin")
		r.Get("/signout", middleware.Auth(), h.Logout()).Name("signout")
		r.Get("/user", middleware.Auth(), h.FetchAuthUser()).Name("user")

	}, "auth.")

}
