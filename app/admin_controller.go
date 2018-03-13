package app

import (
	"io"
	"net/http"

	"github.com/sagarp-webonise/todo/pkg/framework"
)

func (app *App) adminPanel(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/admin.html"}
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}

func (app *App) logout(w *framework.Response, r *framework.Request) {
	if r.Method == "POST" {
		sessionCookie, _ := r.Cookie("user_session")
		deleted := app.Redis.DeleteKey(sessionCookie.Value)
		app.Log.Info(deleted)
		err := app.Session.DestroySession("user_session", w, r)
		if err != nil {
			app.Log.Error(err)
		}
		w.Redirect("/", r.Request)
	}
}
