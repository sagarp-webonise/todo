package session

import (
	"net/http"
	"time"

	"github.com/kaddiya/todo/internal/models"
	"github.com/kaddiya/todo/pkg/framework"
	uuid "github.com/satori/go.uuid"
)

type AppSession struct {
}

type Session interface {
	CreateSession(cookieName string, w *framework.Response, expires time.Time) (uuid.UUID, error)
	DestroySession(cookieName string, w *framework.Response, r *framework.Request) error
	CheckSession(cookieName string, r *http.Request, db models.XODB) (bool, error)
}

func (se *AppSession) CreateSession(cookieName string, w *framework.Response, expires time.Time) (uuid.UUID, error) {
	sessionId, err := uuid.NewV4()
	if err != nil {
		return sessionId, err
	}
	cookie := http.Cookie{
		Name:    cookieName,
		Value:   sessionId.String(),
		Expires: expires,
	}
	http.SetCookie(w.ResponseWriter, &cookie)
	return sessionId, nil
}

func (se *AppSession) DestroySession(cookieName string, w *framework.Response, r *framework.Request) error {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return err
	}
	cookie.MaxAge = -1
	http.SetCookie(w.ResponseWriter, cookie)
	return nil
}

func (se *AppSession) CheckSession(cookieName string, r *http.Request, db models.XODB) (bool, error) {
	sessionCookie, err := r.Cookie(cookieName)
	if err != nil {
		return false, err
	} else {
		authUser, err := models.UsersSessionBySessionID(db, sessionCookie.Value)
		if authUser != nil {
			return true, nil
		} else {
			return false, err
		}
	}
}
