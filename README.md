## TEST NULL's issue

If you have at least one instance of cassandra running just:

```
$ go get github.com/daddye/gocqltest
$ cd $GOPATH/src/github.com/daddye/gocqltest
$ make
```

This command will create drop/create the keyspace for the test.

### The Test:

Seems that `prepared` statements [1](https://github.com/gocql/gocql/issues/296) [2](https://groups.google.com/a/lists.datastax.com/forum/#!topic/java-driver-user/cHE3OOSIXBU/discussion) [3](https://issues.apache.org/jira/browse/CASSANDRA-7304)
(defaults in gocql) suffers of a "bug" where *not* setting a
*column* means setting it `NULL` which (seems only with prepared statements) causes the creation of
a tombstone in cassandra.

This problem _seems_ amplified with _arrays_ (slices). Basically an empty slice is `nil` or `null` for cassandra.

The output will be:

```
cqlsh -e "DROP KEYSPACE IF EXISTS gocqltest"
cqlsh -f schema.cql
go run main.go
cqlsh -e "SELECT * FROM gocqltest.test"

 id | categories | name | status
----+------------+------+--------
  5 |       null | null |   null
 10 |       null | null |   null
 16 |       null | null |   null
 13 |       null | null |   null
 11 |       null | null |   null
  1 |       null | null |   null
 19 |       null | null |   null
  8 |       null | null |   null
  0 |       null | null |   null
  2 |       null | null |   null
  4 |       null | null |   null
 18 |       null | null |   null
 15 |       null | null |   null
  7 |       null | null |   null
  6 |       null | null |   null
  9 |       null | null |   null
 14 |       null | null |   null
 17 |       null | null |   null
 12 |       null | null |   null
  3 |       null | null |   null

(20 rows)
```

Which according to [jira](https://issues.apache.org/jira/browse/CASSANDRA-7304) would create
tombstones (probably when you use more than one node) in cassandra since gocql comes with prepared
statements by default.
