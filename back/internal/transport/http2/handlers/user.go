package handlers

import (
	"back/internal/dto"
	"back/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(req.TgId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) CheckUserExists(w http.ResponseWriter, r *http.Request) {
	tgIdStr := r.URL.Query().Get("tgId")
	if tgIdStr == "" {
		http.Error(w, "tgId parameter is required", http.StatusBadRequest)
		return
	}

	tgId, err := strconv.ParseInt(tgIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid tgId", http.StatusBadRequest)
		return
	}

	exists, err := h.service.CheckUserExists(tgId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"exists": exists}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUserByTgID(w http.ResponseWriter, r *http.Request) {
	tgIdStr := r.URL.Query().Get("tgId")
	if tgIdStr == "" {
		http.Error(w, "tgId parameter is required", http.StatusBadRequest)
		return
	}

	tgId, err := strconv.ParseInt(tgIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid tgId", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByTgId(tgId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
