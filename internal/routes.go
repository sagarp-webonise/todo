package internal

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {
	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping))
	app.Router.Get("/api/roles", app.handle(app.GetAllRoles))
	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/roles", app.renderView(app.DisplayRoles))

}
