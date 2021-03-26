package sort_array

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//sort array struct
type SortArrayHandle struct {
	Unsorted []int `json:"unsorted"`
	Sorted   []int `json:"sorted"`
}

func (s SortArrayHandle) String() string {
	return fmt.Sprintf("unsorted: %v , sorted: %v", s.Unsorted, s.Sorted)
}

func response(w http.ResponseWriter, data *SortArrayHandle, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	resp := make(map[string][]int)
	resp["unsorted"] = data.Unsorted
	resp["sorted"] = data.Sorted
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func Merge(left, right []int) []int {
	sortArray := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(sortArray, right...)
		}
		if len(right) == 0 {
			return append(sortArray, left...)
		}
		if left[0] <= right[0] {
			sortArray = append(sortArray, left[0])
			left = left[1:]
		} else {
			sortArray = append(sortArray, right[0])
			right = right[1:]
		}
	}
	return sortArray
}

func MergeSort(unsortedArray []int) []int {
	if len(unsortedArray) <= 1 {
		return unsortedArray
	}
	mid := len(unsortedArray) / 2
	left := MergeSort(unsortedArray[:mid])
	right := MergeSort(unsortedArray[mid:])
	return Merge(left, right)
}

func separateRepeated(sortedArray []int) []int {
	var repeat []int
	index := 0
	for i := 0; i < len(sortedArray)-1; i++ {
		if sortedArray[i] == sortedArray[i+1] {
			repeat = append(repeat, sortedArray[i])
		} else {
			sortedArray[index] = sortedArray[i]
			if i == len(sortedArray)-2 {
				index++
				sortedArray[index] = sortedArray[i+1]
			}
			index++
		}
	}

	for _, val := range repeat {
		sortedArray[index] = val
		index++
	}
	return sortedArray
}

func (sa *SortArrayHandle) SortArrayView(w http.ResponseWriter, r *http.Request) {
	var sortRes SortArrayHandle
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sortRes)
	//sort array
	sol := MergeSort(sortRes.Unsorted)
	sortRes.Sorted = separateRepeated(sol)

	if err != nil {
		response(w, &sortRes, "error to decode body", http.StatusInternalServerError)
		return
	}

	response(w, &sortRes, "Success", http.StatusOK)
}
