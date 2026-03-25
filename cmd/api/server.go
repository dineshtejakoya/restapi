package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

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
			//Parse form data (necessary for x-www-form-urlencoded)
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			fmt.Println("Form: ", r.Form)

			response := make(map[string]interface{})
			for k, v := range r.Form {
				response[k] = v[0]
			}

			fmt.Println("Processed Response Map:", response)

			// RAW Body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				return
			}
			defer r.Body.Close()

			fmt.Println("Raw Body:", body)
			fmt.Println("RAW Body:", string(body))

			//If you expect json data then unmarshall it
			var userInstance User
			err = json.Unmarshal(body, &userInstance)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled JSON into an instance of user struct", userInstance)
			fmt.Println("Received Username is :", userInstance.Name)

			// Prepare response data
			response1 := make(map[string]interface{})
			for k, v := range r.Form {
				response1[k] = v[0]
			}

			err = json.Unmarshal(body, &response1)
			if err != nil {
				return
			}
			fmt.Println("Unmarshalled JSON into a map ", response1)

			//Access the request details
			fmt.Println("Body:", r.Body)
			fmt.Println("Form:", r.Form)
			fmt.Println("Header:", r.Header)
			fmt.Println("Context:", r.Context())
			fmt.Println("ContentLength:", r.ContentLength)
			fmt.Println("Host:", r.Host)
			fmt.Println("Method:", r.Method)
			fmt.Println("Protocol:", r.Proto)
			fmt.Println("Remote Addr:", r.RemoteAddr)
			fmt.Println("Request URI:", r.RequestURI)
			fmt.Println("TLS:", r.TLS)
			fmt.Println("URL:", r.URL)
			fmt.Println("User Agent:", r.UserAgent())
			fmt.Println("Port:", r.URL.Port())
			fmt.Println("Scheme:", r.URL.Scheme)

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
		// fmt.Fprintf(w, "Hello Root Router")
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
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
