package main

import (
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var creds model.Credentials
    
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
        return
    }
    
    if creds.Username =="" || creds.Password =="" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
        return
    }
    
    _, ok := db.Users[creds.Username]
    if ok {
        w.WriteHeader(http.StatusConflict)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username already exist"})
        return
    }
    w.WriteHeader(http.StatusOK)
    db.Users[creds.Username] = creds.Password
    json.NewEncoder(w).Encode(model.SuccessResponse{Username: creds.Username, Message: "Register Success"})
    return
    
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds model.Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)

    expectedPassword, ok := db.Users[creds.Username]
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
        return
    } else if creds.Username =="" || creds.Password =="" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Username or Password empty"})
        return
    } else if !ok || expectedPassword != creds.Password {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Wrong User or Password!"})
        return
    } 

    sessionToken := uuid.NewString()
    expiresAt := time.Now().Add(5 * time.Hour)
    http.SetCookie(w, &http.Cookie{
        Name:    "session_token",
        Value:   sessionToken,
        Expires: expiresAt, 
    })

    db.Sessions[sessionToken] = model.Session{
        Username: creds.Username,
        Expiry:   expiresAt,
    }


    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(model.SuccessResponse{Username: creds.Username, Message: "Login Success"})
    return
    
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
            return
        }
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    session := c.Value
    id := uuid.NewString()
    todo := model.Todo{}

    error :=  json.NewDecoder(r.Body).Decode(&todo)
    if error != nil {
        panic(error)
    }

    todo.Id = id 

    username := db.Sessions[session].Username
    db.Task[username] = append(db.Task[username], todo)

    
    w.WriteHeader(http.StatusOK)
    msg := fmt.Sprintf("Task %v added!", todo.Task)
    json.NewEncoder(w).Encode(model.SuccessResponse{Username: username, Message: msg})

}

func ListToDo(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
            return
        }
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    session := c.Value
    username := db.Sessions[session].Username
    task, ok := db.Task[username]
    
    if !ok {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Todolist not found!"})
        return
    } 

    data, err := json.Marshal(task)
    if err != nil {
        panic(err)
    }
    w.Write(data)

}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
            return
        }
        w.WriteHeader(http.StatusUnauthorized)
        return
    } 
    session := c.Value
    username := db.Sessions[session].Username
    db.Task[username] = []model.Todo{}
    json.NewEncoder(w).Encode(model.SuccessResponse{Username: username, Message: "Clear ToDo Success"})

}

func Logout(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	c, err := r.Cookie("session_token")
    if err != nil {
        if err == http.ErrNoCookie {
            json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})
            return
        }
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    session := c.Value

    delete(db.Sessions, session)

    http.SetCookie(w, &http.Cookie{
        Name:    "session_token",
        Value:   "",
        Expires: time.Now(),
    })
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(model.SuccessResponse{Username: username, Message: "Logout Success"})

}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(middleware.Auth(http.HandlerFunc(ClearToDo))))

	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
