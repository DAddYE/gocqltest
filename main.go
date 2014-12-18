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

	// test null values
	for i := 0; i < 10; i++ {
		err := session.Query("INSERT INTO test (id) VALUES (?)", i).Exec()
		if err != nil {
			log.Fatalf("first update: %s", err)
		}
	}

	// test empty slice
	for i := 10; i < 20; i++ {
		err := session.Query("INSERT INTO test (id, categories) VALUES (?, ?)", i, &[]string{}).Exec()
		if err != nil {
			log.Fatalf("first update: %s", err)
		}
	}
}
