package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"sort"
)

type ContractType string
type Status string
type Order string
type SortBy string

const portNum string = ":8080"
const userName string = "postgres"
const password string = "password123"

const (
	Ok    Status = "ok"
	Error Status = "error"
)

const (
	Deployed ContractType = "deployed"
	Updated  ContractType = "updated"
)

const (
	Name        SortBy = "name"
	Address     SortBy = "address"
	Transaction SortBy = "transaction"
)

type JsonPayload struct {
	SortBy     string
	Descending bool `json:",string"`
}

type Contract struct {
	Name         string       `db:"name"`
	Address      string       `db:"address"`
	Transaction  string       `db:"transaction_id"`
	Block        int          `db:"block"`
	ContractType ContractType `db:"contractType"`
	Status       Status       `db:"status"`
}

var contracts []Contract

// Handler functions.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func getContracts(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Contracts function")

	json.NewEncoder(w).Encode(contracts)
}

func getContractsSort(w http.ResponseWriter, r *http.Request) {
	log.Println("Into Sorting function")
	vars := mux.Vars(r)

	sortType := vars["sort_type"]
	isDescending := false
	if vars["sort"] == "asc" {
		isDescending = true
	}
	
	var myMap = map[string]func([]Contract, bool){
		"name":        sortByName,
		"address":     sortByAddress,
		"transaction": sortByTransaction,
	}
	sortFunc := myMap[sortType]
	sortFunc(contracts, isDescending)
	json.NewEncoder(w).Encode(contracts)
}

func startHttpServer() {
	log.Println("Starting our simple http server.")
	router := mux.NewRouter()

	// Registering our handler functions, and creating paths.
	router.HandleFunc("/", Home)
	router.HandleFunc("/contracts", getContracts).Methods("GET")
	router.HandleFunc("/contracts/{sort_type}/sort:{sort}", getContractsSort).Methods("GET")

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe(portNum, router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db, err := initDB()
	if err != nil {
		fmt.Printf("failed to init db")
		return
	}

	names, err := queryContracts(db)
	if err != nil {
		fmt.Println("failed to query DB")
		return
	}

	for i, item := range names {
		fmt.Println(i, "--", item)
	}

	startHttpServer()
}

func initDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/postgres?sslmode=disable", userName, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("failed here", err)
		return nil, err
	}

	if err != nil {

		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	
	fmt.Println("The database is connected")
	return db, err
}

func queryContracts(db *sql.DB) ([]Contract, error) {
	rows, err := db.Query("SELECT name, address, transaction_id, block, contractType, status FROM Contracts")

	if err != nil {
		fmt.Printf("failed to fire query")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data Contract
		if err := rows.Scan(&data.Name, &data.Address, &data.Transaction, &data.Block, &data.ContractType, &data.Status); err != nil {
			fmt.Println("failed to scan row")
			return nil, err
		}

		contracts = append(contracts, data)
	}

	return contracts, nil
}

func sortByName(contracts []Contract, descending bool) {
	if descending {
		sort.Slice(contracts[:], func(i, j int) bool {
			return contracts[i].Name > contracts[j].Name
		})
	} else {
		sort.Slice(contracts[:], func(i, j int) bool {
			return contracts[i].Name < contracts[j].Name
		})
	}
}

func sortByAddress(contracts []Contract, descending bool) {
	if descending {
		sort.Slice(contracts[:], func(i, j int) bool {
			return contracts[i].Address > contracts[j].Address
		})
	} else {
		sort.Slice(contracts[:], func(i, j int) bool {
			return contracts[i].Address < contracts[j].Address
		})
	}
}

func sortByTransaction(contracts []Contract, descending bool) {
	if descending {
		sort.Slice(contracts[:], func(i, j int) bool {
			return contracts[i].Transaction > contracts[j].Transaction
		})
	} else {
		sort.Slice(contracts[:], func(i, j int) bool {
			return contracts[i].Transaction < contracts[j].Transaction
		})
	}
}

