package controllers

import "net/http"

//RegisterControllers registers the controllers and routes
func RegisterControllers() {
	placeController := newPlaceController()

	http.Handle("/places", *placeController)
	http.Handle("/places/", *placeController)

}
