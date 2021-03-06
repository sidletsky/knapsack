package bruteforce

import (
	"fmt"
	"math"

	"knapsack/repository"
)

type Knapsack struct {
	TotalValue           int
	TotalWeight          int
	CharacteristicVector []byte
}

func (knapsack *Knapsack) String() string {
	return fmt.Sprintf(
		"Characteristic vector: %d \n"+
			"Total value: %d \n"+
			"Total weight: %d",
		knapsack.CharacteristicVector,
		knapsack.TotalValue,
		knapsack.TotalWeight)
}

// GetKnapsack finds a perfect knapsack from all possible solutions using brute force method.
// Perfect knapsack is the one that has the biggest total value of items and has weight
// lower or equal to the capacity of a required knapsack.
// GetKnapsack finds perfect knapsack on the fly, without producing any helper matrices
func GetKnapsack(items []repository.Item, knapsackCapacity int) (perfectKnapsack Knapsack) {
	itemsSize := len(items)
	totalKnapsacks := int(math.Pow(2, float64(itemsSize)))
	perfectKnapsack = Knapsack{}
BinaryIterator:
	for i := 0; i < totalKnapsacks; i++ {
		// create characteristicVector representation
		characteristicVector := DecToBin(i, itemsSize)
		// build knapsack with items
		currentKnapsack := Knapsack{}
		for i, v := range characteristicVector {
			// 49 represents 1, 48 represents 0
			if v == 49 {
				currentKnapsack.TotalValue += items[i].Value
				currentKnapsack.TotalWeight += items[i].Weight
			}
			if currentKnapsack.TotalWeight > knapsackCapacity {
				continue BinaryIterator
			}
		}
		// choosing best knapsack
		if perfectKnapsack.TotalValue < currentKnapsack.TotalValue {
			perfectKnapsack = currentKnapsack
			perfectKnapsack.CharacteristicVector = characteristicVector
		}
	}
	return
}
