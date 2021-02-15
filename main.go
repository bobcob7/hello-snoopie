package main

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

type output struct {
	Method string `json:"method"`
	Host string `json:"host"`
	Path string `json:"path"`
	RemoteAddr string `json:"remote-addr"`
	Nonce string `json:"nonce"`
}

func getNonce(r io.Reader, size int) string {
	var output string
	remaining := size
	for {
		if len(output) >= size {
			return output
		}
		buff := make([]byte, remaining)
		_, err := r.Read(buff)
		if err != nil {
			panic(err)
		}
		for _, ch := range buff {
			if ch >= 'a' && ch <= 'z' {
				output += string(ch)
				remaining--
			}
		}
	}
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
			Nonce: getNonce(rand.Reader, 16),
		})
		if err != nil {
			log.WithField("error", err.Error()).Error("Failed to encode response")
		} else {
			log.Info("Received Request")
		}
	}))
	if err != nil {
		logrus.WithField("error", err.Error()).Info("Failed to listen")
	}
}