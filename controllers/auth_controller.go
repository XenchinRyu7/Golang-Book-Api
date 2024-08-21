package controllers

import (
	"encoding/json"
	"golang-book-api/helpers"
	"golang-book-api/services"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := c.authService.Register(data.Username, data.Password, data.Email)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Register Gagal")
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, user)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := c.authService.Login(data.Email, data.Password)
	if err != nil {
		if err.Error() == "email tidak ditemukan" {
			helpers.RespondError(w, http.StatusNotFound, err.Error())
			return
		}
		if err.Error() == "password salah" {
			helpers.RespondError(w, http.StatusUnauthorized, err.Error())
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "Login Gagal")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, user)
}
