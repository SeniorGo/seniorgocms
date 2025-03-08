package api

import (
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
			<h1>Entra al CMS!!</h1>
			<p>Página de login</p>
			Usuario: <input type="text" name="username" /><br>
			Contraseña: <input type="text" name="password" /><br>
			<button>Entrar</button>
		`))
}
