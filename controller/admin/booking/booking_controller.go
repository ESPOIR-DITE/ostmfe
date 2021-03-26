package booking

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/booking"
	"ostmfe/io/booking_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Get("/delete/{bookId}", DeleteHandler(app))
	return r
}

func DeleteHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		bookId := chi.URLParam(r, "bookId")
		_, err := booking_io.DeleteBooking(bookId)
		if err != nil {
			fmt.Println(err, " error delete booking")
		}
		http.Redirect(w, r, "/admin_user/booking/", 301)
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}

		var unknown_error string
		var backend_error string
		var bookingList []booking.Booking
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}

		bookings, err := booking_io.ReadBookings()
		if err != nil {
			fmt.Println(err, " error reading bookings")
		} else {
			bookingList = getBookingData(bookings)
		}
		type PageData struct {
			Bookings      []booking.Booking
			SidebarData   misc.SidebarData
			Backend_error string
			Unknown_error string
		}
		data := PageData{bookingList, misc.GetSideBarData("booking", ""), backend_error, unknown_error}
		files := []string{
			app.Path + "admin/booking/booking.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
func getBookingData(booking2 []booking.Booking) []booking.Booking {
	var bookings []booking.Booking
	for _, bookingData := range booking2 {
		bookings = append(bookings, booking.Booking{bookingData.Id,
			bookingData.Name,
			bookingData.PhoneNumber,
			bookingData.Organisation,
			bookingData.Language,
			bookingData.Purpose,
			bookingData.Country,
			bookingData.Province,
			bookingData.Message,
			misc.FormatDateMonth(bookingData.Date),
			bookingData.IsFirstVisit,
			bookingData.Need})
	}
	return bookings
}
