package handler

import (
	"l0/internal/domain/dto"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) getOrders(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		vars = mux.Vars(r)
		id   = vars["id"]
	)

	if id == "" {
		SendErr(w, http.StatusBadRequest, "id is empty")

		return
	}

	result, err := s.services.Orders.GetOrdersByID(ctx, id)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())

		return
	}

	SendOK(w, http.StatusOK, dto.OrderToRest(result))
}
