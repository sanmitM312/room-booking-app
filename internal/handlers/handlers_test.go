package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sanmitM312/room-booking-app/internal/models"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals-quarters", "GET", http.StatusOK},
	{"ms", "/majors-suite", "GET", http.StatusOK},
	{"sa", "/search-availability", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
	{"mr", "/make-reservation", "GET", http.StatusOK},

	// {"post-search-avail","/search-availability","POST",[]postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-02"},
	// },http.StatusOK},
	
	// {"make reservation post","/make-reservation","POST",[]postData{
	// 	{key: "first_name", value: "John"},
	// 	{key: "last_name", value: "Smith"},
	// 	{key: "email", value: "me@here.com"},
	// 	{key: "phone", value: "555-555-5555"},
	// },http.StatusOK},

}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	// test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func TestRepository_PostReservation(t *testing.T){
	reqBody := "start_date=2050-01-01"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=2050-01-02")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=John")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=1")

	req, _ := http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// form post header
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")

	// request recorder
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusSeeOther {
		t.Errorf("PostReservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for missing post body
	req, _ = http.NewRequest("POST", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusTemporaryRedirect{
		t.Errorf("PostReservation handler returned wrong response code for missing post body: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for invalid end date
	reqBody = "start_date=invalid"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=2050-01-02")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=John")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusTemporaryRedirect{
		t.Errorf("PostReservation handler returned wrong response code for invalid start date: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}
	// test for invalid end date
	reqBody = "start_date=2050-01-01"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=invalid")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=John")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusTemporaryRedirect{
		t.Errorf("PostReservation handler returned wrong response code for invalid end date: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for invalid room id
	reqBody = "start_date=2050-01-01"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=2050-01-02")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=John")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=invalid")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusTemporaryRedirect{
		t.Errorf("PostReservation handler returned wrong response code for invalid room id: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for invalid form data
	reqBody = "start_date=2050-01-01"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=2050-01-02")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=J")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=1")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusSeeOther{
		t.Errorf("PostReservation handler returned wrong response code for invalid form data: got %d, wanted %d", rr.Code, http.StatusSeeOther)
	}

	// test for failure of insertion into database
	reqBody = "start_date=2050-01-01"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=2050-01-02")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=John")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=2")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusTemporaryRedirect{
		t.Errorf("PostReservation handler returned wrong response code for failure to insert data : got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}

	// test for failure to insert restriction into database
	// in this case insertion of restriction for room id 1000 fails
	reqBody = "start_date=2050-01-01"
	reqBody= fmt.Sprintf("%s&%s",reqBody,"end_date=2050-01-02")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"first_name=John")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"last_name=Smith")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"email=john@smith.com")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"phone=123456789")
	reqBody= fmt.Sprintf("%s&%s",reqBody,"room_id=1000")

	req, _ = http.NewRequest("POST", "/make-reservation", strings.NewReader(reqBody))
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	// request recorder
	rr = httptest.NewRecorder()

	handler = http.HandlerFunc(Repo.PostReservation)

	handler.ServeHTTP(rr,req)
	if rr.Code != http.StatusTemporaryRedirect{
		t.Errorf("PostReservation handler returned wrong response code for failure to insert room restriction : got %d, wanted %d", rr.Code, http.StatusTemporaryRedirect)
	}
	
	
}
// 120
func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}
	// get a context, store our sessional variable in it and then store it in our request
	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// request recorder
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(rr, req) // fakes routes

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test case where reservation is not in session (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

	// test for non existent room (GetRoomByID in test-repo)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code: got %d, wanted %d", rr.Code, http.StatusOK)
	}

}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx // know about the header
}

func TestRepository_AvailabilityJSON(t *testing.T){
	// first case is rooms are not available
	reqBody := "start=2050-01-01"
	reqBody = fmt.Sprintf("%s&%s",reqBody,"end=2050-01-02")
	reqBody = fmt.Sprintf("%s&%s",reqBody,"room_id=1")

	// create request 
	req, _ := http.NewRequest("POST","/search-availability-json",strings.NewReader(reqBody))

	// get context with session
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	// set the request header
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")

	// make handler a handler func
	handler := http.HandlerFunc(Repo.AvailabilityJSON)

	// make request to our handler
	rr := httptest.NewRecorder()

	//make request to our handler
	handler.ServeHTTP(rr,req)

	var j jsonResponse
	err := json.Unmarshal(([]byte(rr.Body.String())),&j)
	if err != nil{
		t.Error("failed to parse json")
	}
}