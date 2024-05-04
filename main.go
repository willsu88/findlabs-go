package main

import (
	"fmt"
	"log"
	"net/http"
	// "sort"
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

// Handler functions.
func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Info page")
}

func main() {
    log.Println("Starting our simple http server.")

    // Registering our handler functions, and creating paths.
    http.HandleFunc("/", Home)
    http.HandleFunc("/info", Info)

    log.Println("Started on port", portNum)
    fmt.Println("To close connection CTRL+C :-)")

    // Spinning up the server.
    err := http.ListenAndServe(portNum, nil)
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