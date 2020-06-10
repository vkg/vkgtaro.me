package handler

import (
	"net/http"
	"strings"

	"github.com/vkg/vkgtaro.me/images"
)

const help = `
       _         _                                  
__   _| | ____ _| |_ __ _ _ __ ___   _ __ ___   ___ 
\ \ / / |/ / _' | __/ _' | '__/ _ \ | '_ ' _ \ / _ \
 \ V /|   < (_| | || (_| | | | (_) || | | | | |  __/
  \_/ |_|\_\__, |\__\__,_|_|  \___(_)_| |_| |_|\___|
            __/ |
           |___/           github.com/vkg/vkgtaro.me

Usage:
 $ curl -L vkgtaro.me
 $ curl -L vkgtaro.me/{SIZE}       (xs, s, m, l, xl)
 $ curl -L vkgtaro.me/{SIZE}/{OPT} (i, g)
 $ curl -L vkgtaro.me/h            (help)
Examples:
 $ curl -L vkgtaro.me/m            (medium)
 $ curl -L vkgtaro.me/l/i          (large, introverted)
 $ curl -L vkgtaro.me/xl/g         (extra large, grayscale)

vkgtaro.me is inspired and highly respecting dogs.sh (github.com/fortwire/dogs.sh)
`

const helpHTML = `
<body style='background-color:black; color:white;'>
<pre>       _         _                                  
__   _| | ____ _| |_ __ _ _ __ ___   _ __ ___   ___ 
\ \ / / |/ / _' | __/ _' | '__/ _ \ | '_ ' _ \ / _ \
 \ V /|   < (_| | || (_| | | | (_) || | | | | |  __/
  \_/ |_|\_\__, |\__\__,_|_|  \___(_)_| |_| |_|\___|
            __/ |
           |___/           github.com/vkg/vkgtaro.me</pre>

<pre style='font-size: 80%; font-family:Courier New,Courier,Lucida Sans Typewriter,Lucida Typewriter,monospace;'>
<br><b>Usage:</b><br>
 $ curl -L vkgtaro.me<br>
 $ curl -L vkgtaro.me/{SIZE}       (xs, s, m, l, xl)<br>
 $ curl -L vkgtaro.me/{SIZE}/{OPT} (i, g)<br>
 $ curl -L vkgtaro.me/h            (help)<br>
<br><b>Examples:</b><br>
 $ curl -L vkgtaro.me/m            (medium)<br>
 $ curl -L vkgtaro.me/l/i          (large, introverted)<br>
 $ curl -L vkgtaro.me/xl/g         (extra large, grayscale)<br>

vkgtaro.me is inspired and highly respecting dogs.sh (github.com/fortwire/dogs.sh)<br>
</pre>
</body>
`

func Handler(w http.ResponseWriter, r *http.Request) {
	// From browser
	if !strings.Contains(r.Header.Get("User-Agent"), "curl") {
		respond(w, helpHTML)
		return
	}

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
