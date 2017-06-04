package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Root struct {
	Took         int         `json:"took"`
	Timeout      bool        `json:"timed_out"`
	Shards       Shard       `json:"_shards"`
	Hits         Hit         `json:"hits"`
	Aggregations Aggregation `json:"aggregations"`
}

type Shard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type Hit struct {
	Total    int   `json:"total"`
	MaxScore int   `json:"max_score"`
	Hits     []Hit `json:"hits"`
}

type Aggregation struct {
	DR DayRaw `json:"day.raw"`
}

type DayRaw struct {
	Dceub   int         `json:"doc_count_error_upper_bound"`
	Sodc    int         `json:"sum_other_doc_count"`
	Buckets []DayBucket `json:"buckets"`
}

type DayBucket struct {
	Key     string     `json:"key"`
	Dc      int        `json:"doc_count"`
	ChanRaw ChannelRaw `json:"channel.raw"`
}

type ChannelRaw struct {
	Dceub   int          `json:"doc_count_error_upper_bound"`
	Sodc    int          `json:"sum_other_doc_count"`
	Buckets []ChanBucket `json:"buckets"`
}

type ChanBucket struct {
	Key     string     `json:"key"`
	Dc      int        `json:"doc_count"`
	ActtRaw ActtypeRaw `json:"acttype.raw"`
}

type ActtypeRaw struct {
	Dceub   int          `json:"doc_count_error_upper_bound"`
	Sodc    int          `json:"sum_other_doc_count"`
	Buckets []ActtBucket `json:"buckets"`
}

type ActtBucket struct {
	Key string `json:"key"`
	Dc  int    `json:"doc_count"`
	Ct  Ct     `json:"ct"`
}

type Ct struct {
	Value int `json:"value"`
}

func main() {
	f, err := os.Open("json.log")
	checkErr(err)

	data, err := ioutil.ReadAll(f)
	checkErr(err)

	var root Root
	err = json.Unmarshal(data, &root)
	checkErr(err)

	result, err := json.MarshalIndent(&root, "", "    ")
	checkErr(err)
	fmt.Printf("%s\n", result)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
