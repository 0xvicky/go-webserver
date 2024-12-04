package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	//define the endpoints
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := "static/index.html"

		htmlFile, err := os.Open(filepath)

		if err != nil {
			http.Error(w, "Error opening the file", http.StatusInternalServerError)
			fmt.Println("Error opening the file", err)
			return
		}
		defer htmlFile.Close()

		//set the respose header to somethign else
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath)

	})

	//run the server

	fmt.Printf("Server is going to run on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Failed to start server")

	}
}
