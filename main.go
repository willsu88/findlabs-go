package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"slices"
	"strings"
)

const portNum string = ":8080"
var userName string = os.Getenv("postgresUser")
var password string = os.Getenv("postgresPassword")
var database *sql.DB

type Contract struct {
	Name         string `db:"name"`
	Address      string `db:"address"`
	Transaction  string `db:"transaction_id"`
	Block        int    `db:"block"`
	ContractType string `db:"contractType"`
	Status       string `db:"status"`
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Printf("failed to init db")
		return
	}
	database = db

	startHttpServer()
}

/* ---------- HTTP and Handler functions ----------  */
func startHttpServer() {
	router := initHttpHandler()

	// Spinning up the server.
	log.Println("Started on port", portNum)
	log.Println("To close connection CTRL+C :-)")
	err := http.ListenAndServe(portNum, router)
	if err != nil {
		log.Fatal(err)
	}
}

func initHttpHandler() http.Handler {
    router := mux.NewRouter()

	// Registering our handler functions, and creating paths.
	router.Path("/contracts").HandlerFunc(getAllContracts).Methods("GET")
	router.Path("/contracts/sort").Queries("reverse", "{is_reverse}").HandlerFunc(getContractsSort).Methods("GET")
	router.Path("/contracts/name/{name}").HandlerFunc(getContractByNamePrefix).Methods("GET")
	return router
}

func getAllContracts(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Contracts function")
	contracts, err := queryContractsfromDB(database)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
    	w.Write([]byte("404 - Error querying the database!"))
		return
	}
	json.NewEncoder(w).Encode(contracts)
}

func getContractByNamePrefix(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Contract by Name function")
	contracts, err := queryContractsfromDB(database)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
    	w.Write([]byte("404 - Error querying the database!"))
		return
	}

	vars := mux.Vars(r)
	nameString := vars["name"]
	log.Println(nameString)
	contracts = filter(contracts, nameString)
	json.NewEncoder(w).Encode(contracts)
}

func getContractsSort(w http.ResponseWriter, r *http.Request) {
	log.Println("Sort Contracts Based on Name")
	contracts, err := queryContractsfromDB(database)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Error querying the database!"))
		return
	}
	
	var isReverse bool
	vars := mux.Vars(r)
	isReverseString, exist := vars["is_reverse"]
	if exist {
		isReverse, _ = strconv.ParseBool(isReverseString)
	} else {
		isReverse = false
	}

	sort.Slice(contracts[:], func(i, j int) bool {
		return contracts[i].Name < contracts[j].Name
	})

	if isReverse { slices.Reverse(contracts)}
	json.NewEncoder(w).Encode(contracts)
}

/* ---------- Database functions ----------  */
func initDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/postgres?sslmode=disable", userName, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("failed here", err)
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	log.Println("The database is connected")
	return db, err
}

func queryContractsfromDB(db *sql.DB) ([]Contract, error) {
	rows, err := db.Query("SELECT name, address, transaction_id, block, contractType, status FROM Contracts")

	if err != nil {
		log.Printf("failed to fire query")
		return nil, err
	}
	defer rows.Close()

	var contracts []Contract
	for rows.Next() {
		var data Contract
		err := rows.Scan(&data.Name, &data.Address, &data.Transaction, &data.Block, &data.ContractType, &data.Status)
		if err != nil {
			log.Println("failed to scan row")
			return nil, err
		}
		contracts = append(contracts, data)
	}

	return contracts, nil
}

// Sorting / Filtering Functions
func filter(contracts []Contract, prefix string) []Contract {
	var out []Contract
	for _, contract := range contracts {
	   if strings.HasPrefix(contract.Name,prefix) {
		  out = append(out, contract)
	   }
	}
	return out
 }