CREATE TABLE teams (
  id integer PRIMARY KEY NOT NULL unique,
  name varchar(48) NOT NULL unique,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);
