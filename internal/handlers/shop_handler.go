package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/horlabyc/grocery-app/internal/domain/models"
	"github.com/horlabyc/grocery-app/internal/services"
)

type ShopHandler struct {
	service *services.ShopService
}

func NewShopHandler(service *services.ShopService) *ShopHandler {
	return &ShopHandler{
		service: service,
	}
}

func (h *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	var shop models.Shop
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateShop(r.Context(), &shop); err != nil {
		http.Error(w, "Failed to create shop", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shop)
}

func (h *ShopHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	shops, err := h.service.GetAllShops(r.Context())
	if err != nil {
		fmt.Println("Error getting all shops: ", err)
		http.Error(w, "Failed to get shops", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shops)
}

func (h *ShopHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid shop ID", http.StatusBadRequest)
		return
	}
	shop, err := h.service.GetShopByID(r.Context(), id)
	if err != nil {
		fmt.Println("Error getting all shop: ", err)
		http.Error(w, "Shop not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shop)
}

func (h *ShopHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var shop models.Shop
	if err := json.NewDecoder(r.Body).Decode(&shop); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shop.ID = id
	if err := h.service.UpdateShop(r.Context(), &shop); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shop)
}
