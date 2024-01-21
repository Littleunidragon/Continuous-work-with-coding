// Daily challenge - FAILED
//The best solution i've found - by xopoww

package main

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func rob(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    l1 := nums[0]
    l2 := max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        l2, l1 = max(l1 + nums[i], l2), l2
    }
    return l2
}

func main(){
  rob([]int{3,1,2,5,7,34,12,54,0}
      }
