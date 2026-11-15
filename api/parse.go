// Package handler is the Vercel serverless function behind the web demo: GET /api/parse?numbers=a%0Ab
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nadeem0305/ghphone"
)

type result struct {
	Input   string `json:"input"`
	Valid   bool   `json:"valid"`
	E164    string `json:"e164,omitempty"`
	Local   string `json:"local,omitempty"`
	Network string `json:"network,omitempty"`
}

const maxNumbers = 200

func Handler(w http.ResponseWriter, r *http.Request) {
	var out []result
	for _, in := range strings.Split(r.URL.Query().Get("numbers"), "\n") {
		if in = strings.TrimSpace(in); in == "" {
			continue
		}
		if len(out) == maxNumbers {
			http.Error(w, "too many numbers (max 200)", http.StatusBadRequest)
			return
		}
		n, err := ghphone.Parse(in)
		out = append(out, result{Input: in, Valid: err == nil, E164: n.E164, Local: n.Local, Network: n.Network})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}
