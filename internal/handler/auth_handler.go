package handler

import (
	"encoding/json"
	"net/http"

	"todo-api/internal/model"
	"todo-api/internal/service"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", h.Register).Methods("POST")
	r.HandleFunc("/auth/login", h.Login).Methods("POST")
}

// User Register godoc
// @Summary 사용자 등록
// @Description 새로운 사용자를 등록합니다.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "등록할 유저 정보"
// @Success 200 {object} model.ApiResponse
// @Failure 400 {object} model.ApiResponse
// @Failure 500 {object} model.ApiResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := model.User{}
	res := model.ApiResponse{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = "invalid request"
		json.NewEncoder(w).Encode(res)
		return
	}

	if err := h.AuthService.Register(req.Username, req.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "registration failed: " + err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res.Message = "user registered"
	json.NewEncoder(w).Encode(res)
}

// User Login godoc
// @Summary 사용자 로그인
// @Description 사용자 인증 후 JWT 토큰을 반환합니다.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body model.User true "로그인 정보"
// @Success 200 {object} model.ApiResponse
// @Failure 400 {object} model.ApiResponse
// @Failure 401 {object} model.ApiResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := model.User{}
	res := model.ApiResponse{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res.Message = "invalid request"
		json.NewEncoder(w).Encode(res)
		return
	}

	token, err := h.AuthService.Login(req.Username, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		res.Message = "login failed"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Message = "login success"
	res.Data = map[string]string{
		"token": token,
	}
	json.NewEncoder(w).Encode(res)
}
