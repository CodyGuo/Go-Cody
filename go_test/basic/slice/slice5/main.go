package main

import "sort"
import "strings"
import (
	"fmt"
	"time"
)

var unsorted = []string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}

func main() {
	timeNow := time.Now()
	sorted := VersionSort(unsorted)
	for _, v := range sorted {
		println(v)
	}

	fmt.Println("用时：", time.Since(timeNow))
}
func VersionSort(versions []string) []string {
	s2 := make([]string, len(versions))
	for i, v := range versions {
		s2[i] = strings.Replace(v, ".", "~", -1) + "~"
	}
	sort.Strings(s2)
	for i, v := range s2 {
		s2[i] = strings.Replace(v[:len(v)-1], "~", ".", -1)
	}
	return s2
}
