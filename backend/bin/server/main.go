package main

import (
	"encoding/json"
	"fmt"
	"log"
	"masoodrb/packform/repositories/deliveries"
	"masoodrb/packform/utils"
	"net/http"
	"time"
)

const dateFormat = "Mon Jan 02 2006 15:04:05 GMT-0700"

func main_() {

}

func main() {

	port := utils.GoDotEnvVariable("SERVER_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "%s", "Please use frontend port to interact with the <a href='http://localhost:5173'>app</a>!")
	})

	http.HandleFunc("/order-details", func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query().Get("query")
		if query != "" {
			fmt.Println(query)
			payload := deliveries.GetDeliveryDetails(query)
			json.NewEncoder(w).Encode(payload)

		} else {
			payload := deliveries.GetAllDeliveryDetails()
			json.NewEncoder(w).Encode(payload)
		}

	})

	http.HandleFunc("/filter-by-date", func(w http.ResponseWriter, r *http.Request) {

		loc, err := time.LoadLocation("Australia/Melbourne")
		if err != nil {
			log.Fatal("unable to get time zone location")
		}

		startDateStr := r.URL.Query().Get("startDate")
		endDateStr := r.URL.Query().Get("endDate")
		searchString := r.URL.Query().Get("query")

		var startDate time.Time
		var endDate time.Time

		if startDateStr != "" {
			startDate, err = time.ParseInLocation(dateFormat, startDateStr, loc)
			if err != nil {
				log.Fatalf("Error parsing startDate")
			}
		}

		if endDateStr != "" {
			endDate, err = time.ParseInLocation(dateFormat, endDateStr, loc)
			if err != nil {
				log.Fatalf("Error parsing endDate")
			}
		}

		payload := deliveries.GetDeliveryDetailsByDate(startDate.UTC(), endDate.UTC(), searchString)
		json.NewEncoder(w).Encode(payload)
	})

	// listen to port
	http.ListenAndServe(":"+port, nil)
}
