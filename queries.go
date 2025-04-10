package main

var queries = map[int]string{
	1:  `SELECT title, salary FROM vacancies ORDER BY salary DESC LIMIT 10;`,
	2:  `SELECT location, AVG(salary) FROM vacancies GROUP BY location;`,
	3:  `SELECT employers.name, count(*) FROM vacancies JOIN employers ON employer_id = employers.id GROUP BY employers.name ORDER BY count(*) DESC;`,
	4:  `SELECT title FROM vacancies JOIN requirements ON requirements.vacancy_id = id          JOIN skills ON skill_id = skills.id WHERE skills.name = 'Python';`,
	5:  `SELECT title, published_date FROM vacancies WHERE EXTRACT(MONTH FROM published_date) = EXTRACT(MONTH FROM now());`,
	6:  `SELECT width_bucket(salary, 0, 500000, 10) AS bucket, CONCAT('$', (width_bucket(salary, 0, 500000, 10) - 1) * 50000, ' - $', width_bucket(salary, 0, 500000, 10) * 50000) AS salary_range, COUNT(*) AS vacancies_count FROM vacancies WHERE salary IS NOT NULL GROUP BY bucket, salary ORDER BY bucket;`,
	7:  `SELECT skills.name, COUNT(*) AS skill_count FROM requirements JOIN skills ON requirements.skill_id = skills.id GROUP BY skills.name ORDER BY skill_count DESC;`,
	8:  `SELECT location, COUNT(*) FROM vacancies GROUP BY location;`,
	9:  `SELECT AVG(now() - published_date) FROM vacancies;`,
	10: `SELECT employers.name, AVG(salary) AS avg_salary FROM vacancies JOIN employers ON employers.id = vacancies.employer_id GROUP BY employers.name ORDER BY avg_salary DESC LIMIT 5;`,
	11: `SELECT s.name AS skill, COUNT(DISTINCT r.vacancy_id) AS vacancy_count FROM skills s JOIN requirements r ON s.id = r.skill_id GROUP BY s.id ORDER BY vacancy_count DESC LIMIT 5;`,
	12: `SELECT v.title, v.location, v.salary FROM vacancies v JOIN (SELECT location, AVG(salary) AS avg_salary FROM vacancies GROUP BY location) avg_table ON v.location = avg_table.location WHERE v.salary > avg_table.avg_salary;`,
	13: `SELECT v.id, v.title, COUNT(r.skill_id) AS skills_required FROM vacancies v JOIN requirements r ON v.id = r.vacancy_id GROUP BY v.id ORDER BY skills_required DESC LIMIT 5;`,
	14: `SELECT id, title, (max_salary - min_salary) AS salary_diff FROM vacancies ORDER BY salary_diff DESC LIMIT 5;`,
	15: `SELECT e.name, AVG(v.closed_date - v.published_date) AS avg_days_to_close FROM vacancies v JOIN employers e ON v.employer_id = e.id WHERE v.closed_date IS NOT NULL GROUP BY e.name ORDER BY avg_days_to_close LIMIT 5;`,
	16: `SELECT location, ROUND(AVG(skill_count), 2) AS avg_skills FROM (SELECT v.id, v.location, COUNT(r.skill_id) AS skill_count FROM vacancies v LEFT JOIN requirements r ON v.id = r.vacancy_id GROUP BY v.id, v.location) AS skill_stats GROUP BY location ORDER BY avg_skills DESC;`,
	17: `SELECT s.name AS skill, ROUND(AVG(v.salary), 2) AS avg_salary FROM skills s JOIN requirements r ON s.id = r.skill_id JOIN vacancies v ON r.vacancy_id = v.id GROUP BY s.name ORDER BY avg_salary DESC;`,
	18: `SELECT e.name AS employer, ROUND(COUNT(*) / COUNT(DISTINCT DATE_TRUNC('month', v.published_date)), 2) AS avg_vacancies_per_month FROM vacancies v JOIN employers e ON v.employer_id = e.id GROUP BY e.name ORDER BY avg_vacancies_per_month DESC LIMIT 5;`,
	19: `SELECT location, (COUNT(*)::NUMERIC / COUNT(DISTINCT employer_id), 2) AS vacancies_per_employer FROM vacancies GROUP BY location ORDER BY vacancies_per_employer DESC LIMIT 5;`,
	20: `SELECT e.name AS employer, ROUND(AVG(skill_count), 2) AS avg_skills_per_vacancy FROM (SELECT v.id, v.employer_id, COUNT(r.skill_id) AS skill_count FROM vacancies v LEFT JOIN requirements r ON v.id = r.vacancy_id GROUP BY v.id, v.employer_id) AS vacancy_skills JOIN employers e ON vacancy_skills.employer_id = e.id GROUP BY e.name ORDER BY avg_skills_per_vacancy DESC LIMIT 5;`,
}
