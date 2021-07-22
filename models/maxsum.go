package models

import "math"

type ListRequest struct {
	List []float64 `json:"list"`
}

func Kadanes(list []float64) int {
	//simple Kadane's Algorithms implementation for O(n) solution

	//Declaring globalMax to keep track of general sum, and currentMax to keep track of specific index sum
	var currentMax, globalMax = list[0], list[0]

	//iterating through our list
	for index, element := range list {
		//Because the only possible subArray ending at index 0 is the element itself, we can skip it
		if index > 0 {

			//compare which one is higher, current element alone, or current element + previous highest sum
			currentMax = math.Max(element, currentMax+element)

			//if ou new subArray sum is higher than the previous, update it.
			if currentMax > globalMax {
				globalMax = currentMax
			}
		}
	}

	return int(globalMax)
}

//line 20-> Kadane provou por contradição que a maior soma de um subarray terminado em um indice n, é
//ou o elemento em si, ou a soma do elemento com o subarray de maior soma terminado no indice n-1.
