-- 1. Топ-10 вакансий по уровню зарплаты.
SELECT title, salary
FROM vacancies
ORDER BY salary DESC
LIMIT 10;

-- 2. Средняя зарплата для каждой локации.
SELECT location, AVG(salary)
FROM vacancies
GROUP BY location;

-- 3. Вывести компании с наибольшим количеством открытых вакансий.
SELECT employers.name, count(*)
FROM vacancies
         JOIN employers ON employer_id = employers.id
GROUP BY employers.name
ORDER BY count(*) DESC;

-- 4. Найти вакансии, требующие определенного навыка (например, Python).
SELECT title
FROM vacancies
         JOIN requirements ON requirements.vacancy_id = id
         JOIN skills ON skill_id = skills.id
WHERE skills.name = 'Python';

-- 5. Список вакансий, опубликованных последний месяц.
SELECT title, published_date
FROM vacancies
WHERE EXTRACT(MONTH FROM published_date) = EXTRACT(MONTH FROM now());

-- 6. Распределение вакансий по категориям зарплат (гистограмма).
SELECT width_bucket(salary, 0, 500000, 10) AS bucket,
       CONCAT(
               '$',
               (width_bucket(salary, 0, 500000, 10) - 1) * 50000,
               ' - $',
               width_bucket(salary, 0, 500000, 10) * 50000
       )                                   AS salary_range,
       COUNT(*)                            AS vacancies_count
FROM vacancies
WHERE salary IS NOT NULL
GROUP BY bucket, salary
ORDER BY bucket;


-- 7. Список навыков, наиболее часто встречающихся в вакансиях.
SELECT skills.name, COUNT(*) AS skill_count
FROM requirements
         JOIN skills ON requirements.skill_id = skills.id
GROUP BY skills.name
ORDER BY skill_count DESC;

-- 8. Количество вакансий в каждой локации.
SELECT location, COUNT(*)
FROM vacancies
GROUP BY location;

-- 9. Среднее время существования вакансий на рынке.
SELECT AVG(now() - published_date)
FROM vacancies;

-- 10. Топ-5 компаний с самым высоким средним уровнем зарплат.
SELECT employers.name, AVG(salary) AS avg_salary
FROM vacancies
         JOIN employers ON employers.id = vacancies.employer_id
GROUP BY employers.name
ORDER BY avg_salary DESC
LIMIT 5;

-- 11. Топ-5 самых востребованных навыков (по количеству вакансий, где они требуются).
SELECT s.name AS skill, COUNT(DISTINCT r.vacancy_id) AS vacancy_count FROM skills s JOIN requirements r ON s.id = r.skill_id GROUP BY s.id ORDER BY vacancy_count DESC LIMIT 5;


-- 12. Вакансии с зарплатой выше средней в их локации.
SELECT v.title, v.location, v.salary FROM vacancies v JOIN (SELECT location, AVG(salary) AS avg_salary FROM vacancies GROUP BY location) avg_table ON v.location = avg_table.location WHERE v.salary > avg_table.avg_salary;

-- 13. Вакансии с самыми длинными требованиями (по количеству нужных навыков).
SELECT v.id, v.title, COUNT(r.skill_id) AS skills_required FROM vacancies v JOIN requirements r ON v.id = r.vacancy_id GROUP BY v.id ORDER BY skills_required DESC LIMIT 5;

-- 14. Вакансии с наибольшим разбросом зарплат (разница между min/max зарплатой).
-- Предположим, что есть поля min_salary и max_salary
SELECT id, title, (max_salary - min_salary) AS salary_diff FROM vacancies ORDER BY salary_diff DESC LIMIT 5;


-- 15. Какие компании быстрее всех закрывают вакансии (по среднему времени закрытия)?
-- Требуется поле closed_date: DATE
SELECT e.name, AVG(v.closed_date - v.published_date) AS avg_days_to_close FROM vacancies v JOIN employers e ON v.employer_id = e.id WHERE v.closed_date IS NOT NULL GROUP BY e.name ORDER BY avg_days_to_close LIMIT 5;

-- 16. Среднее количество навыков, требуемых в вакансиях по каждой локации.
SELECT location, ROUND(AVG(skill_count), 2) AS avg_skills FROM (SELECT v.id, v.location, COUNT(r.skill_id) AS skill_count FROM vacancies v LEFT JOIN requirements r ON v.id = r.vacancy_id GROUP BY v.id, v.location) AS skill_stats GROUP BY location ORDER BY avg_skills DESC;


-- 17. Средняя зарплата по навыкам (какие технологии дороже?).
SELECT s.name AS skill, ROUND(AVG(v.salary), 2) AS avg_salary FROM skills s JOIN requirements r ON s.id = r.skill_id JOIN vacancies v ON r.vacancy_id = v.id GROUP BY s.name ORDER BY avg_salary DESC;


-- 18. Какие вакансии быстрее всего публикуют компании (среднее число новых вакансий в месяц)?
SELECT e.name AS employer, ROUND(COUNT(*) / COUNT(DISTINCT DATE_TRUNC('month', v.published_date)), 2) AS avg_vacancies_per_month FROM vacancies v JOIN employers e ON v.employer_id = e.id GROUP BY e.name ORDER BY avg_vacancies_per_month DESC LIMIT 5;


-- 19. 5 городов с самой высокой конкуренцией (самое большое число вакансий на одного работодателя).
SELECT location, (COUNT(*)::NUMERIC / COUNT(DISTINCT employer_id), 2) AS vacancies_per_employer FROM vacancies GROUP BY location ORDER BY vacancies_per_employer DESC LIMIT 5;


-- 20. В каких компаниях работают специалисты с самым большим количеством навыков?
SELECT e.name AS employer, ROUND(AVG(skill_count), 2) AS avg_skills_per_vacancy FROM (SELECT v.id, v.employer_id, COUNT(r.skill_id) AS skill_count FROM vacancies v LEFT JOIN requirements r ON v.id = r.vacancy_id GROUP BY v.id, v.employer_id) AS vacancy_skills JOIN employers e ON vacancy_skills.employer_id = e.id GROUP BY e.name ORDER BY avg_skills_per_vacancy DESC LIMIT 5;
