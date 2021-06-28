package visit

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/domain/booking"
	"ostmfe/io/booking_io"
	"ostmfe/io/pageData_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHandler(app))
	r.Post("/book", BookHandler(app))

	return r
}

func BookHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		isFirstVisit := false

		name := r.PostFormValue("name")
		phoneNumber := r.PostFormValue("phoneNumber")
		organisation := r.PostFormValue("organisation")
		language := r.PostFormValue("language")
		purpose := r.PostFormValue("purpose")
		need := r.PostFormValue("need")
		visitTime := r.PostFormValue("visitTime") // choose from first time or have visited before.
		date := r.PostFormValue("date")
		Country := r.PostFormValue("country")
		province := r.PostFormValue("province")
		message := r.PostFormValue("Message")

		if visitTime != "yes" {
			isFirstVisit = true
		}
		fmt.Println(isFirstVisit)
		bookingObject := booking.Booking{"", name, phoneNumber, organisation, language, purpose, Country, province, message, date, isFirstVisit, need}
		_, err := booking_io.CreateBooking(bookingObject)
		if err != nil {
			fmt.Println(err, " error create booking")
		}

		http.Redirect(w, r, "/visit", 301)
		return
	}
}

func homeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			VisitPageSection VisitPageSection
			VisitBanner      string
		}

		var bannerImage string
		banner, err := misc.GetBanner("Visit-page")
		if err != nil {
			fmt.Println(err, " There is an error when reading people pageBanner")
		} else {
			bannerImage = banner.Id
		}
		data := PageData{getPageData("Visit-page"), bannerImage}
		files := []string{
			app.Path + "visit/visit.html",
			app.Path + "base_templates/navigator.html",
			app.Path + "base_templates/footer.html",
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

type VisitPageSection struct {
	Welcome       string
	MuseumIntro   string
	MuseumService string
	MuseumAddress string
	BookingInfo   string
}

func getPageData(pageName string) VisitPageSection {
	var visitPage VisitPageSection
	var welcome string
	var museumIntro string
	var museumService string
	var museumAddress string
	var bookingInfo string

	page, err := pageData_io.ReadPageDataWIthName(pageName)
	if err != nil {
		fmt.Println(err, " error reading page, this may not exist")
		return visitPage
	} else {
		pageDateSectionObject, err := pageData_io.ReadPageSectionAllOf(page.Id)
		if err != nil {
			fmt.Println(err, " error reading page")
		}
		for _, pageDateSection := range pageDateSectionObject {
			pageSection, err := pageData_io.ReadSection(pageDateSection.SectionId)
			if err != nil {
				fmt.Println(err, " error reading page")
			} else {
				if pageSection.SectionName == "Welcome" {
					//fmt.Println(" Introduction",pageSection)
					welcome = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "museumIntro" {
					//fmt.Println(" StaffIntroTitle",pageSection)
					museumIntro = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "museumService" {
					//fmt.Println(" MemberIntroTitle",pageSection)
					museumService = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "museumAddress" {
					//fmt.Println(" MemberIntroTitle",pageSection)
					museumAddress = misc.ConvertingToString(pageDateSection.Content)
				}
				if pageSection.SectionName == "bookingInfo" {
					//fmt.Println(" MemberIntroTitle",pageSection)
					bookingInfo = misc.ConvertingToString(pageDateSection.Content)
				}

			}
		}
	}
	return VisitPageSection{welcome, museumIntro, museumService, museumAddress, bookingInfo}
}
