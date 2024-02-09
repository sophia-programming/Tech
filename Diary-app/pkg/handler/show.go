package handler

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("/tmp/test.jpg")
	defer file.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeImageWithTemplate(w, "show", &img)
}

func writeImageWithTemplate(w http.ResponseWriter, tmpl string, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("Unable to encode image.")
	}
	//	w.Header().Set("Content-Type", "image/jpeg")
	//	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	//	if _, err := w.Write(buffer.Bytes()); err != nil {
	//		log.Println("unable to write image.")
	//	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	data := map[string]interface{}{"Title": tmpl, "Image": str}
	RenderTemplate(w, tmpl, data)
}
