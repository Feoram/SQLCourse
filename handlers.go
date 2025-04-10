package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // или "http://localhost:5173"
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		h(w, r)
	}
}

func SalaryByLocation(w http.ResponseWriter, r *http.Request) {
	min := r.URL.Query().Get("min")
	max := r.URL.Query().Get("max")

	query := `
		SELECT location, AVG(salary) AS avg_salary
		FROM vacancies
		WHERE salary IS NOT NULL
	`

	if min != "" {
		query += " AND salary >= " + min
	}
	if max != "" {
		query += " AND salary <= " + max
	}

	query += " GROUP BY location ORDER BY avg_salary DESC"

	rows, err := DB.Query(query)
	if err != nil {
		http.Error(w, "Ошибка SQL: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	respondJSON(w, rows)
}

func VacanciesWithSkills(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	skill := r.URL.Query().Get("skill")

	query := `
		SELECT v.title, e.name AS employer, s.name AS skill, v.location
		FROM vacancies v
		JOIN employers e ON v.employer_id = e.id
		JOIN requirements r ON v.id = r.vacancy_id
		JOIN skills s ON r.skill_id = s.id
		WHERE 1=1
	`

	if location != "" {
		query += " AND v.location = '" + location + "'"
	}
	if skill != "" {
		query += " AND s.name ILIKE '%" + skill + "%'"
	}

	query += " ORDER BY v.title, s.name"

	rows, err := DB.Query(query)
	if err != nil {
		http.Error(w, "Ошибка SQL: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	respondJSON(w, rows)
}

func SkillDemand(w http.ResponseWriter, r *http.Request) {
	min := r.URL.Query().Get("min_count")
	query := `
		SELECT s.name, COUNT(*) AS demand
		FROM requirements r
		JOIN skills s ON r.skill_id = s.id
		GROUP BY s.name
		HAVING COUNT(*) >= 1
	`

	if min != "" {
		query = strings.Replace(query, ">= 1", ">= "+min, 1)
	}

	query += " ORDER BY demand DESC"

	rows, err := DB.Query(query)
	if err != nil {
		http.Error(w, "Ошибка SQL: "+err.Error(), 500)
		return
	}
	defer rows.Close()

	respondJSON(w, rows)
}

func respondJSON(w http.ResponseWriter, rows *sql.Rows) {
	cols, _ := rows.Columns()
	values := make([]interface{}, len(cols))
	ptrs := make([]interface{}, len(cols))

	result := []map[string]interface{}{}

	for rows.Next() {
		for i := range ptrs {
			ptrs[i] = &values[i]
		}
		rows.Scan(ptrs...)
		row := make(map[string]interface{})
		for i, col := range cols {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}
		result = append(result, row)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || queries[id] == "" {
		http.Error(w, "Неверный номер запроса", http.StatusBadRequest)
		return
	}

	rows, err := DB.Query(queries[id])
	if err != nil {
		log.Printf("Ошибка запроса #%d: %v", id, err)
		http.Error(w, "Ошибка выполнения запроса", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	values := make([]interface{}, len(cols))
	ptrs := make([]interface{}, len(cols))

	result := []map[string]interface{}{}

	for rows.Next() {
		for i := range ptrs {
			ptrs[i] = &values[i]
		}
		rows.Scan(ptrs...)
		row := make(map[string]interface{})
		for i, col := range cols {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}
		result = append(result, row)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}
