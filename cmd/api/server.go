package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Method on Teachers Route"))
		fmt.Println("Hello Get Method on Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello Put Method on Teachers Route"))
		fmt.Println("Hello Put Method on Teachers Route")
	case http.MethodPost:
		w.Write([]byte("Hello Post Method on Teachers Route"))
		fmt.Println("Hello Post Method on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method on Teachers Route"))
		fmt.Println("Hello Patch Method on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method on Teachers Route"))
		fmt.Println("Hello Delete Method on Teachers Route")
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Method on Students Route"))
		fmt.Println("Hello Get Method on Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello Put Method on Students Route"))
		fmt.Println("Hello Put Method on Students Route")
	case http.MethodPost:

		w.Write([]byte("Hello Post Method on Students Route"))
		fmt.Println("Hello Post Method on Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method on Students Route"))
		fmt.Println("Hello Patch Method on Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method on Students Route"))
		fmt.Println("Hello Delete Method on Students Route")
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get Method on Execs Route"))
		fmt.Println("Hello Get Method on Execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello Put Method on Execs Route"))
		fmt.Println("Hello Put Method on Execs Route")
	case http.MethodPost:

		w.Write([]byte("Hello Post Method on Execs Route"))
		fmt.Println("Hello Post Method on Execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello Patch Method on Execs Route"))
		fmt.Println("Hello Patch Method on Execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello Delete Method on Execs Route"))
		fmt.Println("Hello Delete Method on Execs Route")
	}
}
func main() {
	port := ":3000"
	cert := "../../cert.pem"
	key := "../../key.pem"

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/teachers/", teachersHandler)
	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", execsHandler)
	// http.HandleFunc("/", rootHandler)
	// http.HandleFunc("/teachers/", teachersHandler)
	// http.HandleFunc("/students/", studentsHandler)
	// http.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//Create custom server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
		// Handler:   nil,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}

}
