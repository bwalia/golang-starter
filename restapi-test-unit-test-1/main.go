package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// step 1- add a array of cars, matching the structure above and populate array with some dummy rackets
// step 2- add a http server in golang, serve data in json format from the array from step 1
// step 3- end of this task test api is working manually by typing http://localhost:8080/api/v1/cars in chrome
// step 4 - write unit test in golang automatically everytime a developer changes code
type car struct {
	Model        string `json:"Model"`
	Make         string `json:"Make"`
	Year         int    `json:"Year"`
	NumberOfTyre int    `json:"NumberOfTyre"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlePost).Methods("POST")
	r.HandleFunc("/", handleGet).Methods("GET")

	a := car{Model: "carens", Make: "kia", Year: 2022, NumberOfTyre: 4}
	b := car{Model: "scorpio", Make: "mahindra", Year: 2022, NumberOfTyre: 4}
	var carData []car
	carData = append(carData, a)
	carData = append(carData, b)
	j, _ := json.Marshal(carData)
	//j, _ = json.MarshalIndent(carData, "", "  ")
	fmt.Println(string(j))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()

}

func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "post\n")
}
