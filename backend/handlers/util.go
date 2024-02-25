package handlers

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/url-shortener/entities"
)

const shortenLength = 6

func CreateErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(
		entities.Message{
			Message: "Not Found",
		},
	)
}

func CreateJsonResponse(w http.ResponseWriter, jsonObj any) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jsonObj)
}

func Shorten(url string) (short string, err error) {
	hashed := sha1.New()
	_, err = io.Copy(
		hashed,
		strings.NewReader(url),
	)
	if err != nil {
		return "", err
	}

	encoded := bytes.NewBuffer([]byte{})
	_, err = io.Copy(
		base64.NewEncoder(base64.URLEncoding, encoded),
		bytes.NewReader(hashed.Sum(nil)),
	)
	if err != nil {
		return "", err
	}

	encoded.Truncate(shortenLength)

	return encoded.String(), nil
}

func RealHost(r *http.Request) string {
	forwardedFor := r.Header.Get("X-Forwarded-Host")

	if !IsBlank(forwardedFor) {
		return forwardedFor
	}

	return r.Host
}

func GetScheme(r *http.Request) string {
	scheme := r.Header.Get("X-Forwarded-Proto")

	if IsBlank(scheme) {
		scheme = r.URL.Scheme
	}

	return scheme
}

func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
