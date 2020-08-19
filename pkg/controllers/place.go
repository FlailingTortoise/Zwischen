package controllers

import "net/http"

type placeController struct{}

func (pc placeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response from Place Controller"))
}

func newPlaceController() *placeController {
	return &placeController{}
}
