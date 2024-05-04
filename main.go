package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type ContractType string
type Status string
type Order string
type SortBy string

const portNum string = ":8080"
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
	Name         string
	Address      string
	Transaction  string
	Block        int `json:",string"`
	ContractType ContractType
	Status       Status
}

var contracts []Contract

// Handler functions.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func getContracts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contracts)
}

func main() {
	initTestLocalData()

	db, err := initDB()
	if err != nil {
		fmt.Printf("failed to init db")
		return
	}

	fmt.Println(db)
	// names, err := queryContracts(db)
	// if err != nil {
	// 	fmt.Printf("failed to query DB")
	//     return
	// }

	// fmt.Println("The iterated elements are:")
	// for i, item := range names {
	//    fmt.Println(i, "--", item)
	// }

	startHttpServer()
}

func initTestLocalData() {
	c1 := Contract{"abc", "addres", "0x111", 123, "deployed", "ok"}
	c2 := Contract{"abc", "addres", "0x111", 123, "deployed", "ok"}

	contracts = append(contracts, c1)
	contracts = append(contracts, c2)
}

func initDB() (*sql.DB, error) {
	connStr := "postgres://postgres:password123@localhost/postgres?sslmode=disable"
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
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
	return db, err;
}

func queryContracts(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT * FROM Contracts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	return names, nil
}

func startHttpServer() {
	log.Println("Starting our simple http server.")
	router := mux.NewRouter()

	// Registering our handler functions, and creating paths.
	router.HandleFunc("/", Home)
	router.HandleFunc("/itemsLocal", getContracts).Methods("GET")

	log.Println("Started on port", portNum)
	fmt.Println("To close connection CTRL+C :-)")

	// Spinning up the server.
	err := http.ListenAndServe(portNum, router)
	if err != nil {
		log.Fatal(err)
	}
}

// func sortByName(contracts []Contract, descending bool) {
// 	if descending {
// 		sort.Slice(contracts[:], func(i, j int) bool {
// 			return contracts[i].Name > contracts[j].Name
// 		})
// 	} else {
// 		sort.Slice(contracts[:], func(i, j int) bool {
// 			return contracts[i].Name < contracts[j].Name
// 		})
// 	}
// }

// func sortByAddress(contracts []Contract, descending bool) {
// 	if descending {
// 		sort.Slice(contracts[:], func(i, j int) bool {
// 			return contracts[i].Address > contracts[j].Address
// 		})
// 	} else {
// 		sort.Slice(contracts[:], func(i, j int) bool {
// 			return contracts[i].Address < contracts[j].Address
// 		})
// 	}
// }

// func sortByTransaction(contracts []Contract, descending bool) {
// 	if descending {
// 		sort.Slice(contracts[:], func(i, j int) bool {
// 			return contracts[i].Transaction > contracts[j].Transaction
// 		})
// 	} else {
// 		sort.Slice(contracts[:], func(i, j int) bool {
// 			return contracts[i].Transaction < contracts[j].Transaction
// 		})
// 	}
// }

// func main() {
// 	fmt.Println("hello world")

// 	var data JsonPayload
// 	file, err := ioutil.ReadFile("FlowJson.json")
// 	if err != nil {
// 		fmt.Printf("failed to read json file, error: %v", err)
// 		return
// 	}

// 	err = json.Unmarshal([]byte(file), &data)
// 	if err != nil {
// 		fmt.Println("failed to unmarshall file:", err)
// 		return
// 	}

// 	fmt.Println("Sort by:", data.SortBy)
// 	fmt.Println("Sort by:", data.Descending)

// var myMap = map[string]func([]Contract, bool){
// 	"name":        sortByName,
// 	"address":     sortByAddress,
// 	"transaction": sortByTransaction,
// }
// sortFunc := myMap[data.SortBy]
// sortFunc(data.Contracts, data.Descending)
// fmt.Println(data.SortBy, data.Descending)
// fmt.Println("Contracts:", data.Contracts)
// }
