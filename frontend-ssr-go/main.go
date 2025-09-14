package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Priority    string    `json:"priority"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

type UpdateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    string `json:"priority"`
}

type PageData struct {
	Todos           []Todo
	Filter          string
	CompletedCount  int
	TotalCount      int
	ActiveCount     int
	Error           string
	Success         string
}

var (
	templates *template.Template
	apiURL    string
)

func init() {
	apiURL = os.Getenv("API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080"
	}
}

func main() {
	// Parse templates
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}

	// Setup routes
	r := mux.NewRouter()
	
	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	
	// Health check
	r.HandleFunc("/health", healthHandler).Methods("GET")
	
	// Todo routes
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/todos", createTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}/toggle", toggleTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}/update", updateTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id}/delete", deleteTodoHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	fmt.Printf("Go SSR Frontend server starting on port %s\n", port)
	fmt.Printf("API URL: %s\n", apiURL)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	if filter == "" {
		filter = "all"
	}

	todos, err := fetchTodos()
	if err != nil {
		renderError(w, "Failed to load todos: "+err.Error())
		return
	}

	// Filter todos
	var filteredTodos []Todo
	completedCount := 0
	for _, todo := range todos {
		if todo.Completed {
			completedCount++
		}
		
		switch filter {
		case "active":
			if !todo.Completed {
				filteredTodos = append(filteredTodos, todo)
			}
		case "completed":
			if todo.Completed {
				filteredTodos = append(filteredTodos, todo)
			}
		default:
			filteredTodos = append(filteredTodos, todo)
		}
	}

	data := PageData{
		Todos:          filteredTodos,
		Filter:         filter,
		CompletedCount: completedCount,
		TotalCount:     len(todos),
		ActiveCount:    len(todos) - completedCount,
		Success:        r.URL.Query().Get("success"),
		Error:          r.URL.Query().Get("error"),
	}

	renderTemplate(w, "index", data)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/?error=Invalid+form+data", http.StatusSeeOther)
		return
	}

	todo := CreateTodo{
		Title:       strings.TrimSpace(r.FormValue("title")),
		Description: strings.TrimSpace(r.FormValue("description")),
		Priority:    r.FormValue("priority"),
	}

	if todo.Title == "" {
		http.Redirect(w, r, "/?error=Title+is+required", http.StatusSeeOther)
		return
	}

	if todo.Priority == "" {
		todo.Priority = "medium"
	}

	err := createTodo(todo)
	if err != nil {
		http.Redirect(w, r, "/?error=Failed+to+create+todo", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/?success=Todo+created+successfully", http.StatusSeeOther)
}

func toggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Redirect(w, r, "/?error=Invalid+todo+ID", http.StatusSeeOther)
		return
	}

	err = toggleTodo(id)
	if err != nil {
		http.Redirect(w, r, "/?error=Failed+to+toggle+todo", http.StatusSeeOther)
		return
	}

	filter := r.URL.Query().Get("filter")
	redirectURL := "/"
	if filter != "" {
		redirectURL += "?filter=" + filter
	}
	
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Redirect(w, r, "/?error=Invalid+todo+ID", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/?error=Invalid+form+data", http.StatusSeeOther)
		return
	}

	todo := UpdateTodo{
		Title:       strings.TrimSpace(r.FormValue("title")),
		Description: strings.TrimSpace(r.FormValue("description")),
		Priority:    r.FormValue("priority"),
	}

	if todo.Title == "" {
		http.Redirect(w, r, "/?error=Title+is+required", http.StatusSeeOther)
		return
	}

	err = updateTodo(id, todo)
	if err != nil {
		http.Redirect(w, r, "/?error=Failed+to+update+todo", http.StatusSeeOther)
		return
	}

	filter := r.URL.Query().Get("filter")
	redirectURL := "/?success=Todo+updated+successfully"
	if filter != "" {
		redirectURL += "&filter=" + filter
	}
	
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Redirect(w, r, "/?error=Invalid+todo+ID", http.StatusSeeOther)
		return
	}

	err = deleteTodo(id)
	if err != nil {
		http.Redirect(w, r, "/?error=Failed+to+delete+todo", http.StatusSeeOther)
		return
	}

	filter := r.URL.Query().Get("filter")
	redirectURL := "/?success=Todo+deleted+successfully"
	if filter != "" {
		redirectURL += "&filter=" + filter
	}
	
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

// API client functions
func fetchTodos() ([]Todo, error) {
	resp, err := http.Get(apiURL + "/api/v1/todos")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var todos []Todo
	err = json.NewDecoder(resp.Body).Decode(&todos)
	return todos, err
}

func createTodo(todo CreateTodo) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiURL+"/api/v1/todos", "application/json", strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return nil
}

func updateTodo(id int, todo UpdateTodo) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/todos/%d", apiURL, id), strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return nil
}

func toggleTodo(id int) error {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/todos/%d/toggle", apiURL, id), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return nil
}

func deleteTodo(id int) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/todos/%d", apiURL, id), nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return nil
}

// Template helpers
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func renderError(w http.ResponseWriter, message string) {
	data := PageData{
		Error: message,
	}
	renderTemplate(w, "index", data)
}