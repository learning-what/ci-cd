package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var Props struct {
	Path struct {
		Base string `json:"base"`
	} `json:"path"`
	Serve struct {
		Port string `json:"port"`
	} `json:"serve"`
	DB struct {
		PG struct {
			Dsn string `json:"dsn"`
		} `json:"pg"`
	} `json:"db"`
}

type Form struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = map[string]string{
	"zhangsan": "123456",
	"lisi": "abcdef",
	"foo": "bar",
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}

	form := Form{}
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		log.Printf("decode request body: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if form.Username == "" || form.Password == "" {
		http.Error(w, "username or password can not be empty", http.StatusBadRequest)
		return
	}

	password, exist := users[form.Username]
	if !exist {
		http.Error(w, "user is not exist", http.StatusBadRequest)
		return
	}
	if password != form.Password {
		http.Error(w, "username or password is not correct", http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"code": 200,
		"msg":  "~",
	})
}

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Llongfile)

	// load props
	dir := os.Getenv("CONFIG_DIR")
	if dir == "" {
		log.Fatalf("get empty config directory")
	}
	f, err := os.Open(filepath.Join(dir, "ci-cd.dev.json"))
	if err != nil {
		log.Fatalf("open config file: %s", err)
	}
	err = json.NewDecoder(f).Decode(&Props)
	if err != nil {
		log.Fatalf("decode props: %s", err)
	}
	log.Println("configuration is loaded")

	// define request router
	r := mux.NewRouter()
	r.HandleFunc("/api/login", Login).Methods("POST")
	log.Printf("server is running at %s", Props.Serve.Port)
	err = http.ListenAndServe(Props.Serve.Port, r)
	if err != nil {
		log.Fatalf("server run: %s", err)
	}
}
