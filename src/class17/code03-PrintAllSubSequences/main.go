package main

import "fmt"

func subs(s string) []string {
	path := ""
	ans := make([]string, 0)
	process1(s, 0, &ans, path)
	return ans
}

func process1(s string, index int, ans *[]string, path string) {
	if index == len(s) {
		*ans = append(*ans, path)
		return
	}
	// 不要当前index
	process1(s, index+1, ans, path)
	// 要当前index
	process1(s, index+1, ans, path+string(s[index]))
}

func subsNoRepeat(s string) []string {
	path := ""
	ans := make([]string, 0)
	set := make(map[string]struct{}, 0)
	process2(s, 0, set, path)
	for k := range set {
		ans = append(ans, k)
	}
	return ans
}

func process2(s string, index int, set map[string]struct{}, path string) {
	if index == len(s) {
		set[path] = struct{}{}
		return
	}
	// 不要当前index
	process2(s, index+1, set, path)
	// 要当前index
	process2(s, index+1, set, path+string(s[index]))
}

func main() {
	test := "acccc"
	ans1 := subs(test)
	for i := range ans1 {
		fmt.Println(ans1[i])
	}
	fmt.Println("====================")
	ans2 := subsNoRepeat(test)
	for k := range ans2 {
		fmt.Println(ans2[k])
	}
}
