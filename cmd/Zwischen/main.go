package main

import (
	"net/http"

	"github.com/HibiscusJuice/Zwischen/pkg/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
