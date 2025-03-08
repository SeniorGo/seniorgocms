package api

import (
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
			<h1>Hello World!!</h1>
			<p>Página de inicio</p>
			<a href="/login">Login</a>


			<div id="version" style="position: absolute; left: 0; bottom: 0;"></div>
			<script>
				fetch('/version')
					.then(req => req.text())
					.then(version => {
						document.getElementById("version").innerText = version;
					});
			</script>
		`))
}
