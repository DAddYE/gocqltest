CREATE KEYSPACE gocqltest WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 2
};

USE gocqltest;

CREATE TABLE test (
  id INT,
  categories LIST<TEXT>,
  name TEXT,
  status BIGINT,
  PRIMARY KEY (id)
);
