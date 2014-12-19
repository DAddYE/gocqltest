package main

import (
	"log"

	"github.com/gocql/gocql"
)

func main() {
	c := gocql.NewCluster("127.0.0.1")
	c.Consistency = gocql.LocalOne
	c.Keyspace = "gocqltest"
	session, err := c.CreateSession()
	if err != nil {
		log.Fatalf("error while creating session: %s", err)
	}
	defer session.Close()

	const rounds = 1000

	// fill before with values
	for i := 0; i < rounds; i++ {
		err := session.Query(`
			INSERT INTO test (id, status, name, categories) VALUES (?, ?, ?, ?)
		`, i, 1, "test", []string{"test"},
		).Exec()
		if err != nil {
			log.Fatalf("insert: %s", err)
		}
	}

	// test null values
	for i := 0; i < rounds; i++ {
		err := session.Query("INSERT INTO test (id) VALUES (?)", i).Exec()
		if err != nil {
			log.Fatalf("first update: %s", err)
		}
	}

	// test empty slice
	for i := 0; i < rounds; i++ {
		err := session.Query("INSERT INTO test (id, categories) VALUES (?, ?)", i, &[]string{}).Exec()
		if err != nil {
			log.Fatalf("second update: %s", err)
		}
	}
}
