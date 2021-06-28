package page

import (
	"bufio"
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"io/ioutil"
	"net/http"
	"ostmfe/config"
	"ostmfe/controller/admin/adminHelper"
	"ostmfe/controller/misc"
	"ostmfe/domain/pageData"
	"ostmfe/io/pageData_io"
)

func PageHome(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pageHandler(app))
	r.Get("/edit/{pageId}", EditePageHandler(app))
	r.Get("/edit_section/{pageId}", EditePageHandler(app))
	r.Get("/delete/{pageId}", DeletePageHandler(app))
	r.Get("/delete_section/{pageId}", DeletePageHandler(app))
	r.Post("/create", CreatePageHandler(app))
	r.Post("/update_section", UpdatePageSectionHandler(app))
	r.Post("/create_section", CreatePageSectionHandler(app))
	r.Post("/create_banner", CreatePageBannerHandler(app))
	r.Post("/update_banner", UpdatePageBannerHandler(app))
	return r
}

func UpdatePageBannerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()
		var contentFile []byte
		file, _, err := r.FormFile("file")
		pageId := r.PostFormValue("pageId")
		bannerId := r.PostFormValue("bannerId")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should not happen>>>")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
			return
		} else {
			reader := bufio.NewReader(file)
			contentFile, _ = ioutil.ReadAll(reader)
		}

		if pageId != "" && bannerId != "" {
			bannerObject := pageData.Banner{bannerId, contentFile}
			_, err := pageData_io.UpdateBanner(bannerObject)
			//fmt.Println("banner: ", banner)
			if err != nil {
				fmt.Println(err, " error updating Banner")
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			}

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create a section")
			http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
			return
		}
		fmt.Println(" error missing field")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/page", 301)
		return

	}
}

func CreatePageBannerHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()
		var contentFile []byte
		file, _, err := r.FormFile("file")
		pageId := r.PostFormValue("pageId")
		description := r.PostFormValue("description")
		pageName := r.PostFormValue("pageName")
		if err != nil {
			fmt.Println(err, "<<<<<< error reading file>>>>This error should not happen>>>")
			if app.Session.GetString(r.Context(), "user-create-error") != "" {
				app.Session.Remove(r.Context(), "user-create-error")
			}
			app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
			http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
			return
		} else {
			reader := bufio.NewReader(file)
			contentFile, _ = ioutil.ReadAll(reader)
		}

		if description != "" && pageId != "" && pageName != "" {
			bannerObject := pageData.Banner{"", contentFile}
			banner, err := pageData_io.CreateBanner(bannerObject)
			//fmt.Println("banner: ", banner)
			if err != nil {
				fmt.Println(err, " error creating Banner")
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			}
			pageDataObject, err := pageData_io.ReadPageData(pageId)
			if err != nil {
				fmt.Println(err, " error creating Banner")
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			}
			newPageData := pageData.PageData{pageDataObject.Id, pageDataObject.PageName, banner.Id, pageDataObject.Description}
			_, err = pageData_io.UpdatePageData(newPageData)
			if err != nil {
				fmt.Println(err, " error updating PageData with BannerId")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create a section")
			http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
			return
		}
		fmt.Println(" error missing field")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/page", 301)
		return

	}
}

func UpdatePageSectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()
		pageSectionId := r.PostFormValue("pageSectionId")
		content := r.PostFormValue("myArea")
		pageId := r.PostFormValue("pageId")
		//pageName := r.PostFormValue("pageName")
		//description := r.PostFormValue("description")

		fmt.Println("pageSectionId", pageSectionId)
		if pageSectionId != "" && pageId != "" {
			pageSection, err := pageData_io.ReadPageSection(pageSectionId)
			if err != nil {
				fmt.Println(err, " error reading PageSection")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			} else {
				newPageSectionObject := pageData.PageSection{pageSectionId, pageId, pageSection.SectionId, misc.ConvertToByteArray(content)}
				_, err := pageData_io.UpdatePageSection(newPageSectionObject)
				if err != nil {
					fmt.Println(err, " error Updting PageSection")
					if app.Session.GetString(r.Context(), "user-create-error") != "" {
						app.Session.Remove(r.Context(), "user-create-error")
					}
					app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
					http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
					return
				}
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create a section")
			http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
			return
		}
		fmt.Println(" error missing field")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/page", 301)
		return

	}
}

func CreatePageSectionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()
		pageId := r.PostFormValue("pageId")
		name := r.PostFormValue("name")
		//pageName := r.PostFormValue("pageName")
		description := r.PostFormValue("description")
		content := r.PostFormValue("mytextarea")

		fmt.Println(pageId, "<<pageId ", name, " <<<name ", description, " description ", content, " <<content")
		if name != "" && description != "" && content != "" {
			sectionObject := pageData.SectionBlock{"", name, description}
			section, err := pageData_io.CreateSection(sectionObject)
			if err != nil {
				fmt.Println(err, " error creating section")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			}
			contentByteArray := []byte(content)
			pageSectionObject := pageData.PageSection{"", pageId, section.Id, contentByteArray}

			_, errx := pageData_io.CreatePageSection(pageSectionObject)
			if errx != nil {
				fmt.Println(errx, " error creating PageSection")
				fmt.Println(" deleting the section created ")
				_, err := pageData_io.DeleteSection(section.Id)
				if err != nil {
					fmt.Println(err, " error deleting section")
				}
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
				return
			}

			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully create a section")
			http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
			return

		}
		fmt.Println(" One field missing")
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/page/edit/"+pageId, 301)
		return
	}
}

func DeletePageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pageId := chi.URLParam(r, "pageId")

	}
}

func CreatePageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		r.ParseForm()
		name := r.PostFormValue("name")
		description := r.PostFormValue("description")

		if name != "" && description != "" {
			pageObject := pageData.PageData{"", name, "", description}
			page, err := pageData_io.CreatePageData(pageObject)
			if err != nil {
				fmt.Println(err, " error creating page")
				if app.Session.GetString(r.Context(), "user-create-error") != "" {
					app.Session.Remove(r.Context(), "user-create-error")
				}
				app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
				http.Redirect(w, r, "/admin_user/page", 301)
				return
			}
			if app.Session.GetString(r.Context(), "creation-successful") != "" {
				app.Session.Remove(r.Context(), "creation-successful")
			}
			fmt.Println(" successfully updated")
			app.Session.Put(r.Context(), "creation-successful", "You have successfully updating: "+page.PageName+"  project. ")
			http.Redirect(w, r, "/admin_user/page", 301)
			return

		}
		if app.Session.GetString(r.Context(), "user-create-error") != "" {
			app.Session.Remove(r.Context(), "user-create-error")
		}
		app.Session.Put(r.Context(), "user-create-error", "An error has occurred, Please try again late")
		http.Redirect(w, r, "/admin_user/page", 301)
		return
	}
}

func EditePageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
		pageId := chi.URLParam(r, "pageId")
		pageDataObject, err := pageData_io.ReadPageData(pageId)
		if err != nil {
			app.InfoLog.Println(err, " error reading PageData")
		}
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
		type PageData struct {
			Backend_error      string
			Unknown_error      string
			SidebarData        misc.SidebarData
			PageAggregatedData PageAggregatedData
			PageBanner         pageData.BannerImageHelper
			AdminName          string
			AdminImage         string
		}
		data := PageData{backend_error, unknown_error,
			misc.GetSideBarData("page", pageId),
			GetPageAggregatedData(pageId), misc.GetPageBannerData(pageDataObject),
			adminName, adminImage}
		files := []string{
			app.Path + "admin/page/edit_page.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
			app.Path + "base_templates/footer.html",
			app.Path + "admin/template/cards.html",
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

func pageHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !adminHelper.CheckAdminInSession(app, r) {
			http.Redirect(w, r, "/administration/", 301)
			return
		}
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
		pages, err := pageData_io.ReadPageDatas()
		if err != nil {
			fmt.Println(err, " error reading pages")
		}
		adminName, adminImage, isTrue := adminHelper.CheckAdminDataInSession(app, r)
		if !isTrue {
			fmt.Println(isTrue, "error reading adminData")
		}

		type PagePage struct {
			Backend_error string
			Unknown_error string
			Pages         []pageData.PageData
			SidebarData   misc.SidebarData
			AdminName     string
			AdminImage    string
		}

		data := PagePage{backend_error, unknown_error, pages,
			misc.GetSideBarData("page", ""), adminName, adminImage}
		files := []string{
			app.Path + "admin/page/pages.html",
			app.Path + "admin/template/navbar.html",
			app.Path + "admin/template/topbar.html",
			app.Path + "base_templates/footer.html",
			app.Path + "admin/template/cards.html",
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

type PageAggregatedData struct {
	PageDate    pageData.PageData
	PageSection []pageData.PageSectionHelper
	Section     []pageData.SectionBlock
}

//This method will return all the sections and contents of a page
func GetPageAggregatedData(pageId string) PageAggregatedData {
	var pageAggregatedData PageAggregatedData
	var sections []pageData.SectionBlock
	var pageSectionhelper []pageData.PageSectionHelper

	page, err := pageData_io.ReadPageData(pageId)
	if err != nil {
		fmt.Println(err, " error reading page")
		return pageAggregatedData
	}
	pageSections, err := pageData_io.ReadPageSectionAllOf(pageId)
	if err != nil {
		fmt.Println(err, " error reading all the page sections")
	} else {
		for _, pageSection := range pageSections {
			section, err := pageData_io.ReadSection(pageSection.SectionId)
			if err != nil {
				fmt.Println(err, " error reading section for the following sectionId: ", pageSection.SectionId)
			} else {
				pageSectionHelperObject := pageData.PageSectionHelper{pageSection.Id, pageSection.PageId, pageSection.SectionId, section.SectionName, section.Description, misc.ConvertingToString(pageSection.Content)}
				pageSectionhelper = append(pageSectionhelper, pageSectionHelperObject)
				pageSectionHelperObject = pageData.PageSectionHelper{}
				sections = append(sections, section)
				section = pageData.SectionBlock{}
			}
		}
	}
	return PageAggregatedData{page, pageSectionhelper, sections}
}
