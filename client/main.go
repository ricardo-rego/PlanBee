package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorules/zen-go"
)

func main() {
	// Initialize the Zen Engine with default configuration
	engine := zen.NewEngine(zen.EngineConfig{})

	// Read the JSON decision graph file
	graph, err := os.ReadFile("./v6.json")
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	// Create a decision from the graph
	decision, err := engine.CreateDecision(graph)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	// Prepare the data for decision evaluation
	data := map[string]interface{}{
		"documento":     "35077857000194",
		"nsu":           1,
		"tipoPessoa":    "PJ",
		"tipoPagamento": "A",
		"trxPos":        1500.00,
		"trxAcum":       10000,
		"trxMes":        5000000.00,
		"trxHorario":    "1330",
		"renda":         90000,
		"faturamento":   100000,
		"cnae":          3313901,
		"porte":         "ME",
		"mccRisco":      "F",
		"csnu":          "F",
		"fronteira":     "F",
	}

	// Evaluate the decision with the prepared data
	response, err := decision.Evaluate(data)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	// Convert the response to a pretty-printed JSON string
	responseJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	// Print the response
	fmt.Println(string(responseJSON))
}
