package main

import(
	"fmt"
)

func main(){
	fmt.Println("Maps Problem Sets")
	nums1 := []int{1,2,3,4,5,6,7,7}
	nums2 := []int{7,7,7,7}
	fmt.Print(intersect(nums1,nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
    // make map
    maps := make(map[int]int)
    for _, n := range nums1 {
        _, key_exists := maps[n] 
        if !key_exists {
            maps[n] = maps[n] + 1
        }
    }
    // make empty array
    var ans []int
    // check second array with map
    for _, n := range nums2{
        if maps[n] == 1 {
            maps[n] = maps[n] - 1
            ans = append(ans, n) 
        }
    }
    // return new array
    return ans
}