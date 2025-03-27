package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	_ "github.com/lib/pq"
)

type LocationSalary struct {
	Location  string
	AvgSalary float64
}

func main() {
	// Подключение к PostgreSQL
	connStr := "user=feoram password=swedaqws dbname=itjobs host=172.17.0.2 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	// SQL-запрос на среднюю зарплату по локациям
	query := `
		SELECT location, AVG(salary) AS avg_salary
		FROM vacancies
		GROUP BY location
		ORDER BY avg_salary DESC;
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Ошибка выполнения запроса:", err)
	}
	defer rows.Close()

	// Сбор данных в структуру
	var results []LocationSalary
	for rows.Next() {
		var loc LocationSalary
		if err := rows.Scan(&loc.Location, &loc.AvgSalary); err != nil {
			log.Fatal("Ошибка сканирования строки:", err)
		}
		results = append(results, loc)
	}

	// Генерация HTML-графика
	generateChart(results)
}

// Строим график с помощью go-echarts
func generateChart(data []LocationSalary) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Средняя зарплата по локациям"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Локация"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Средняя зарплата"}),
	)

	// Наполняем оси
	locations := make([]string, 0)
	salaries := make([]opts.BarData, 0)

	for _, item := range data {
		locations = append(locations, item.Location)
		salaries = append(salaries, opts.BarData{Value: item.AvgSalary})
	}

	bar.SetXAxis(locations).
		AddSeries("Зарплата", salaries)

	// Сохраняем HTML-отчёт
	f, err := os.Create("avg_salary_by_location.html")
	if err != nil {
		log.Fatal("Ошибка создания HTML-файла:", err)
	}
	defer f.Close()

	bar.Render(f)
	fmt.Println("Отчёт создан: avg_salary_by_location.html")
}
