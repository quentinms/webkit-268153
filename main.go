package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/playlist.m3u8", playlistHandler)
	fs := http.FileServer(http.Dir("media"))
	http.Handle("/media/", http.StripPrefix("/media/", fs))

	log.Fatal(http.ListenAndServe(":9999", nil))
}

func playlistHandler(w http.ResponseWriter, r *http.Request) {
	primaryAuto := r.URL.Query().Get("p_auto")
	primaryDefault := r.URL.Query().Get("p_default")
	primaryLanguage := r.URL.Query().Get("p_lang")

	adAuto := r.URL.Query().Get("ad_auto")
	adDefault := r.URL.Query().Get("ad_default")
	adLanguage := r.URL.Query().Get("ad_lang")

	// Load the template file
	tmpl, err := template.ParseFiles("playlist.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Define the data to be passed to the template
	data := struct {
		PrimaryAutoselect string
		PrimaryDefault    string
		PrimaryLanguage   string
		ADAutoselect      string
		ADDefault         string
		ADLanguage        string
	}{
		PrimaryAutoselect: primaryAuto,
		PrimaryDefault:    primaryDefault,
		PrimaryLanguage:   primaryLanguage,
		ADAutoselect:      adAuto,
		ADDefault:         adDefault,
		ADLanguage:        adLanguage,
	}

	w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")

	// Execute the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
