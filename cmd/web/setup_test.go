package main

import (
	"net/http"
	"os"
	"testing"
)
func TestMain(m *testing.M){
	// do something
	os.Exit(m.Run())
}

// object that satisfies the http handler request for the nosurf middleware
type myHandler struct{}

func(mh *myHandler)ServeHTTP (w http.ResponseWriter, r *http.Request){

}