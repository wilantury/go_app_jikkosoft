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

func sortArray(unsortedArray []int) {
	if len(unsortedArray) > 1 {
		var mid int = len(unsortedArray) / 2
		left := unsortedArray[:mid]
		right := unsortedArray[mid:]
		//divide and conquer: until length of array is less than 2
		sortArray(left)
		sortArray(right)

		//sort each array
		i, j, k := 0, 0, 0

		for i < len(left) && j < len(right) {

		}
	}
}

func (sa *SortArrayHandle) SortArrayView(w http.ResponseWriter, r *http.Request) {
	var sortRes SortArrayHandle
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sortRes)
	if err != nil {
		response(w, &sortRes, "error to decode body", http.StatusInternalServerError)
		return
	}
	fmt.Println(sortRes)
	response(w, &sortRes, "Success", http.StatusOK)
}
