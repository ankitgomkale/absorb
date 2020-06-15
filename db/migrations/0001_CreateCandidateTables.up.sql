CREATE TABLE candidates_test (
    id serial PRIMARY KEY,
    name VARCHAR (100) NOT NULL,
    email VARCHAR (100) UNIQUE NOT NULL,
    phone_number VARCHAR (100) NOT NULL,
    interested_role VARCHAR (100) NOT NULL,
    years_of_exp VARCHAR (20) NOT NULL,
    candidate_location VARCHAR (50) NOT NULL,
    last_org VARCHAR (100) NOT NULL,
    notable_org VARCHAR (100),
    college_name VARCHAR (100),
    remote VARCHAR (10),
    relocate VARCHAR (10),
    description VARCHAR (5000) NOT NULL,
    salary integer NOT NULL,
    entrepreneur VARCHAR (10),
    communication_method VARCHAR (100),
    resume_link VARCHAR (500),
    website VARCHAR (500),
    github VARCHAR (500),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

