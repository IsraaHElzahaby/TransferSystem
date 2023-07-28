package route

import (
	"flashcards-api/app/security"
	userController "flashcards-api/controller/user"
	creditController "flashcards-api/controller/credit"
	authController "flashcards-api/controller/auth"
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
