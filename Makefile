SHELL:=/bin/bash
PATH:=$(PATH):/usr/local/cassandra/bin

default: test

test: setup
	go run main.go
	cqlsh -e "SELECT * FROM gocqltest.test"
	nodetool cfstats gocqltest

setup:
	cqlsh -e "DROP KEYSPACE IF EXISTS gocqltest"
	cqlsh -f schema.cql

.PHONY: test setup
