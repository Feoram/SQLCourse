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

type SkillDemand struct {
	Skill string
	Count int
}

func main() {
	connStr := "user=feoram password=swedaqws dbname=itjobs host=172.17.0.2 port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT s.name, COUNT(*) as demand
		FROM requirements r
		JOIN skills s ON r.skill_id = s.id
		GROUP BY s.name
		ORDER BY demand DESC;
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var skills []SkillDemand
	for rows.Next() {
		var s SkillDemand
		if err := rows.Scan(&s.Skill, &s.Count); err != nil {
			log.Fatal(err)
		}
		skills = append(skills, s)
	}

	generateSkillChart(skills)
}

func generateSkillChart(data []SkillDemand) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Востребованность навыков"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Навык"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Количество вакансий"}),
	)

	skillNames := make([]string, 0)
	counts := make([]opts.BarData, 0)

	for _, s := range data {
		skillNames = append(skillNames, s.Skill)
		counts = append(counts, opts.BarData{Value: s.Count})
	}

	bar.SetXAxis(skillNames).AddSeries("Навыки", counts)

	f, err := os.Create("skill_demand_chart.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bar.Render(f)
	fmt.Println("Диаграмма создана: skill_demand_chart.html")
}
