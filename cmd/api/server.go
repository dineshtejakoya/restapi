package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Router")
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Router")
		// fmt.Println(r.Method)
		// if r.Method == http.MethodGet {
		// 	w.Write([]byte("Hello Get Method on Teachers Route"))
		// 	fmt.Println("Hello Get Method on Teachers Route")
		// 	return
		// }

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
		w.Write([]byte("Hello Teachers Route"))
		fmt.Println("Hello Teachers Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Router")
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Router")
		w.Write([]byte("Hello Execs Route"))
		fmt.Println("Hello Execs Route")
	})

	fmt.Println("Server is running on port:", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}

}
