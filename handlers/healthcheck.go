package handlers

import (
	"encoding/json"
	"net/http"
)

// GitSHA populated at build-time with:
// -ldflags "-X github.com/cbsinteractive/bakery/handlers.GitSHA=$(git rev-parse HEAD)"
var GitSHA string

const HealthcheckPath = "/healthcheck"

// Healthcheck is returned when querying for service health
type Healthcheck struct {
	GitSHA string `json:"git_sha"`
}

// HealthcheckHandler responds to health check requests
type HealthcheckHandler struct{}

// ServeHTTP will return a http.StatusOK code if the service is up
func (HealthcheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(Healthcheck{GitSHA: GitSHA})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
