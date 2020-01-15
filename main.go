package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Port string
}

func main() {
	var con = Config{
		Port: ":8888",
	}
	r, cors := con.NewRouter()

	//r.HandleFunc("/", APISet)
	r.HandleFunc("/hc", Test)

	err := RunServe(con.Port, r, cors)
	if err != nil {
		log.Println(err)
	}
}

func RunServe(port string, r *mux.Router, cors func(http.Handler) http.Handler) error {
	return http.ListenAndServe(port, cors(r))
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test OK!")
}

func (rou *Config) NewRouter() (*mux.Router, func(http.Handler) http.Handler) {
	r := mux.NewRouter()

	str := "http://localhost"
	str += rou.Port

	allowedOrigins := handlers.AllowedOrigins([]string{str})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})

	return r, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)
}
