package route

import (
	"TRANSFERSYSTEM/app/security"
	userController "TRANSFERSYSTEM/controller/user"
	creditController "TRANSFERSYSTEM/controller/credit"
	authController "TRANSFERSYSTEM/controller/auth"
	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
}

func RegisterRoutes() *chi.Mux {
	handleAPI()

	return router
}

func handleAPI() {
	// Auth
	router.Post("/api/login", authController.Login)
	// router.Post("/api/register", authController.Register)

	// Secured Routes
	router.Route("/api", func(router chi.Router) {
		router.Use(security.Auth)

		// user
		router.Get("/user", userController.FindAll)
		router.Post("/credit/tranfer", creditController.Transfer)
	})
}
