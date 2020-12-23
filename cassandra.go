package main

import (
	"fmt"
	"github.com/gocql/gocql"
)

func main() {
    var cqlshrc_host = "10.101.33.41"
    var _username = "cassandra"
    var _password = "cassandra"

    config := gocql.NewCluster(cqlshrc_host)
    config.Authenticator = gocql.PasswordAuthenticator{
        Username: _username,
        Password: _password,
    }
    

    fmt.Println("starting..")
    session, _ := config.CreateSession()
    
    defer session.Close()
    var _version string
    var _query = "select release_version from system.local"
    iter := session.Query(_query).Iter()
    for iter.Scan(&_version) {
        fmt.Printf("Version: %s\n", _version)
    }
    iter.Close()

}