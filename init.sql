-- Создание таблицы работодателей
CREATE TABLE employers
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    contact_info TEXT
);

-- Заполнение таблицы работодателей (10 записей)
INSERT INTO employers (name, contact_info)
VALUES ('Google', 'contact@google.com'),
       ('Microsoft', 'contact@microsoft.com'),
       ('Amazon', 'contact@amazon.com'),
       ('Facebook', 'contact@facebook.com'),
       ('Apple', 'contact@apple.com'),
       ('Netflix', 'contact@netflix.com'),
       ('Tesla', 'contact@tesla.com'),
       ('IBM', 'contact@ibm.com'),
       ('Adobe', 'contact@adobe.com'),
       ('Samsung', 'contact@samsung.com');

-- Создание таблицы вакансий
CREATE TABLE vacancies
(
    id             SERIAL PRIMARY KEY,
    title          VARCHAR(255) NOT NULL,
    employer_id    INT REFERENCES employers (id) ON DELETE CASCADE,
    salary         NUMERIC(10, 2),
    location       VARCHAR(255),
    published_date DATE         NOT NULL,
    closed_date    DATE,         -- Добавляем поле закрытия вакансии
    min_salary     NUMERIC(10, 2), -- Добавляем минимальную зарплату
    max_salary     NUMERIC(10, 2)  -- Добавляем максимальную зарплату
);

-- Заполнение таблицы вакансий (15 записей)
INSERT INTO vacancies (title, employer_id, salary, location, published_date, min_salary, max_salary)
VALUES ('Software Engineer', 1, 120000, 'San Francisco', '2025-01-01', 110000, 130000),
       ('Data Scientist', 2, 110000, 'New York', '2025-01-02', 100000, 120000),
       ('DevOps Engineer', 3, 115000, 'Seattle', '2025-01-03', 105000, 125000),
       ('Backend Developer', 4, 105000, 'Los Angeles', '2025-01-04', 95000, 115000),
       ('Frontend Developer', 5, 100000, 'Austin', '2025-01-05', 90000, 110000),
       ('Cybersecurity Analyst', 6, 95000, 'Boston', '2025-01-06', 85000, 105000),
       ('AI Engineer', 7, 130000, 'Palo Alto', '2025-01-07', 120000, 140000),
       ('Cloud Architect', 8, 125000, 'Chicago', '2025-01-08', 115000, 135000),
       ('Database Administrator', 9, 98000, 'San Diego', '2025-01-09', 90000, 105000),
       ('Product Manager', 10, 140000, 'San Jose', '2025-01-10', 130000, 150000),
       ('Game Developer', 1, 90000, 'San Francisco', '2025-01-11', 80000, 100000),
       ('Machine Learning Engineer', 2, 125000, 'New York', '2025-01-12', 115000, 135000),
       ('Full Stack Developer', 3, 110000, 'Seattle', '2025-01-13', 100000, 120000),
       ('Business Analyst', 4, 85000, 'Los Angeles', '2025-01-14', 75000, 95000),
       ('IT Support Specialist', 5, 70000, 'Austin', '2025-01-15', 65000, 75000);

-- Создание таблицы навыков
CREATE TABLE skills
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- Заполнение таблицы навыков (10 записей)
INSERT INTO skills (name)
VALUES ('Python'),
       ('JavaScript'),
       ('AWS'),
       ('Docker'),
       ('SQL'),
       ('Machine Learning'),
       ('Cybersecurity'),
       ('Cloud Computing'),
       ('C++'),
       ('Product Management');

-- Создание таблицы требований (связь вакансий и навыков)
CREATE TABLE requirements
(
    vacancy_id INT REFERENCES vacancies (id) ON DELETE CASCADE,
    skill_id   INT REFERENCES skills (id) ON DELETE CASCADE,
    PRIMARY KEY (vacancy_id, skill_id)
);

-- Заполнение таблицы требований (15 записей)
INSERT INTO requirements (vacancy_id, skill_id)
VALUES (1, 1),
       (1, 2),  -- Software Engineer -> Python, JavaScript
       (2, 1),
       (2, 5),  -- Data Scientist -> Python, SQL
       (3, 3),
       (3, 4),  -- DevOps Engineer -> AWS, Docker
       (4, 2),
       (4, 5),  -- Backend Developer -> JavaScript, SQL
       (5, 2),
       (5, 4),  -- Frontend Developer -> JavaScript, Docker
       (6, 7),
       (6, 5),  -- Cybersecurity Analyst -> Cybersecurity, SQL
       (7, 1),
       (7, 6),  -- AI Engineer -> Python, Machine Learning
       (8, 3),
       (8, 8),  -- Cloud Architect -> AWS, Cloud Computing
       (9, 5),
       (9, 4),  -- Database Administrator -> SQL, Docker
       (10, 10),
       (10, 1), -- Product Manager -> Product Management, Python
       (11, 2),
       (11, 9), -- Game Developer -> JavaScript, C++
       (12, 6),
       (12, 1), -- ML Engineer -> Machine Learning, Python
       (13, 2),
       (13, 5), -- Full Stack Developer -> JavaScript, SQL
       (14, 5),
       (14, 8), -- Business Analyst -> SQL, Cloud Computing
       (15, 4),
       (15, 7); -- IT Support Specialist -> Docker, Cybersecurity;

-- Добавление поля closed_date для учета времени закрытия вакансии
ALTER TABLE vacancies
    ADD COLUMN closed_date DATE;
