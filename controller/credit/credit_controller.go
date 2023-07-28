package credit

import (
	"net/http"
	"encoding/json"
	"TRANSFERSYSTEM/app/api"
	userRepository "TRANSFERSYSTEM/repository/user"
)

type Credit struct {
    Transferer string 
	Transferee string
    Amount float64 
}

type Balance struct {
    Balance float64 
}

func Transfer(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
    var creditRequest Credit
    
	err := decoder.Decode(&creditRequest)
    if err != nil {
        panic(err)
    }

	transferer := userRepository.FindById(creditRequest.Transferer)
	transferee := userRepository.FindById(creditRequest.Transferee)

	var isInValidAccounts = transferer.ID == "" || transferee.ID == ""

	if isInValidAccounts{
		api.Json(w).RespondError(api.ErrorRes{Error: "User not found", StatusCode: http.StatusNotFound})
		return
	}

	if transferer.Balance < creditRequest.Amount{
		api.Json(w).RespondError(api.ErrorRes{Error: "Credit is not suffiecient", StatusCode: http.StatusNotFound})
		return
	}
	
	var newTransfererBalance = transferer.Balance - creditRequest.Amount	
	transferer = userRepository.UpdateBalance(creditRequest.Transferer, newTransfererBalance)
	
	if transferer.ID == ""{
		api.Json(w).RespondError(api.ErrorRes{Error: "User not found", StatusCode: http.StatusNotFound})
		return
	}
	
	var newTransfereeBalance = transferee.Balance + creditRequest.Amount
	transferee = userRepository.UpdateBalance(creditRequest.Transferee, newTransfereeBalance)
	
	if transferee.ID == ""{
		api.Json(w).RespondError(api.ErrorRes{Error: "User not found", StatusCode: http.StatusNotFound})
		return
	}

	api.Json(w).Respond(api.DataRes{
		Data:          transferer,
		StatusCode:    http.StatusOK,
		StatusMessage: "Credit deducted successfully",
	})
}
