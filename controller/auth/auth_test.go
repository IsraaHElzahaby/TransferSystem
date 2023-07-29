package auth

import (
	"TRANSFERSYSTEM/app/api"
	"testing"
	"net/http"
)

type credentials struct {
	Name    string `json:"name"`
	Password string `json:"password"`
}

func TestFailureLogin(t *testing.T){

    got := Login(credentials{"Chatterbridge", 126},)
    want := api.Json().RespondError(api.ErrorRes{Error: "Wrong credentials", StatusCode: http.StatusInternalServerError})

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestSuccessLogin(t *testing.T){

    got := Login(credentials{"Chatterbridge", 123456},)
    
	want := api.Json().Respond(api.DataRes{
		Data:          map[string]interface{}{"user": user, "token": signedToken},
		StatusCode:    http.StatusCreated,
		StatusMessage: "User created successfully",
	})

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
