package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/sagarp-webonise/todo/app/models"
	"github.com/sagarp-webonise/todo/pkg/framework"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) registerUser(w *framework.Response, r *framework.Request) {
	if r.Method == "POST" {
		passwordHash, error := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 10)
		if error != nil {
			panic(errors.New("password could not get hashed"))
		}

		var user models.User
		user.Username = r.FormValue("username")
		user.Password = string(passwordHash)
		user.FirstName = sql.NullString{String: r.FormValue("firstname"), Valid: true}
		user.LastName = sql.NullString{String: r.FormValue("lastname"), Valid: true}
		user.Email = sql.NullString{String: r.FormValue("email"), Valid: true}
		user.LoginAttempts = 0
		user.Modified = pq.NullTime{Time: time.Now().UTC(), Valid: true}
		user.Created = pq.NullTime{Time: time.Now().UTC(), Valid: true}

		err := user.Insert(app.DB)
		if err != nil {
			app.Log.Error("User cannot be saved!")
		}
		w.Redirect("/", r.Request)
	}
}

func (app *App) loginUser(w *framework.Response, r *framework.Request) {
	if r.Method == "POST" {
		userDetails, err := models.UserByUsername(app.DB, r.FormValue("username"))
		if err != nil {
			app.Log.Error(err)
		}
		if hashErr := bcrypt.CompareHashAndPassword([]byte(userDetails.Password), []byte(r.FormValue("password"))); hashErr == nil {
			//successful login, redirect to admin page
			sessionId, err := app.Session.CreateSession("user_session", w, time.Now().UTC().Add(time.Minute*30))
			if err != nil {
				app.Log.Error(err)
			}

			userSession := models.UsersSession{
				UserID:    sql.NullInt64{Int64: int64(userDetails.ID), Valid: true},
				SessionID: sessionId.String(),
				Modified:  pq.NullTime{Time: time.Now().UTC(), Valid: true},
				Created:   pq.NullTime{Time: time.Now().UTC(), Valid: true},
			}

			//save the session in db
			//userSession.Upsert(app.DB)
			redisUserObject, _ := json.Marshal(userSession)

			//save session in Redis
			app.Redis.StoreKeyValue(sessionId.String(), string(redisUserObject), 0)
			w.Redirect("/home", r.Request)

		} else {
			if userDetails.LoginAttempts == 5 {
				app.Log.Info("Login Attempts Over! Game Over!")
				w.Redirect("/", r.Request)
			} else {
				userDetails.LoginAttempts += 1
				userDetails.Save(app.DB)
				w.Redirect("/", r.Request)
			}
		}
	}
}
