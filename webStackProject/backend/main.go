
package main

import (
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
    "github.com/gorilla/mux"
    "myapp/handlers"
)

func main() {
    godotenv.Load()

    r := mux.NewRouter()
    r.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/api/users/{id}", handlers.DeleteUser).Methods("DELETE")
    r.HandleFunc("/api/users/{id}", handlers.UpdateUser).Methods("PUT")

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
