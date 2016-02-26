package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Version []string

func (v Version) VersionSort() []string {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v)-1-i; j++ {
			numS1 := v.verSplist(v[j])
			numS2 := v.verSplist(v[j+1])
			num1, _ := strconv.Atoi(numS1[1])
			num2, _ := strconv.Atoi(numS2[1])

			if num1 > num2 {
				fmt.Println("正在排序2:", v[j], v[j+1])
				v.swap(j, j+1)
			}
		}

	}

	return nil
}

func (v Version) verSplist(value string) []string {
	return strings.Split(value, ".")
}

func (v Version) strToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
func (v Version) swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Version) verLen(ver []string) {
}
func (v Version) tagPrint(tag []string) {
}

func main() {
	timeGo := time.Now()
	var ver Version
	ver = []string{"1.8", "1.10", "1.2", "1.5", "1.3", "1.1", "1.0", "1.7"}
	// ver = []string{
	// 	"1.6beta1",
	// 	"1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.5.2", "1.5alpha1",
	// 	"1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4",
	// 	"1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3",
	// 	"1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2",
	// 	"1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2",
	// }

	ver.tagPrint([]string{"beta", "alpha", "rc"})
	fmt.Println("排序前:")
	fmt.Println(ver)

	fmt.Println("排序后:")
	ver.VersionSort()
	fmt.Println(ver)

	fmt.Println("用时: ", time.Since(timeGo))
}
