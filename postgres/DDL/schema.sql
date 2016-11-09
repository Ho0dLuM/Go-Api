DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id serial PRIMARY KEY,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL
);

\i /Users/Ho0dLuM/GOspace/src/github.com/ho0dlum/leagueapi/postgres/DML/users.sql
