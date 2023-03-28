package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetMovieList)
	router.GET("/users", h.GetMovieByUUID)
	router.GET("/users", h.CreateMovie)
	router.GET("/users", h.UpdateMovie)
	router.GET("/users", h.PartialUpdateMoive)
}

func (h *handler) GetMovieList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of users"))
}
