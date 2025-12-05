package http_utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetParam(r *http.Request, name string) string {
	return chi.URLParam(r, name)
}

func GetParamUUID(r *http.Request, name string) (uuid.UUID, error) {
	return uuid.Parse(GetParam(r, name))
}

type StartGenericHTTPServerContext struct {
	Router *chi.Mux
}

func StartGenericHTTPServer(ctx StartGenericHTTPServerContext) {
	r := chi.NewRouter()
	r.Use(LoggingMiddleware)

	r.Mount("/", ctx.Router)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", r))
}
