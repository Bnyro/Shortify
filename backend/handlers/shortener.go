package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/url-shortener/database"
	"github.com/url-shortener/entities"
)

func CreateShortcut(w http.ResponseWriter, r *http.Request) {
	shortcut := new(entities.Shortcut)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		CreateErrorResponse(w, err)
		return
	}

	err = json.Unmarshal(reqBody, shortcut)
	if err != nil || shortcut.Full == "" {
		CreateErrorResponse(w, errors.New("Bad Request"))
		return
	}

	shortcut.Short, err = Shorten(shortcut.Short)
	if err != nil {
		CreateErrorResponse(w, err)
		return
	}

	database.Db.Create(shortcut)

	shortcut.Short = "http://" + r.Host + "/" + shortcut.Short
	CreateJsonResponse(w, shortcut)
}

func ReadShortcut(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short")

	shortcut := entities.Shortcut{
		Short: shortUrl,
	}

	database.Db.Find(&shortcut, shortcut)

	if shortcut.Full == "" {
		CreateErrorResponse(w, errors.New("Not Found"))
		return
	}

	http.Redirect(w, r, shortcut.Full, http.StatusMovedPermanently)
}
