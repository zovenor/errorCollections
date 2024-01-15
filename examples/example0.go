package main

import (
	"fmt"

	"github.com/zovenor/errorCollections/errorCollections"
)

func main() {
	errorCollections := errorCollections.NewErrorCollections()

	err := fmt.Errorf("error type 1")
	errorCollections.AddError(err)

	for i := 0; i < 20; i++ {
		errorCollections.AddError(fmt.Errorf("error type 2"))
	}

	for _, errorCollection := range errorCollections.ErrorCollections {
		fmt.Printf("error collection: %v, count: %v\n", errorCollection.Error(), errorCollection.Count())
	}
}
