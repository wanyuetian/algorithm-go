package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Employee struct {
	Happy int
	Nexts []*Employee
}

func maxHappy1(x *Employee) int {
	if x == nil {
		return 0
	}
	return process1(x, false)
}

func process1(x *Employee, up bool) int {
	if up {
		ans := 0
		for _, next := range x.Nexts {
			ans += process1(next, false)
		}
		return ans
	} else {
		p1 := x.Happy
		p2 := 0
		for _, next := range x.Nexts {
			p1 += process1(next, true)
			p2 += process1(next, false)
		}
		return int(math.Max(float64(p1), float64(p2)))
	}
}

type Info struct {
	Yes int
	No  int
}

func maxHappy2(x *Employee) int {
	allInfo := process2(x)
	return int(math.Max(float64(allInfo.Yes), float64(allInfo.No)))
}

func process2(x *Employee) *Info {
	if x == nil {
		return &Info{
			Yes: 0,
			No:  0,
		}
	}
	yes := x.Happy
	no := 0
	for _, next := range x.Nexts {
		allInfo := process2(next)
		no += int(math.Max(float64(allInfo.Yes), float64(allInfo.No)))
		yes += allInfo.No
	}
	return &Info{
		Yes: yes,
		No:  no,
	}
}

func generateBoss(maxLevel int, maxNexts int, maxHappy int) *Employee {
	if rand.Intn(100) <= 2 {
		return nil
	}
	boss := &Employee{
		Happy: rand.Intn(maxHappy + 1),
		Nexts: make([]*Employee, 0),
	}
	generateNexts(boss, 1, maxLevel, maxNexts, maxHappy)
	return boss
}

func generateNexts(e *Employee, level int, maxLevel int, maxNexts int, maxHappy int) {
	if level > maxLevel {
		return
	}
	nextSize := rand.Intn(maxNexts + 1)
	for i := 0; i < nextSize; i++ {
		next := &Employee{
			Happy: rand.Intn(maxHappy + 1),
			Nexts: make([]*Employee, 0),
		}
		e.Nexts = append(e.Nexts, next)
		generateNexts(next, level+1, maxLevel, maxNexts, maxHappy)
	}
}

func main() {
	maxLevel := 4
	maxNexts := 7
	maxHappy := 100
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		boss := generateBoss(maxLevel, maxNexts, maxHappy)
		if maxHappy1(boss) != maxHappy2(boss) {
			fmt.Println("Oops!")
			return
		}
	}
	fmt.Println("finish!")
}
