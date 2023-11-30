package webmanage

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Kwynto/gosession"

	"github.com/Kwynto/GracefulDB/internal/base/basicsystem/gauth"
)

// Handler after authorization
func homeDefault(w http.ResponseWriter, r *http.Request, login string) {
	// This function is complete
	err := templatesMap[HOME_TEMP_NAME].Execute(w, nil)
	if err != nil {
		slog.Debug("Internal Server Error", slog.String("err", err.Error()))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Authorization Handler
func homeAuth(w http.ResponseWriter, r *http.Request) {
	// This function is complete
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			slog.Debug("Bad request", slog.String("err", err.Error()))
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		isAuth := gauth.CheckUser(username, password)
		if isAuth {
			sesID := gosession.Start(&w, r)
			sesID.Set("auth", username)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		err := templatesMap[AUTH_TEMP_NAME].Execute(w, nil)
		if err != nil {
			slog.Debug("Internal Server Error", slog.String("err", err.Error()))
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// Handler for the main route
func home(w http.ResponseWriter, r *http.Request) {
	// This function is complete
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	sesID := gosession.Start(&w, r)
	auth := sesID.Get("auth")
	if auth == nil {
		homeAuth(w, r)
	} else {
		login := fmt.Sprint(auth)
		homeDefault(w, r, login)
	}
}

// Exit handler
func logout(w http.ResponseWriter, r *http.Request) {
	// This function is complete
	sesID := gosession.Start(&w, r)
	sesID.Remove("auth")
	http.Redirect(w, r, "/", http.StatusFound)
}

func firstmsg(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is FirstMsg!"))
	if err != nil {
		slog.Debug("The response has not been sent", slog.String("err", err.Error()))
	}
}

func mainunit(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit. This is main unit."))
	if err != nil {
		slog.Debug("The response has not been sent", slog.String("err", err.Error()))
	}
}