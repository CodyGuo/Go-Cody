package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type VersionList []string

func (v VersionList) Len() int {
	return len(v)
}
func (v VersionList) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func (v VersionList) Less(i, j int) bool {
	if v.containLetter(v[i]) && v.containLetter(v[j]) {
		return v[i] < v[j]
	}
	if !v.containLetter(v[i]) && !v.containLetter(v[j]) {
		return v[i] < v[j]
	}
	fields := strings.Split(v[i], ".")
	fields2 := strings.Split(v[j], ".")
	if fields[0] < fields2[0] {
		return true
	} else if fields[0] > fields2[0] {
		return false
	} else {
		if fields[1][0] < fields2[1][0] {
			return true
		} else if fields[1][0] == fields2[1][0] {
			if v.containLetter(v[i]) {
				return true
			} else {
				return false
			}
		} else if fields[1][0] > fields2[1][0] {
			return false
		}
	}
	return v[i] < v[j]
}
func (v VersionList) containLetter(version string) bool {
	if strings.Contains(version, "rc") || strings.Contains(version, "beta") || strings.Contains(version, "alpha") {
		return true
	}
	return false
}
func VersionSort(versions []string) []string {
	sort.Sort(VersionList(versions))
	return versions
}
func main() {
	timeNow := time.Now()
	versions := []string{"1.6beta1", "1.5rc1", "1.5beta2", "1.5beta1", "1.5.1", "1.5", "1.4rc2", "1.4rc1", "1.4beta1", "1.4.2", "1.4.1", "1.4", "1.3rc2", "1.3rc1", "1.3beta2", "1.3beta1", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2rc5", "1.2rc4", "1.2rc3", "1.2rc2", "1.2rc1", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.5.2", "1.5alpha1"}
	VersionSort(versions)
	fmt.Println(versions)

	fmt.Println("用时：", time.Since(timeNow))
}
