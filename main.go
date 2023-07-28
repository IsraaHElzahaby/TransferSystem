package main
 
import (
	// "flashcards-api/app/datafixtures"
	"flashcards-api/route"
    "log"
    "net/http"
)
 
type User struct {

    Exported string
    ExportedID string `json:"id"`
    ExportedName string `json:"name"`
    ExportedBalance string `json:"balance"`
}

func main() {
	// datafixtures.LoadUsers()
    router := route.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}