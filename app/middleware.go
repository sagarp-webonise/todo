package app

import (
	"net/http"

	"github.com/sagarp-webonise/todo/app/models"
	"github.com/sagarp-webonise/todo/pkg/framework"
)

// Handle will be serving only those requests that dont need to be authed
func (app *App) handle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		handler(&res, &req)
		res.Write()
	})
}

//RenderView renders a view
func (app *App) renderView(viewHandler func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		viewHandler(w, r)
	})
}

//Handle will be serving authed requests
func (app *App) secureHandle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, _ := r.Cookie("user_session")
		authUser, _ := models.UsersSessionBySessionID(app.DB, sessionCookie.Value)
		if authUser != nil {
			res := framework.NewResponse(w)
			req := framework.Request{Request: r}
			handler(&res, &req)
			res.Write()
		} else {
			http.Redirect(w, r, "/", 401)
		}
	})
}

//RenderView renders a view
func (app *App) secureRenderView(viewHandler func(w http.ResponseWriter, req *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie("user_session")
		if sessionCookie == nil {
			app.Log.Info(err)
			http.Redirect(w, r, "/", 401)
		} else {
			val, err := app.Redis.GetKeyValue(sessionCookie.Value)
			app.Log.Info(val)
			if err != nil {
				app.Log.Error(err)
				http.Redirect(w, r, "/", 401)
			} else {
				app.Log.Info(val)
				viewHandler(w, r)
			}
		}
	})
}
