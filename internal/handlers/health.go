package handlers

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //1
	// w.Write([]byte("all ok"))
	w.WriteHeader(http.StatusOK)       //2
	w.Write([]byte(`{"status":"ok"}`)) //3
}
