package server

import (
	"encoding/json"
	"log"
	"net/http"
	"html/template"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"todo-go/internal/database"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.indexHandler)
	r.Post("/create", s.createTodoHandler)
    r.Get("/todos/{id}", s.getTodoHandler)
    r.Put("/todos/{id}", s.markDoneHandler)
    r.Delete("/todos/{id}", s.deleteTodoHandler)
	r.Get("/health", s.healthHandler)

	return r
}


func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	todoos, err := s.db.GetAllTodos()
	if err != nil {
		log.Fatalf("Could not get all todos form db. Err: %v", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(w, todoos)
	if err != nil {
		log.Fatalf("Could not execute template. Err: %v", err)
	}
}

// Create a new todo
func (s *Server) createTodoHandler(w http.ResponseWriter, r *http.Request){
	var todo database.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.CreateTodo(todo.Todo)
	if err != nil {
		http.Error(w, "Faled to create todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Get a specific todo
func (s *Server) getTodoHandler(w http.ResponseWriter, r *http.Request){
	idStr :=chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	todo, err := s.db.GetTodo(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	jsonResp, _ := json.Marshal(todo)
	_, _ = w.Write(jsonResp)
}

// Mark a todo as done
func (s *Server) markDoneHandler(w http.ResponseWriter, r *http.Request){
	idStr :=chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = s.db.MarkDone(id)
	if err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete a todo
func (s *Server) deleteTodoHandler(w http.ResponseWriter, r *http.Request){
	idStr :=chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = s.db.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}