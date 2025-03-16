CREATE TABLE vacancies (
    id SERIAL primary key,
    email VARCHAR(255) not null,
    role VARCHAR(255),
    company VARCHAR(255),
    salary VARCHAR(255),
    type VARCHAR(255),
    location VARCHAR(255)
)