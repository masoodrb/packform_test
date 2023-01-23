package main

import (
	"fmt"
	"masoodrb/packform/repositories/deliveries"
	"masoodrb/packform/utils"
	"net/http"
)

func main() {
	deliveries.GetAllDeliveryDetails(1)
}

func main_() {

	port := utils.GoDotEnvVariable("SERVER_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "%s", "Please use frontend port to interact with the <a href='http://localhost:5173'>app</a>!")
	})

	http.HandleFunc("/orderDetails", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query != "" {
			fmt.Println(query)

		} else {
			fmt.Println("No query provided")
		}
	})

	// listen to port
	http.ListenAndServe(":"+port, nil)
}
