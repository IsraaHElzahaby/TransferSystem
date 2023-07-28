package datafixtures

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"encoding/json"
	"gorm.io/gorm"	
	"flashcards-api/app/database"
	"golang.org/x/crypto/bcrypt"

)
 
type User struct {
	gorm.Model
    ID string `json:"id"`
    Name string `json:"name"`
    Balance string `json:"balance"`
    Password string `json:"password"`
}


func LoadUsers() {
	var password = "123456"
	response, err := http.Get("https://gist.githubusercontent.com/paytabs-engineering/c470210ebb19511a4e744aefc871974f/raw/6296df58428c89b8f852a6a83b0a5d0ac38289b6/accounts-mock.json")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	json.Unmarshal([]byte(responseData), &users)

	for _,userData := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		
		if err != nil {
			log.Fatal(err)
		}
		
		userData.Password = string(hashedPassword)

		database.DB.Create(&userData)
	}
	fmt.Println("Users synced")
}
