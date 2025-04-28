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

	// Добавление
	r.HandleFunc("/api/employers", withCORS(CreateEmployer)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/vacancies", withCORS(CreateVacancy)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/skills", withCORS(CreateSkill)).Methods("POST", "OPTIONS")

	// Удаление
	r.HandleFunc("/api/employers/{id}", withCORS(DeleteEmployer)).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/vacancies/{id}", withCORS(DeleteVacancy)).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/skills/{id}", withCORS(DeleteSkill)).Methods("DELETE", "OPTIONS")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
