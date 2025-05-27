package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"backend/ecommerce-api/db"
	"backend/ecommerce-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	json.NewDecoder(r.Body).Decode(&order)

	order.ID = primitive.NewObjectID()
	order.Status = "Pending"
	order.CreatedAt = time.Now()

	collection := db.Client.Database("ecommerce").Collection("orders")
	_, err := collection.InsertOne(context.TODO(), order)

	if err != nil {
		http.Error(w, "Order failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}
