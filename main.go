package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/api/query/{id}", withCORS(QueryHandler)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/report/salary-by-location", withCORS(SalaryByLocation)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/report/vacancies-with-skills", withCORS(VacanciesWithSkills)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/report/skill-demand", withCORS(SkillDemand)).Methods("GET", "OPTIONS")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
