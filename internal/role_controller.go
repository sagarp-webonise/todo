package internal

import (
	"errors"
	"io"
	"net/http"

	"github.com/kaddiya/todo/internal/models"
	"github.com/kaddiya/todo/pkg/framework"
)

func (a App) GetAllRoles(w *framework.Response, r *framework.Request) {

	role, err := models.GetAllPortalRoles(a.DB)
	if err != nil {
		a.Log.Info(err.Error())
		w.NotFound(errors.New("could not find the role"))
		return
	}
	a.Log.Info(role)
	if err != nil {
		panic(err)
		return
	}
	w.PutInData("roles", role)
}

//RenderIndex renders the index page
func (app *App) DisplayRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := models.GetAllPortalRoles(app.DB)
	tmplList := []string{"./web/views/base.html",
		"./web/views/roles/roles.html"}
	data := struct {
		Roles []*models.PortalRole
	}{roles}
	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}
