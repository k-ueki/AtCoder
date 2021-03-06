package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var p int
	var s string
	ac := make(map[int]struct{})
	wa := make(map[int]int)
	for i := 0; i < m; i++ {
		fmt.Scan(&p, &s)
		if s == "AC" {
			ac[p] = struct{}{}
		} else if _, o := ac[p]; !o {
			wa[p]++
		}
	}

	pen := 0
	for i, v := range wa {
		if _, o := ac[i]; o {
			pen += v
		}
	}
	fmt.Println(len(ac), pen)
}

/*


//===========================
//
// Library
//
//===========================

//===========================
//  round(12.54, 0) -> 13.0
//  round(12.54, -1) -> 12.5
//===========================

func round(num float64, digit uint) float64 {
    d := float64(digit)
    return float64(int(num * math.Pow(10, d) + 0.5)) / math.Pow(10, d)
}

func max(a, b int) int {
    return int(math.Max(float64(a), float64(b)));
}

func min(a, b int) int {
    return int(math.Min(float64(a), float64(b)));
}

func gcd(x, y int) int {
    if y == 0 {
        return x
    }

    return gcd(y, x%y)
}

//===========================
//  var q Queue
//  q.Enqueue(n)
//  n := q.Dequeue()
//===========================
type Queue struct {
    data []int
}

func (q *Queue) Enqueue(n int) {
    q.data = append(q.data, n)
}

func (q *Queue) Dequeue() int {
    tmp := q.data[0]
    q.data = q.data[1:]
    return tmp
}

// 数値を桁区切りのスライスにする
func digit(n int, list []int) []int {
   if n > 0 {
       return digit(n/10, append(list, n%10))
   }
   return list
}

func sortIntsRev(list []int) {
    sort.Sort(sort.Reverse(sort.IntSlice(list)))
}

func strReverse(s string) string {
    runes := []rune(s)
    for i := 0; i < len(s)/2; i++ {
        runes[i], runes[len(s)-1-i] = runes[len(s)-1-i], runes[i]
    }
    return string(runes)
}

func nextInt() int {
    sc.Scan()
    i, _ := strconv.Atoi(sc.Text())
    return i
}

func nextStr() string {
    sc.Scan()
    return sc.Text()
}

*/
