package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)
	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/services/spa", dynamicMiddleware.ThenFunc(app.showSPA))
	mux.Get("/services/cosmetology", dynamicMiddleware.ThenFunc(app.showCosmetology))
	mux.Get("/products/selfcare", dynamicMiddleware.ThenFunc(app.showSelfcare))
	mux.Get("/info/aboutUs", dynamicMiddleware.ThenFunc(app.showAboutUs))
	mux.Get("/booking", dynamicMiddleware.ThenFunc(app.showBooking))
	mux.Get("/news", dynamicMiddleware.ThenFunc(app.showNews))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
