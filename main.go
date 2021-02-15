package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	err := http.ListenAndServe(":80", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		logrus.WithField("remoteAddr", r.RemoteAddr).Info("Received Request")
	}))
	if err != nil {
		logrus.WithField("error", err.Error()).Info("Failed to listen")
	}
}