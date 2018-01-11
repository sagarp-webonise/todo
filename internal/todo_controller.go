package internal

import (
	"errors"
	"io"
	"net/http"

	"github.com/kaddiya/todo/internal/models"
	"github.com/kaddiya/todo/pkg/framework"
)

func (a App) getAllTodos(w *framework.Response, r *framework.Request) {

	role, err := models.GetAllTodos(a.DB)
	if err != nil {
		a.Log.Info(err.Error())
		w.NotFound(errors.New("could not find the role"))
		return
	}
	if err != nil {
		panic(err)
		return
	}
	w.PutInData("roles", role)
}

//RenderIndex renders the index page
func (app *App) displayTodos(w http.ResponseWriter, r *http.Request) {
	roles, err := models.GetAllTodos(app.DB)
	tmplList := []string{"./web/views/base.html",
		"./web/views/roles/todo.html"}
	data := struct {
		Roles []*models.Todo
	}{roles}
	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}
