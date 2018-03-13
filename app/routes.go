package app

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {
	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping))
	app.Router.Get("/api/todo/", app.handle(app.GetAllTodos))
	app.Router.Post("/registerUser", app.handle(app.registerUser))
	app.Router.Post("/login", app.handle(app.loginUser))
	app.Router.Post("/logout", app.handle(app.logout))

	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/todo/", app.renderView(app.DisplayTodos))
	app.Router.Get("/register", app.renderView(app.RenderRegisterUserPage))
	app.Router.Get("/home", app.secureRenderView(app.adminPanel))

}
