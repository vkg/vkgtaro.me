package main

import (
	"io/ioutil"
	"net/http"

	"github.com/vkg/vkgtaro.me/images"
)

var help = ""

func main() {
	h, err := ioutil.ReadFile("help.txt")
	if err != nil {
		panic(err)
	}
	help = string(h)

	http.HandleFunc("/", vkgtaroHandler)
	_ = http.ListenAndServe(":8080", nil)
}

func vkgtaroHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/h":
		respond(w, help)

	// normal images
	case "/xs":
		respond(w, images.AsciiVkgtaro30)
	case "/s":
		respond(w, images.AsciiVkgtaro60)
	case "/m":
		respond(w, images.AsciiVkgtaro90)
	case "/l":
		respond(w, images.AsciiVkgtaro120)
	case "/xl":
		respond(w, images.AsciiVkgtaro150)

	// inverted images
	case "/xs/i":
		respond(w, images.AsciiVkgtaro30_i)
	case "/s/i":
		respond(w, images.AsciiVkgtaro60_i)
	case "/m/i":
		respond(w, images.AsciiVkgtaro90_i)
	case "/l/i":
		respond(w, images.AsciiVkgtaro120_i)
	case "/xl/i":
		respond(w, images.AsciiVkgtaro150_i)

	// grayscale images
	case "/xs/g":
		respond(w, images.AsciiVkgtaro30_g)
	case "/s/g":
		respond(w, images.AsciiVkgtaro60_g)
	case "/m/g":
		respond(w, images.AsciiVkgtaro90_g)
	case "/l/g":
		respond(w, images.AsciiVkgtaro120_g)
	case "/xl/g":
		respond(w, images.AsciiVkgtaro150_g)

	default:
		respond(w, images.AsciiVkgtaro90)
	}
}

func respond(w http.ResponseWriter, image string) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(image))
}
