package main

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

// album represents data about a record album.
type album struct {
	Count  int
	Title  string
	Artist string
	Price  float64
}

func BuildResponse(sizeOfSlice int) []album {
	var t1 []album
	temp_count := 0
	var temp = album{
		Count:  temp_count,
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	}
	for i := 0; i < sizeOfSlice; i++ {
		temp.Count += 1
		t1 = append(t1, temp)
	}
	return t1
}

func MarshallingUnmarshallingJSON(t1 []album) {
	var result []*album
	byteReponse, _ := json.Marshal(t1)
	json.Unmarshal(byteReponse, &result)
}

func MarshallingUnmarshallingSONIC(t1 []album) {
	var result []album
	byteReponse, _ := sonic.Marshal(t1)
	sonic.Unmarshal(byteReponse, &result)
}

func main() {
	fmt.Println("Why Sonic is better than other Sherde options")
}
