package homeHandler

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("origin"))
	w.Write([]byte("Yo Homie !!!"))
}