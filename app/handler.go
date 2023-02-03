package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("tz")
	locations := strings.Split(query, ",")
	w.Header().Set("Content-Type", "application/json")
	times := map[string]string{}
	if len(locations) >= 1 {
		for _, location := range locations {
			zone, err := time.LoadLocation(location)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "Invalid timezone")
				return
			}
			times[location] = time.Now().In(zone).String()
		}
		json.NewEncoder(w).Encode(times)
	} else {
		now := time.Now().UTC().String()
		times["current_time"] = now
		json.NewEncoder(w).Encode(times)
	}

}
