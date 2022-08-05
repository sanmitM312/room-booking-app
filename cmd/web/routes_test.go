package main

import (
	"testing"
	"fmt"
	"github.com/go-chi/chi"

	"github.com/sanmitM312/room-booking-app/internal/config"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T",v))
	}
}