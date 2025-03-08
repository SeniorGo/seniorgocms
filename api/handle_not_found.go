package api

import (
	"net/http"
)

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`
			<h1>404 Not Found!</h1>
			<p>La página que buscas no la encontramos por ningún sitio :S mala suerte</p>
			<a href="/">Volver</a>
		`))
}
