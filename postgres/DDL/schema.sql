DROP TABLE IF EXISTS champions;

CREATE TABLE champions (
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  image varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  enemytips varchar(255) NOT NULL,
  lore varchar(255) NOT NULL,
  passive_name varchar(255) NOT NULL,
  passive_image varchar(255) NOT NULL,
  passive_description varchar(255) NOT NULL,
  spells_q_name varchar(255) NOT NULL,
  spells_q_image varchar(255) NOT NULL,
  spells_q_description varchar(255) NOT NULL,
  spells_w_name varchar(255) NOT NULL,
  spells_w_image varchar(255) NOT NULL,
  spells_w_description varchar(255) NOT NULL,
  spells_e_name varchar(255) NOT NULL,
  spells_e_image varchar(255) NOT NULL,
  spells_e_description varchar(255) NOT NULL,
  spells_r_name varchar(255) NOT NULL,
  spells_r_image varchar(255) NOT NULL,
  spells_r_description varchar(255) NOT NULL
);

\i /Users/Ho0dLuM/GOspace/src/github.com/ho0dlum/leagueapi/postgres/DML/champions.sql
