package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

var todos = []Todo{
	{ID: "1", Title: "Learn Go", Details: "Study the Go programming language."},
	{ID: "2", Title: "Build API", Details: "Create a RESTful API in Go."},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, todo := range todos {
		if todo.ID == params["id"] {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(len(todos) + 1)
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, todo := range todos {
		if todo.ID == params["id"] {
			todos = append(todos[:i], todos[i+1:]...)
			var updatedTodo Todo
			_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
			updatedTodo.ID = params["id"]
			todos = append(todos, updatedTodo)
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, todo := range todos {
		if todo.ID == params["id"] {
			todos = append(todos[:i], todos[i+1:]...)
			json.NewEncoder(w).Encode(todos)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos = []Todo{}
	json.NewEncoder(w).Encode(todos)
}

func searchTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query().Get("q")
	var results []Todo
	for _, todo := range todos {
		if query != "" && (contains(todo.Title, query) || contains(todo.Details, query)) {
			results = append(results, todo)
		}
	}
	json.NewEncoder(w).Encode(results)
}

func contains(source, query string) bool {
	return len(source) >= len(query) && source[:len(query)] == query
}

func calculateComplexOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	a, err1 := strconv.ParseFloat(params["a"], 64)
	b, err2 := strconv.ParseFloat(params["b"], 64)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	result := math.Pow(a, 2) + math.Pow(b, 2)
	json.NewEncoder(w).Encode(map[string]float64{"result": result})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", getTodoByID).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodoByID).Methods("DELETE")
	r.HandleFunc("/todos", deleteAllTodos).Methods("DELETE")
	r.HandleFunc("/todos/search", searchTodos).Methods("GET")
	r.HandleFunc("/calculate/{a}/{b}", calculateComplexOperation).Methods("GET")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
