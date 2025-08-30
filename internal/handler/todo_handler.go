package handler

import (
	"encoding/json"
	"net/http"

	"todo-api/internal/middleware"
	"todo-api/internal/model"
	"todo-api/internal/service"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	TodoService *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{TodoService: s}
}

func (h *TodoHandler) RegisterRoutes(r *mux.Router) {
	todoRouter := r.PathPrefix("/todos").Subrouter()
	todoRouter.Use(middleware.JWTMiddleware())

	todoRouter.HandleFunc("", h.Find).Methods("GET")
	todoRouter.HandleFunc("", h.Create).Methods("POST")
	todoRouter.HandleFunc("/{id}", h.Update).Methods("PUT")
	todoRouter.HandleFunc("/{id}", h.Delete).Methods("DELETE")
}

// Find godoc
// @Summary Todo 목록 조회
// @Description 인증된 사용자의 Todo 목록을 조회합니다.
// @Tags todos
// @Security BearerAuth
// @Produce json
// @Success 200 {array} model.Todo
// @Failure 400 {object} model.ApiResponse
// @Failure 401 {object} model.ApiResponse
// @Router /todos [get]
func (h *TodoHandler) Find(w http.ResponseWriter, r *http.Request) {
	res := model.ApiResponse{Message: "success"}
	username, _ := middleware.GetUsernameFromContext(r.Context())
	todos, err := h.TodoService.FindByUsername(username)
	if err != nil {
		res.Message = "not found"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Data = todos
	json.NewEncoder(w).Encode(todos)
}

// Create godoc
// @Summary Todo 생성
// @Description 새로운 Todo 항목을 생성합니다.
// @Tags todos
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param todo body model.Todo true "할 일 정보"
// @Success 200 {object} model.ApiResponse
// @Failure 400 {object} model.ApiResponse
// @Failure 401 {object} model.ApiResponse
// @Router /todos [post]
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := model.Todo{}
	res := model.ApiResponse{Message: "success"}
	json.NewDecoder(r.Body).Decode(&req)
	username, _ := middleware.GetUsernameFromContext(r.Context())
	req.Username = username

	if err := h.TodoService.Create(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "create failed: " + err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

// Update godoc
// @Summary Todo 업데이트
// @Description 기존 Todo 항목을 업데이트합니다.
// @Tags todos
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body model.Todo true "할 일 정보"
// @Success 200 {object} model.ApiResponse
// @Failure 400 {object} model.ApiResponse
// @Failure 401 {object} model.ApiResponse
// @Router /todos/{id} [put]
func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	req := model.Todo{}
	res := model.ApiResponse{Message: "success"}
	json.NewDecoder(r.Body).Decode(&req)
	req.Id = mux.Vars(r)["id"]
	username, _ := middleware.GetUsernameFromContext(r.Context())
	req.Username = username

	if err := h.TodoService.Update(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "update failed: " + err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// Delete godoc
// @Summary Todo 삭제
// @Description 기존 Todo 항목을 삭제합니다.
// @Tags todos
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} model.ApiResponse
// @Failure 400 {object} model.ApiResponse
// @Failure 401 {object} model.ApiResponse
// @Router /todos/{id} [delete]
func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	res := model.ApiResponse{Message: "success"}
	id := mux.Vars(r)["id"]

	if err := h.TodoService.Delete(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res.Message = "delete failed: " + err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(res)
}
