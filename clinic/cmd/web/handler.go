package main

import (
	"AdvancedProgAssignment/clinic/pkg/models"
	"AdvancedProgAssignment/clinic/pkg/models/forms"
	"errors"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	id := app.session.GetInt(r, "authenticatedUserID")
	b := app.users.GetUsername(id)

	app.render(w, r, "home.page.tmpl", &templateData{
		User: b,
	})
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("name", "email", "password", "phone")
	form.MaxLength("name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	form.MaxLength("phone", 12)
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	// Try to create a new user record in the database. If the email already exists
	// add an error message to the form and re-display it.
	err = app.users.Insert(form.Get("name"), form.Get("email"), form.Get("password"), form.Get("phone"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Otherwise add a confirmation flash message to the session confirming that
	// their signup worked and asking them to log in.
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	// And redirect the user to the login page.
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "authenticatedUserID", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	app.session.Remove(r, "authenticatedUserID")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) showSPA(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "SPA.page.tmpl", nil)
}

func (app *application) showCosmetology(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "cosmetology.page.tmpl", nil)
}

func (app *application) showSelfcare(w http.ResponseWriter, r *http.Request) {

	p, err := app.products.Highthree()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "selfcareres.page.tmpl", &templateData{
		Products: p,
	})
}

func (app *application) showAboutUs(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "aboutUs.page.tmpl", nil)
}

func (app *application) showBooking(w http.ResponseWriter, r *http.Request) {
	id := app.session.GetInt(r, "authenticatedUserID")
	b, err := app.users.GetUser(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "booking.page.tmpl", &templateData{
		User: b,
	})
}

func (app *application) showNews(w http.ResponseWriter, r *http.Request) {

	n, err := app.news.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "news.page.tmpl", &templateData{
		news: n,
	})
}
