package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type output struct {
	Method string `json:"method"`
	Host string `json:"host"`
	Path string `json:"path"`
	RemoteAddr string `json:"remote-addr"`
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Starting to listen on :80")
	err := http.ListenAndServe(":80", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		log := logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"host": r.Host,
			"path": r.URL.Path,
			"remote-address": r.RemoteAddr,
		})
		err := json.NewEncoder(rw).Encode(output{
			Method: r.Method,
			Host: r.Host,
			Path: r.URL.Path,
			RemoteAddr: r.RemoteAddr,
		})
		if err != nil {
			log.WithField("error", err.Error()).Error("Failed to encode response")
		} else {
			log.Info("Received Request")
		}
		rw.WriteHeader(http.StatusOK)
	}))
	if err != nil {
		logrus.WithField("error", err.Error()).Info("Failed to listen")
	}
}