package service

import (
	_ "bytes"
	"encoding/json"
	"net/http"
	"src/client"
	"src/entity"
)

func OrderService(w http.ResponseWriter, r *http.Request) {
	order := entity.Order{}
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call api to service payment to check the amount, cai nay phai de trong client
	response := client.OrderClient(order, w)

	defer response.Body.Close()
	// parse response from paymentHandler

	var paymentResponse entity.PaymentResponse
	err = json.NewDecoder(response.Body).Decode(&paymentResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// response for orderHandler
	if !paymentResponse.Status {
		http.Error(w, "Transaction failed. Invalid balance", http.StatusBadRequest)
	}
	// encode response body as JSON and send response
	w.WriteHeader(http.StatusOK)
}
