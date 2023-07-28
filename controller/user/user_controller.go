package user

import (
	"flashcards-api/app/api"
	userRepository "flashcards-api/repository/user"
	"net/http"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	users := userRepository.GetAll()
	api.Json(w).Respond(api.DataRes{
		Data:          users,
		StatusCode:    http.StatusOK,
		StatusMessage: "Data retrieved successfully",
	})
}

