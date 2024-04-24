CREATE DATABASE IF NOT EXISTS hexagonal;

USE hexagonal;

CREATE TABLE IF NOT EXISTS arith_history (
  date DATE NOT NULL,
  answer INT NOT NULL,
  operation VARCHAR(15) NOT NULL
);
