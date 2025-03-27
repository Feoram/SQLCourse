package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/template"

	_ "github.com/lib/pq"
)

type VacancySkill struct {
	Vacancy  string
	Employer string
	Skill    string
}

func main() {
	connStr := "user=feoram password=swedaqws dbname=itjobs host=172.17.0.2 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT v.title, e.name, s.name
		FROM vacancies v
		JOIN employers e ON v.employer_id = e.id
		JOIN requirements r ON v.id = r.vacancy_id
		JOIN skills s ON r.skill_id = s.id
		ORDER BY v.title, s.name;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []VacancySkill
	for rows.Next() {
		var vs VacancySkill
		if err := rows.Scan(&vs.Vacancy, &vs.Employer, &vs.Skill); err != nil {
			log.Fatal(err)
		}
		results = append(results, vs)
	}

	createHTMLTable(results)
}

func createHTMLTable(data []VacancySkill) {
	const tpl = `
<!DOCTYPE html>
<html>
<head>
	<title>Вакансии и навыки</title>
	<style>
		table { border-collapse: collapse; width: 100%; }
		th, td { border: 1px solid #ccc; padding: 8px; text-align: left; }
		th { background-color: #f2f2f2; }
	</style>
</head>
<body>
	<h2>Таблица вакансий, компаний и требуемых навыков</h2>
	<table>
		<tr>
			<th>Вакансия</th>
			<th>Компания</th>
			<th>Навык</th>
		</tr>
		{{range .}}
		<tr>
			<td>{{.Vacancy}}</td>
			<td>{{.Employer}}</td>
			<td>{{.Skill}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>`

	t, err := template.New("report").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("vacancies_skills.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	t.Execute(f, data)
	fmt.Println("HTML-таблица создана: vacancies_skills.html")
}
