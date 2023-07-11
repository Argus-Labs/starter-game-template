package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

var EnvCardinalPort = os.Getenv("CARDINAL_PORT")

type CardinalHandlers []struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

func RegisterRpc(port string, handlers CardinalHandlers) {
	log.Printf("Attempting to register %d handlers", len(handlers))
	var paths []string
	for _, h := range handlers {
		http.HandleFunc("/"+h.Path, h.Handler)
		paths = append(paths, h.Path)
	}
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		if err := enc.Encode(paths); err != nil {
			WriteError(w, "cant marshal list", err)
		}
	})

	log.Printf("Cardinal running on server on port %s", port)
	http.ListenAndServe(":"+port, nil)
}

func GetPort() string {
	port := EnvCardinalPort
	if port == "" {
		log.Fatal().Msgf("Must specify a port via %s", EnvCardinalPort)
	}
	return port
}

func WriteError(w http.ResponseWriter, msg string, err error) {
	w.WriteHeader(500)
	payload := struct {
		Msg string
		Err string
	}{Msg: msg, Err: err.Error()}

	enc := json.NewEncoder(w)
	if err := enc.Encode(payload); err != nil {
		WriteError(w, "can't encode", err)
		return
	}

	log.Error().Err(err).Msg(msg)
}

func WriteResult(w http.ResponseWriter, v any) {
	if s, ok := v.(string); ok {
		v = struct{ Msg string }{Msg: s}
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(v); err != nil {
		WriteError(w, "can't encode", err)
		return
	}
}
