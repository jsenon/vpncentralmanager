package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

type wellknownResponse struct {
	Servicename        string `json:"Servicename"`
	Servicedescription string `json:"Servicedescription"`
	Version            string `json:"Version"`
	Versionfull        string `json:"Versionfull"`
	Revision           string `json:"Revision"`
	Branch             string `json:"Branch"`
	Builddate          string `json:"Builddate"`
	Swaggerdocurl      string `json:"Swaggerdocurl"`
	Healthzurl         string `json:"Healthzurl"`
	Metricurl          string `json:"Metricurl"`
	Endpoints          string `json:"Endpoints"`
}

// WellKnownFingerHandler will provide the information about the service.
func WellKnownFingerHandler(w http.ResponseWriter, _ *http.Request) {
	item := wellknownResponse{
		Servicename:        "vpncentralmanager",
		Servicedescription: "VPN Central Manager Application",
		Version:            "0.1",
		Versionfull:        "v.0.1",
		Revision:           "",
		Branch:             "",
		Builddate:          "",
		Swaggerdocurl:      "",
		Healthzurl:         "/healthz",
		Metricurl:          "",
		Endpoints:          ""}
	data, err := json.Marshal(item)
	if err != nil {
		log.Fatalf("Error in marshall: %v", err)
	}
	writeJSONResponse(w, http.StatusOK, data)
}

// Health will provide the information about state of the service.
func Health(w http.ResponseWriter, _ *http.Request) {
	data, err := json.Marshal(healthCheckResponse{Status: "UP"})
	if err != nil {
		log.Fatalf("Error in marshall: %v", err)
	}
	fmt.Println("Debug Marshall health", data)

	writeJSONResponse(w, http.StatusOK, data)
}

// writeJsonResponse will convert response to json
func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		log.Fatalf("Error in writeJSON rest server: %v", err)
	}
}
