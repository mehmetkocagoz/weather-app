package main

import (
	"fmt"
	"math/rand"
	"mehmetkocagz/datascraper"
	"mehmetkocagz/model"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func keepAliveDatabase() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		fmt.Println("Inside keepAliveDatabase()")
		city := []string{"Istanbul", "Ankara", "Izmir", "Bursa", "Adana", "Gaziantep", "Konya", "Antalya", "Kayseri", "Mersin"}
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(10)
		datascraper.ScrapeDataFromOpenWeatherAPI(city[randomNumber])
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	weatherData := model.GetWeatherDataFromPostgres("Istanbul")
	tmpl.Execute(w, weatherData)
}

func main() {
	// keep alive database
	go keepAliveDatabase()
	// create a new router
	r := mux.NewRouter()

	// tell the router what to do with "/assests/" path
	fs := http.FileServer(http.Dir("assests"))
	r.PathPrefix("/assests").Handler(http.StripPrefix("/assests", fs))

	// create routers for pages
	homeRouter := r.Methods(http.MethodGet).Subrouter()
	homeRouter.HandleFunc("/", ServeHome)

	// create a new server
	s := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	fmt.Println("Running server on localhost:", s.Addr)
	s.ListenAndServe()
}
