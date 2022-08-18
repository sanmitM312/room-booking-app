package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sanmitM312/room-booking-app/internal/config"
	
)

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig){
	app = a
}

func ClientError(w http.ResponseWriter, status int){
	app.InfoLog.Println("Client error with status of",status)
	http.Error(w,http.StatusText(status),status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s",err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	// for the client side
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func IsAuthenticated(r *http.Request) bool{
	exists := app.Session.Exists(r.Context(),"user_id")
	return exists
}