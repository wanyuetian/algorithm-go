package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Customer struct {
	ID        int
	Buy       int
	EnterTime int
}

type WhosYourDaddy struct {
	Map        map[int]*Customer
	CandsHeap  *GreaterHeap[*Customer]
	DaddyHeap  *GreaterHeap[*Customer]
	DaddyLimit int
}

func NewWhosYourDaddy[T *Customer](limit int) *WhosYourDaddy {
	return &WhosYourDaddy{
		Map: make(map[int]*Customer),
		DaddyHeap: NewGreaterHeap(func(a, b *Customer) int {
			if a.Buy != b.Buy {
				return b.Buy - a.Buy
			}
			return a.EnterTime - b.EnterTime
		}, limit),
		DaddyLimit: limit,
	}
}

func (d *WhosYourDaddy) topK(arr []int, op []bool) [][]int {
	ans := make([][]int, 0)
	d.CandsHeap = NewGreaterHeap(func(a, b *Customer) int {
		if a.Buy != b.Buy {
			return a.Buy - b.Buy
		}
		return b.EnterTime - a.EnterTime
	}, len(arr))
	for i := 0; i < len(arr); i++ {
		d.operate(i, arr[i], op[i])
		ans = append(ans, d.getDaddies())
	}
	return ans
}

func (d *WhosYourDaddy) operate(i int, id int, isBuy bool) {
	if _, ok := d.Map[id]; !isBuy && !ok {
		return
	}
	if _, ok := d.Map[id]; !ok {
		d.Map[id] = &Customer{ID: id}
	}
	customer := d.Map[id]
	if isBuy {
		customer.Buy++
	} else {
		customer.Buy--
	}
	if customer.Buy == 0 {
		delete(d.Map, id)
	}
	if !d.CandsHeap.Contains(customer) && !d.DaddyHeap.Contains(customer) {
		customer.EnterTime = i
		if d.DaddyHeap.Size < d.DaddyLimit {
			d.DaddyHeap.Push(customer)
		} else {
			d.CandsHeap.Push(customer)
		}
	} else if d.CandsHeap.Contains(customer) {
		if customer.Buy == 0 {
			d.CandsHeap.Remove(customer)
		} else {
			d.CandsHeap.Resign(customer)
		}
	} else {
		if customer.Buy == 0 {
			d.DaddyHeap.Remove(customer)
		} else {
			d.DaddyHeap.Resign(customer)
		}
	}
	d.daddyMove(i)
}

func (d *WhosYourDaddy) daddyMove(i int) {
	if d.CandsHeap.IsEmpty() {
		return
	}
	if d.DaddyHeap.Size < d.DaddyLimit {
		c, _ := d.CandsHeap.Pop()
		c.EnterTime = i
		d.DaddyHeap.Push(c)
	} else {
		c1, _ := d.CandsHeap.Peek()
		c2, _ := d.DaddyHeap.Peek()
		if c1.Buy > c2.Buy {
			oldDaddy, _ := d.DaddyHeap.Pop()
			newDaddy, _ := d.CandsHeap.Pop()
			oldDaddy.EnterTime = i
			newDaddy.EnterTime = i
			d.DaddyHeap.Push(newDaddy)
			d.CandsHeap.Push(oldDaddy)
		}
	}
}

func (d *WhosYourDaddy) getDaddies() []int {
	ans := make([]int, 0)
	for i := 0; i < d.DaddyHeap.Size; i++ {
		ans = append(ans, d.DaddyHeap.Data[i].ID)
	}
	return ans
}

// 不优化
func compare(arr []int, op []bool, k int) [][]int {
	m := make(map[int]*Customer)
	cands := make([]*Customer, 0)
	daddy := make([]*Customer, 0)
	ans := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		id := arr[i]
		isBuy := op[i]
		_, ok := m[id]
		if !isBuy && !ok {
			ans = append(ans, getCurAns(daddy))
			continue
		}
		// 没有发生：用户购买数为0并且又退货了
		// 用户之前购买数是0，此时买货事件
		// 用户之前购买数>0， 此时买货
		// 用户之前购买数>0, 此时退货
		if !ok {
			m[id] = &Customer{
				ID:        id,
				Buy:       0,
				EnterTime: 0,
			}
		}
		// 买、卖
		customer := m[id]
		if isBuy {
			customer.Buy++
		} else {
			customer.Buy--
		}
		if customer.Buy == 0 {
			delete(m, id)
		}
		if !contains(cands, customer) && !contains(daddy, customer) {
			customer.EnterTime = i
			if len(daddy) < k {
				daddy = append(daddy, customer)
			} else {
				cands = append(cands, customer)
			}
		}
		daddy = cleanZeroBuy(daddy)
		cands = cleanZeroBuy(cands)
		sort.Slice(daddy, func(i, j int) bool {
			if daddy[i].Buy != daddy[j].Buy {
				return daddy[i].Buy < daddy[j].Buy
			}
			return daddy[i].EnterTime < daddy[j].EnterTime
		})
		sort.Slice(cands, func(i, j int) bool {
			if cands[i].Buy != cands[j].Buy {
				return cands[i].Buy > cands[j].Buy
			}
			return cands[i].EnterTime < cands[j].EnterTime
		})
		daddy, cands = move(cands, daddy, k, i)
		ans = append(ans, getCurAns(daddy))
	}

	return ans
}

func move(cands []*Customer, daddy []*Customer, k int, i int) ([]*Customer, []*Customer) {
	if len(cands) == 0 {
		return daddy, cands
	}
	if len(daddy) < k {
		customer := cands[0]
		customer.EnterTime = i
		daddy = append(daddy, customer)
		cands = cands[1:]
	} else {
		if cands[0].Buy > daddy[0].Buy {
			oldDaddy := daddy[0]
			oldDaddy.EnterTime = i
			daddy = daddy[1:]

			newDaddy := cands[0]
			newDaddy.EnterTime = i
			cands = cands[1:]

			daddy = append(daddy, newDaddy)
			cands = append(cands, oldDaddy)
		}
	}
	return daddy, cands
}

func contains[T *Customer](arr []T, elem T) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
}

func cleanZeroBuy(arr []*Customer) []*Customer {
	for i := 0; i < len(arr); i++ {
		if arr[i].Buy == 0 {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func getCurAns(daddy []*Customer) []int {
	ans := make([]int, 0)
	for i := 0; i < len(daddy); i++ {
		ans = append(ans, daddy[i].ID)
	}
	return ans
}

type Data struct {
	Arr []int
	Op  []bool
}

func randomData(maxValue int, maxLen int) *Data {
	length := rand.Intn(maxLen + 1)
	arr := make([]int, length)
	op := make([]bool, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue) + 1
		if rand.Float64() < 0.5 {
			op[i] = true
		}
	}
	return &Data{
		Arr: arr,
		Op:  op,
	}
}

func isSameAns(ans1 [][]int, ans2 [][]int) bool {
	if len(ans1) != len(ans2) {
		return false
	}
	for i := 0; i < len(ans1); i++ {
		cur1 := ans1[i]
		cur2 := ans2[i]
		if len(cur1) != len(cur2) {
			return false
		}
		sort.Ints(cur1)
		sort.Ints(cur2)
		for j := 0; j < len(cur1); j++ {
			if cur1[j] != cur2[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxValue := 10
	maxLen := 10
	maxK := 1
	testTimes := 100000

	for i := 0; i < testTimes; i++ {
		testData := randomData(maxValue, maxLen)
		//testData = &Data{
		//	Arr: []int{4, 9, 3, 7, 10, 10},
		//	Op:  []bool{true, true, false, true, true, true},
		//}
		k := rand.Intn(maxK) + 1
		whosYourDaddy := NewWhosYourDaddy(k)
		ans1 := whosYourDaddy.topK(testData.Arr, testData.Op)
		ans2 := compare(testData.Arr, testData.Op, k)
		if !isSameAns(ans1, ans2) {
			fmt.Println("Oops")
			return
		}
	}
	fmt.Println("Nice")
}
