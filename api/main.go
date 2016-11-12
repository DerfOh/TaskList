// main process
//
// REST APIs with Go and MySql.
//
// Usage:
//
//   # run go server in the background
//   $ go run *.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// dbAddress the api will connect to, is set by flag, defaults to localhost
var dbAddress string

// dbUserName database username
var dbUserName string

// dbPassword the password to be used to connect to the database
var dbPassword string

// dbConnectionURL the url used to perform db connection
var dbConnectionURL string

func main() {

	port := flag.Int("port", 1234, "an int")
	dbAddress := flag.String("dbaddress", "localhost", "a string")
	dbUserName := flag.String("dbuser", "compromise", "a string")
	dbPassword := flag.String("dbpassword", "password", "a string")
	dbConnectionURL = *dbUserName + ":" + *dbPassword + "@tcp(" + *dbAddress + ":3306)/compromise"

	// Execute the command-line parsing.
	flag.Parse()

	// Show flag trail in logs, do not show dbPassword intentionally, just check if the password is the default
	fmt.Println("api port:\t\t", *port)
	fmt.Println("database address:\t", *dbAddress)
	fmt.Println("database user:\t\t", *dbUserName)
	if *dbPassword == "password" {
		fmt.Println("default db password:\t true")
		fmt.Println("connection url:", dbConnectionURL)
	} else {
		fmt.Println("default db password:\t false")
	}
	fmt.Println("Comments, remarks, general thoughts:", flag.Args())

	var err string
	portstring := strconv.Itoa(*port)

	mux := http.NewServeMux()
	//mux.Handle("/api/", http.HandlerFunc(APIHandler))
	mux.Handle("/api/users/", http.HandlerFunc(UserAPIHandler))     // Handler for User interactions
	mux.Handle("/api/tasks/", http.HandlerFunc(TaskAPIHandler))     // Handler for Task interactions
	mux.Handle("/api/rewards/", http.HandlerFunc(RewardAPIHandler)) // Handler for Reward interactions
	mux.Handle("/api/groups/", http.HandlerFunc(GroupAPIHandler))   // Hanlder for Group interactions
	mux.Handle("/api/auth/", http.HandlerFunc(AuthAPIHandler))      // Handler for Authentication of users
	mux.Handle("/", http.HandlerFunc(Handler))

	// Start listing on a given port with these routes on this server.
	log.Print("Listening on port " + portstring + " ... ")
	errs := http.ListenAndServe(":"+portstring, mux)
	if errs != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func cleanJSON(s string) string {
	// fmt.Println(s)
	s = strings.Replace(s, "\\\"", "\"", -1)
	// fmt.Println(s)
	s = strings.Replace(s, "}\"", "}", -1)
	// fmt.Println(s)
	s = strings.Replace(s, "\"{", "{", -1)
	// fmt.Println(s)
	return s
}
