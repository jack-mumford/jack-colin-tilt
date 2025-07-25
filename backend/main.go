package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

// Todo represents a single todo item
type Todo struct {
    ID   int    `json:"id"`
    Text string `json:"text"`
    Done bool   `json:"done"`
}

var (
    todos  = make(map[int]Todo)
    nextID = 1
    mu     sync.Mutex
)

// withCORS wraps handlers to add permissive CORS headers for local dev.
func withCORS(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        h(w, r)
    }
}

func main() {
    http.HandleFunc("/todos", withCORS(handleTodos))
    http.HandleFunc("/todos/", withCORS(handleTodoByID)) // matches /todos/{id}

    log.Println("Starting backend on :8080 …")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("server error: %v", err)
    }
}

// handleTodos handles collection endpoints: GET list, POST create
func handleTodos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case http.MethodGet:
        list := make([]Todo, 0, len(todos))
        mu.Lock()
        for _, t := range todos {
            list = append(list, t)
        }
        mu.Unlock()
        if err := json.NewEncoder(w).Encode(list); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

    case http.MethodPost:
        var incoming struct {
            Text string `json:"text"`
        }
        if err := json.NewDecoder(r.Body).Decode(&incoming); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }
        if strings.TrimSpace(incoming.Text) == "" {
            http.Error(w, "text required", http.StatusBadRequest)
            return
        }
        mu.Lock()
        id := nextID
        nextID++
        t := Todo{ID: id, Text: incoming.Text, Done: false}
        todos[id] = t
        mu.Unlock()

        w.WriteHeader(http.StatusCreated)
        _ = json.NewEncoder(w).Encode(t)

    default:
        w.Header().Set("Allow", "GET, POST")
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}

// handleTodoByID handles single resource endpoints: PUT update, DELETE delete
func handleTodoByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Path expected: /todos/{id}
    parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/todos/"), "/")
    if len(parts) == 0 || parts[0] == "" {
        http.Error(w, "id required", http.StatusBadRequest)
        return
    }
    id, err := strconv.Atoi(parts[0])
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    mu.Lock()
    todo, ok := todos[id]
    mu.Unlock()
    if !ok {
        http.Error(w, "not found", http.StatusNotFound)
        return
    }

    switch r.Method {
    case http.MethodPut:
        if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
            http.Error(w, "invalid JSON", http.StatusBadRequest)
            return
        }
        todo.ID = id // ensure ID remains consistent
        mu.Lock()
        todos[id] = todo
        mu.Unlock()
        _ = json.NewEncoder(w).Encode(todo)

    case http.MethodDelete:
        mu.Lock()
        delete(todos, id)
        mu.Unlock()
        w.WriteHeader(http.StatusNoContent)

    default:
        w.Header().Set("Allow", "PUT, DELETE")
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
