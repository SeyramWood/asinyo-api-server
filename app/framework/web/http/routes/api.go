package routes

import (
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/handlers"
	middleware "github.com/SeyramWood/app/framework/web/http/middlewares"
	request "github.com/SeyramWood/app/framework/web/http/requests"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type ApiRouter struct {
	db *database.Adapter
}

func NewApiRouter(db *database.Adapter) *ApiRouter {
	return &ApiRouter{db}
}

func (h *ApiRouter) Router(app *fiber.App) {

	r := app.Group("/api", limiter.New())

	authRouter(r, h.db)

	agentRouter(r, h.db)

	retailMerchantRouter(r, h.db)

	supplierMerchantRouter(r, h.db)

	customerRouter(r, h.db)

	userRouter(r, h.db)

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

	router := r.Group("/auth/merchant/retailers")

	authRouter := r.Group("/merchant/retailers", middleware.Auth())

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateRetailMerchant(), h.Create()).Name("register")

	}, "auth.merchant.retailers.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "merchant.retailers.")

}
func supplierMerchantRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewSupplierMerchantHandler(db)

	router := r.Group("/auth/merchant/suppliers")

	authRouter := r.Group("/merchant/suppliers", middleware.Auth())

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateSupplierMerchant(), h.Create()).Name("register")

	}, "auth.merchant.suppliers.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "merchant.suppliers.")

}

func userRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewUserHandler(db)

	router := r.Group("/auth/users")

	authRouter := r.Group("/users", middleware.Auth())

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateRetailMerchant(), h.Create()).Name("register")

	}, "auth.users.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Get("/:id", h.FetchByID()).Name("fetch.id")
		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "users.")

}
func adminRouter(r fiber.Router, db *database.Adapter) {

	h := handlers.NewAdminHandler(db)

	router := r.Group("/auth/admins")

	authRouter := r.Group("/admins", middleware.Auth())

	router.Route("/", func(r fiber.Router) {

		r.Post("/sign-up", request.ValidateAdmin(), h.Create()).Name("register")

	}, "auth.admins.")

	authRouter.Route("/", func(r fiber.Router) {

		r.Get("/", h.Fetch()).Name("fetch")
		r.Get("/:id", h.FetchByID()).Name("fetch.id")
		r.Get("/", h.Fetch()).Name("fetch")
		r.Delete("/:id", h.Delete()).Name("delete")

	}, "admins.")

}

func authRouter(router fiber.Router, db *database.Adapter) {

	h := handlers.NewAuthHandler(db)

	router.Route("/auth", func(r fiber.Router) {

		r.Post("/login", h.Login()).Name("login")
		r.Get("/logout", middleware.Auth(), h.Logout()).Name("logout")

	}, "auth.")

}
