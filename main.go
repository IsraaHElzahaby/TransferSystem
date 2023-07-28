package main
 
import (
	"TRANSFERSYSTEM/app/datafixtures"
	"TRANSFERSYSTEM/route"
    "log"
    "net/http"
)

func main() {
	datafixtures.LoadUsers()
    router := route.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}