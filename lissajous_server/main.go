// lissajous_server is a server that displays lissajous figures with URL parameters or default values. Exercise 1.12
package main

import (
	"go_exercises/lissajous"
	"log"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles, _ := strconv.Atoi(r.FormValue("cycle"))
		res, _ := strconv.ParseFloat(r.FormValue("res"), 32)
		size, _ := strconv.Atoi(r.FormValue("size"))
		nframe, _ := strconv.Atoi(r.FormValue("nframe"))
		delay, _ := strconv.Atoi(r.FormValue("delay"))
		lissajous.Lissajous(w, cycles, res, size, nframe, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
