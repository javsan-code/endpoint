package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"backend/ecommerce-api/db"
	"backend/ecommerce-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var item model.CartItem
	json.NewDecoder(r.Body).Decode(&item)
	item.ID = primitive.NewObjectID()

	collection := db.Client.Database("ecommerce").Collection("cart")
	_, err := collection.InsertOne(context.TODO(), item)

	if err != nil {
		http.Error(w, "Add to cart failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}
