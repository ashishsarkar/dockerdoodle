package services

import (
	"log"
	"net/http"

	"github.com/gauravgahlot/watchdock/client/rpc"
	"github.com/gauravgahlot/watchdock/client/services/helpers"
	"github.com/gauravgahlot/watchdock/types"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)

// Handler represents struct with some data and request handlers for incoming HTTP requests
type Handler struct {
	hosts   *[]types.Host
	Clients *rpc.Clients
	Routes  map[string]handleFunc
}

// NewHandler initializes the request handler and returns a pointer to it
func NewHandler(hosts *[]types.Host) *Handler {
	h := Handler{hosts: hosts}
	h.initializeRoutes()
	return &h
}

func (h *Handler) initializeRoutes() {
	h.Routes = map[string]handleFunc{
		"/containers": h.dashboard,
		"/container":  h.getContainer,
	}
}

func (h *Handler) dashboard(w http.ResponseWriter, r *http.Request) {
	res, err := helpers.GetContainersCount(h.Clients.DockerService, h.hosts)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Something went wrong!"))
	}
	log.Println(*res)
	w.Write([]byte("Welcome to client"))
}

func (h *Handler) getContainer(w http.ResponseWriter, r *http.Request) {
	helpers.GetContainer(h.Clients.ContainerService)
	log.Println("in getContainer")
	w.Write([]byte("some data"))
}
