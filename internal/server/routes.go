package server

import (
	"encoding/json"
	"log"
	"fmt"
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
	fs := http.FileServer(http.Dir("cmd/web/static")) 
    r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", s.indexHandler)
	r.Post("/create", s.createTodoHandler)
    r.Get("/todo/{id}", s.getTodoHandler)
    r.Put("/todo/{id}", s.markDoneHandler)
    r.Delete("/todo/{id}", s.deleteTodoHandler)
	
	return r
}

func (s *Server) sendTodos(w http.ResponseWriter) {
    todos, err := s.db.GetAllTodos()
    if err != nil {
        fmt.Println("Could not get all todos from db", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.ParseFiles("cmd/web/templates/index.html"))

    err = tmpl.ExecuteTemplate(w, "Todos", todos)
    if err != nil {
        fmt.Println("Could not execute template", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
    }
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	todoos, err := s.db.GetAllTodos()
	if err != nil {
		log.Fatalf("Could not get all todos form db. Err: %v", err)
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
	}

	tmpl := template.Must(template.ParseFiles("cmd/web/templates/index.html"))
	err = tmpl.Execute(w, todoos)
	if err != nil {
		log.Fatalf("Could not execute template. Err: %v", err)
	}
}

// Create a new todo
func (s *Server) createTodoHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm() 
	if err != nil {
        http.Error(w, "Invalid form data", http.StatusBadRequest)
        return
    }

    todoText := r.FormValue("todo")
    if todoText == "" {
        http.Error(w, "Todo text is required", http.StatusBadRequest)
        return
    }

    todo := database.Todo{
        Todo: todoText,
    }

    err = s.db.CreateTodo(todo.Todo)
    if err != nil {
        http.Error(w, "Failed to create todo", http.StatusInternalServerError)
        return
    }
	s.sendTodos(w)
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
	s.sendTodos(w)
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
	s.sendTodos(w)
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
	s.sendTodos(w)
}