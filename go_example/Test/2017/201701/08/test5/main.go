package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

var (
	Web1   = fakeSearch("web1")
	Web2   = fakeSearch("web2")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
)

func Google(query string) (results []Result) {
	c := make(chan Result)
	/**
	 * Search 4
	 */
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	// go func() {
	// 	c <- Web(query)
	// }()
	// go func() {
	// 	c <- Image(query)
	// }()
	// go func() {
	// 	c <- Video(query)
	// }()

	/**
	 * Sarch 3
	 */
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("time out.")
			return
		}
	}

	/**
	 * Search 2
	 */
	// for i := 0; i < 3; i++ {
	// 	result := <-c
	// 	results = append(results, result)
	// }

	/**
	 * Search 1
	 */
	// results = append(results, Web(query))
	// results = append(results, Image(query))
	// results = append(results, Video(query))
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	// results := First("golang", fakeSearch("Web"),
	// 	fakeSearch("Image"), fakeSearch("Video"))
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
