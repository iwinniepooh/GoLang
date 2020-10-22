package main

import (
	"encoding/json"
	"fmt"
)

func pprint(d interface{}) {
	empJSON, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(empJSON))
}

func main() {
	sample := map[string]interface{}{
		"Key1": "Value1",
		"Key2": 2,
		"Key3": 3.3,
	}

	pprint(sample)
}
