package visit

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/misc"
	"ostmfe/io/pageData_io"
)

func Home(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", homeHanler(app))
	r.Post("/book", BookHanler(app))

	return r
}

func BookHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		//var newEvent event2.Event
		//event_name := r.PostFormValue("event_name")
		//date, _ := time.Parse(misc.YYYYMMDD_FORMAT, r.PostFormValue("date"))
		//project := r.PostFormValue("project")
		//description := r.PostFormValue("description")
		//partner := r.PostFormValue("partner")
		//latlng := r.PostFormValue("latlng")
		//place := r.PostFormValue("place")

		http.Redirect(w, r, "/visit", 301)
		return
	}
}

func homeHanler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			VisitPageSection VisitPageSection
		}
		data := PageData{getPageData("Visit")}
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
				if pageSection.SectionName == "welcome" {
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
