package main

import "fmt"

// Largest integer allowed in our input
const size int = 1000000

func solution(A []int) int {
   sortable := [size]int{}
  for _, v := range A {
  // Since we do not care about negative numbers, we can get O(n) sort
  // by using the input values as indices of an array as big as the largest possible.
  // We also do not care about duplicates.
    if v > 0 {
       sortable[v] = 1
    }
  }
  // Now just go through all of them until we hit a zero
  // If so, return the next integer
  for i, v := range sortable[1:] {
      if v > 0 {
      	 continue
      } else {
      	 return i+1
      }
  }
  // If we have only negatives or nothing is passed in, 1 is the first we could get
  return 1
}

func main() {
    example1 := []int{1, 3, 6, 4, 1, 2}
    fmt.Println(example1)
    r1 := solution(example1)
    fmt.Println(r1)

    example2 := []int{1,2,3}
    fmt.Println(example2)
    r2 := solution(example2)
    fmt.Println(r2)

    example3 := []int{-1,-3}
    fmt.Println(example3)
    r3 := solution(example3)
    fmt.Println(r3)
}