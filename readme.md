# Room Booking Management System

This is a simple hotel room management system for a  privately owned hotel business.

## FEATURES FOR CUSTOMER
- Customers can search,book rooms present in the hotel for a duration if the room is not booked already.
- Customers receive mails after a successful reservation.
- Can cancel reservations.

## FEATURES FOR OWNER
- Backend Administration for the owner
- Can manage reserved rooms, like payment processing etc.
- Can block rooms from getting reservation for a time for maintenance etc.
- Can delete reservations if cancelled.


## TECHNOLOGIES USED
- Built in Go version 1.15
- Uses Go Templates and BootStrap5 for rendering the frontend
- Uses the [chi router](github.com/go-chi/chi)
- Uses SCS session management from alexedwards
- Uses [nosurf](github.com/justinas/nosurf) as middleware 
- Uses Soda for database migrations
- Uses postgresql for the database
And many other Golang 3rd party packages


# SNAPSHOTS
## LANDING PAGE
<img src="/static/github-ss/landing_page.png" alt="landing_page" style="height: 400px; width:800px;"/>


## CHECKING IF A PARTICULAR ROOM AVAILABLE IN A TIME PERIOD
<img src="/static/github-ss/room_availability_check.png" alt="room_availability_check" style="height: 400px; width:800px;"/>


## CHECKING FOR ALL AVAILABLE ROOMS IN A IN A TIME PERIOD
<img src="/static/github-ss/check_available_rooms.png" alt="availability_check" style="height: 400px; width:800px;"/>


## RESERVATION FORM
<img src="/static/github-ss/reservation_form.png" alt="reservation_form" style="height: 400px; width:800px;"/>


## OWNER LOGIN
<img src="/static/github-ss/admin_login.png" alt="owner_login" style="height: 400px; width:800px;"/>


## RESERVATION BACKEND ADMINISTRATION( RESERVATION CALENDAR)
<img src="/static/github-ss/reservation_backend_dashboard.png" alt="calendar" style="height: 400px; width:800px;"/>


## PROCESSING A RESERVATION
<img src="/static/github-ss/reservation_processing.png" alt="process_reservation" style="height: 400px; width:800px;"/>
