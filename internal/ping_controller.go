package internal

import (
	"github.com/kaddiya/todo/pkg/framework"
)

//Ping will indicate the health
func (a App) ping(w *framework.Response, r *framework.Request) {
	a.Log.Info("hello from the log side")
	w.Message("pong")
}
