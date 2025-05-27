package handlers

import (
	"backend/ecommerce-api/db"
	"backend/ecommerce-api/model"
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	product.ID = primitive.NewObjectID()
	primitive.NewObjectID()

	collection := db.Client.Database("ecommerce").Collection("products")
	_, err := collection.InsertOne(context.TODO(), product)

	if err != nil {
		http.Error(w, "Insert failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	collection := db.Client.Database("ecommerce").Collection("products")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Fetch failed", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var products []model.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		http.Error(w, "Parse failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}
