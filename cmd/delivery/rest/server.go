package rest

import (
	"fmt"
	"net/http"

	"200lab/cmd/delivery/component"

	"github.com/go-chi/chi/v5"

	"200lab/cmd/delivery/rest/transportation/restaurant"
)

type Server struct {
	AppContext *component.AppContext
}

func NewServer(appCtx *component.AppContext) *Server {
	return &Server{AppContext: appCtx}
}

func (s *Server) ListenAndServe() {
	handler := s.AppContext.GetRestHandler()
	// config := s.AppContext.GetConfig()

	initRoute(handler, s.AppContext)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%s", "", "6886"),
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	server.ListenAndServe()
}

func initRoute(handler *chi.Mux, appCtx *component.AppContext) {
	handler.Route("/delivery", func(r chi.Router) {
		r.Get("/restaurant", restaurant.GetRestaurant(appCtx))
	})
}
