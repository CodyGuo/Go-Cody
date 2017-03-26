package main

import "fmt"

func location(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		continent = "North America"
	default:
		continent = "Unknown"
	}
	return name, continent
}
func main() {
	region, continent := location("Matt", "LA")
	fmt.Printf("%s lives in %s", region, continent)
}
