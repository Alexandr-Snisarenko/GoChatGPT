
package handlers

import (
    "encoding/json"
    "net/http"
    "myapp/db"
    "github.com/gorilla/mux"
)

type User struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Email *string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    database, _ := db.Connect()
    defer database.Close()

    err := database.QueryRow(
        "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
        user.Name, user.Email,
    ).Scan(&user.ID)

    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    database, _ := db.Connect()
    defer database.Close()

    rows, _ := database.Query("SELECT id, name, email FROM users")
    var users []User

    for rows.Next() {
        var u User
        rows.Scan(&u.ID, &u.Name, &u.Email)
        users = append(users, u)
    }

    json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    database, _ := db.Connect()
    defer database.Close()

    _, err := database.Exec("DELETE FROM users WHERE id=$1", id)
    if err != nil {
        http.Error(w, err.Error(), 500)
    }
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    var user User
    json.NewDecoder(r.Body).Decode(&user)

    database, _ := db.Connect()
    defer database.Close()

    _, err := database.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
    if err != nil {
        http.Error(w, err.Error(), 500)
    }
}
