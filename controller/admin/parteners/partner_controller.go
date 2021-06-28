package parteners

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	partner2 "ostmfe/domain/partner"
	"ostmfe/io/partner_io"
)

func PartnerHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", PartenersHandler(app))
	r.Get("/new", NewPartenersHandler(app))
	r.Get("/delete/{partnerId}", DeletePartenersHandler(app))
	r.Post("/create", CreatePartenersHandler(app))
	r.Post("/update", UpdatePartenersHandler(app))
	return r
}

func DeletePartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partnerId := chi.URLParam(r, "partnerId")
		partner, err := partner_io.ReadPartner(partnerId)
		if err != nil {
			fmt.Println(err, " error creating a new partner")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/partner", 301)
			return
		}
		_, erra := partner_io.DeletePartner(partner.Id)
		if erra != nil {
			fmt.Println(erra, " error deleting partner")
		}
		if app.Session.GetString(r.Context(), "creation-successful") != "" {
			app.Session.Remove(r.Context(), "creation-successful")
		}
		app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new partenr : "+partner.Name)
		http.Redirect(w, r, "/admin_user/partner", 301)
		return
	}
}

func UpdatePartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		partner_name := r.PostFormValue("partner_name")
		partnerId := r.PostFormValue("partnerId")
		url := r.PostFormValue("url")
		partner, err := partner_io.ReadPartner(partnerId)
		if err != nil {
			fmt.Println(err, " error creating a new partner")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/partner", 301)
			return
		}
		if url != "" && partner_name != "" && partnerId != "" {
			partner := partner2.Partner{partnerId, partner_name, partner.Description, url}
			partnerResult, err := partner_io.UpdatePartner(partner)
			if err != nil {
				fmt.Println(err, " error creating a new partner")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/partner", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new partenr : "+partnerResult.Name)
			http.Redirect(w, r, "/admin_user/partner", 301)
			return
		}
		fmt.Println(err, " error Missing fields")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/partner", 301)
		return
	}
}
func CreatePartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		partner_name := r.PostFormValue("partner_name")
		description := r.PostFormValue("description")
		url := r.PostFormValue("url")
		if url != "" && description != "" && partner_name != "" {
			partner := partner2.Partner{"", partner_name, description, url}
			partnerResult, err := partner_io.CreatePartner(partner)
			if err != nil {
				fmt.Println(err, " error creating a new partner")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/partner", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create an new partenr : "+partnerResult.Name)
			http.Redirect(w, r, "/admin_user/partner", 301)
			return
		}
	}
}

func EditePartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unknown_error string
		var backend_error string
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}

		type PagePage struct {
			Backend_error string
			Unknown_error string
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PagePage{backend_error, unknown_error,
			misc.GetSideBarData("partner", ""),
			adminName, adminImage}

		files := []string{
			app.Path + "admin/collection/edit_partner.html",
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

func PartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unknown_error string
		var backend_error string
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}
		partners, err := partner_io.ReadPartners()
		if err != nil {
			fmt.Println(err, " error reading Partners")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			Partners      []partner2.Partner
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}
		data := PagePage{backend_error, unknown_error,
			partners, misc.GetSideBarData("partner", ""), adminName, adminImage}
		files := []string{
			app.Path + "admin/partner/partners.html",
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

func NewPartenersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var unknown_error string
		var backend_error string
		if app.Session.GetString(r.Context(), "creation-unknown-error") != "" {
			unknown_error = app.Session.GetString(r.Context(), "creation-unknown-error")
			app.Session.Remove(r.Context(), "creation-unknown-error")
		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			backend_error = app.Session.GetString(r.Context(), "user-create-error")
			app.Session.Remove(r.Context(), "user-create-error")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}
		type PagePage struct {
			Backend_error string
			Unknown_error string
			AdminName     string
			AdminImage    string
		}
		data := PagePage{backend_error, unknown_error, adminName, adminImage}
		files := []string{
			app.Path + "admin/partner/new_partner.html",
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
