package app

import (
	"io"
	"net/http"
)

//RenderIndex renders the index page
func (app *App) RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/index.html"}
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}

func (app *App) RenderRegisterUserPage(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/register.html"}
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}
