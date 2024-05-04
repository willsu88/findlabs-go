package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"io/ioutil"
)

type ContractType string
type Status string
type Order string

const (
	Ok    Status = "ok"
	Error Status = "error"
)

const (
	Deployed ContractType = "deployed"
	Updated  ContractType = "updated"
)

const (
	Ascend Order = "ascend"
	Descend  Order = "descend"
)

type JsonPayload struct {
	SortBy string 
	Order Order
	Contracts []Contract 
}

type Contract struct {
	Name         string 
	Address      string 
	Transaction  string 
	Block        int `json:",string"`
	ContractType ContractType 
	Status       Status 
}

func main() {
	fmt.Println("hello world")

	var data JsonPayload
	file, err := ioutil.ReadFile("FlowJson.json")
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}

	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("failed to unmarshall file:", err)
		return
	}

	fmt.Println("Sort by:", data.SortBy)
	fmt.Println("Sort by:", data.Order)
	fmt.Println("Contracts:", data.Contracts)

	// sortByName(data.Contracts)
	
	// fmt.Println("Contracts:", data.Contracts)

	sortByAddress(data.Contracts)
	fmt.Println("Contracts:", data.Contracts)

}	

func sortByName (contracts []Contract) {
	sort.Slice(contracts[:], func(i, j int) bool {
		return contracts[i].Name < contracts[j].Name
	})
}


func sortByAddress (contracts []Contract) {
	sort.Slice(contracts[:], func(i, j int) bool {
		return contracts[i].Address < contracts[j].Address
	})
}
