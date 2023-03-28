package movie

import (
	"github.com/bnagoc/fav-movies-app/internal/handlers"
	"github.com/bnagoc/fav-movies-app/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	moviesURL = "/movies"
	movieURL  = "/movies/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(moviesURL, h.GetMovieList)
	router.POST(moviesURL, h.CreateMovie)
	router.GET(movieURL, h.GetMovieByUUID)
	router.PUT(movieURL, h.UpdateMovie)
	router.PATCH(movieURL, h.PartialUpdateMovie)
	router.DELETE(movieURL, h.DeleteMovie)
}

func (h *handler) GetMovieList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("list of movies"))
}

func (h *handler) GetMovieByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("movie"))
}

func (h *handler) CreateMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("create movie"))
}

func (h *handler) UpdateMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("fully update movie"))
}

func (h *handler) PartialUpdateMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("partially update movie"))
}

func (h *handler) DeleteMovie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("delete movie"))
}
