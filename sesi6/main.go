package main

import (
	"bangun2D"
	"bangun3D"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func LuasPersegi(w http.ResponseWriter, r *http.Request) {
	sisi := 10
	luas := bangun2D.LuasPersegi(sisi)

	str, err := json.Marshal(luas)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Print(err) //
	} else {
		io.WriteString(w, string(str))
	}

	// io.WriteString(w, string(strconv.FormatFloat(luas, 'f', -1, 64)))
}

func VolumeTabung(w http.ResponseWriter, r *http.Request) {
	jari, _ := strconv.Atoi(r.FormValue("Jari"))
	t, _ := strconv.Atoi(r.FormValue("Tinggi"))

	vol := bangun3D.VolTabung(jari, t)

	str, err := json.Marshal(vol)
	if err != nil {
		fmt.Print(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(str))
	}

}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/LuasPersegi", LuasPersegi)
	mux.HandleFunc("/VolumeTabung", VolumeTabung)

	http.ListenAndServe(":5050", mux)
}
