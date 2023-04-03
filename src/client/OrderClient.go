package client

import (
	"net/http"
	"src/entity"
	"strconv"
)

func OrderClient(order entity.Order, w http.ResponseWriter) *http.Response {
	response, err := http.Get("http://localhost:1522/payment?amount=" + strconv.FormatFloat(order.Amount, 'f', 2, 64))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	return response
}
