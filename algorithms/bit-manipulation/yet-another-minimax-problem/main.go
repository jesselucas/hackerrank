package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

type Uint64Slice []uint64

func (a Uint64Slice) Len() int           { return len(a) }
func (a Uint64Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Uint64Slice) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	var T int
	_, err := fmt.Scan(&T)
	if err != nil {
		log.Fatal(err)
	}

	if T <= 1 {
		log.Fatal("Input must be greater than 1")
	}

	var nums []uint64
	for i := 0; i <= T-1; i++ {
		var n int
		scan(&n)
		nums = append(nums, uint64(n))
	}

	fmt.Println(minMaxXOR(nums))
}

func minMaxXOR(nums Uint64Slice) uint64 {
	// First check to see if they are all the same number
	sort.Sort(nums)
	if nums[0] == nums[len(nums)-1] {
		return 0
	}

	msbs := []uint64{}
	msbPlace := 0
	var mask uint64
	for _, n := range nums {
		for i := 63; i >= 0; i-- {
			mask = 1 << uint64(i)
			if n&mask > 0 {
				if i > msbPlace {
					msbPlace = i
				}

				break
			}
		}
	}

	mask = 1 << uint64(msbPlace)
	for _, n := range nums {
		if n&uint64(mask) > 0 {
			add := true
			// only add if the numbers isn't in the msbs
			for _, m := range msbs {
				if m == n {
					add = false
				}
			}
			if add {
				msbs = append(msbs, n)
			}
		}
	}

	// Remove msbInts from nums
	for _, m := range msbs {
		for i := 0; i < len(nums); i++ {
			if m == nums[i] {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		}
	}

	// XOR against all numbers
	var minXOR uint64 = math.MaxUint64
	for _, m := range msbs {
		for _, n := range nums {
			xor := m ^ n
			if xor < minXOR {
				minXOR = xor
			}
		}
	}

	// If they are all candidates for MSBs test them against themself
	if len(nums) == 0 {
		for i, m := range msbs {
			// Reduce
			msbs[i] = m ^ uint64(mask)
			// fmt.Println("m", m, "msb", msbs[i])
		}

		// Retest
		// TODO
		return minMaxXOR(msbs)
	}

	// fmt.Println(msbs)
	// fmt.Println(nums)

	return minXOR
}

func scan(n *int) int {
	fmt.Scan(n)
	return *n
}
