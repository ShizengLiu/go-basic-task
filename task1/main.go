package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 只出现一次的数字
func OnlyOnceNumber() {
	nums := []int{4, 2, 1, 2, 5, 1, 4}
	map1 := make(map[int]int)
	for _, value := range nums {
		mV, flag := map1[value]
		if flag {
			map1[value] = mV + 1
		} else {
			map1[value] = 1
		}
		//进阶写法
		//map1[value]++
	}

	for key, value := range map1 {
		if value == 1 {
			fmt.Println("只出现一次的数字是:", key)
		}
	}
}

// 回文数
func PalindromicNumber() {
	test1 := 121
	test2 := -121
	test3 := 123
	test4 := 1221
	test5 := 4321

	t1Result := isPalindromicNumber(test1)
	t2Result := isPalindromicNumber(test2)
	t3Result := isPalindromicNumber(test3)
	t4Result := isPalindromicNumber(test4)
	t5Result := isPalindromicNumber(test5)

	fmt.Println("test1 is palindromic number:", t1Result)
	fmt.Println("test2 is palindromic number:", t2Result)
	fmt.Println("test3 is palindromic number:", t3Result)
	fmt.Println("test4 is palindromic number:", t4Result)
	fmt.Println("test5 is palindromic number:", t5Result)

}

func isPalindromicNumber(x int) bool {
	if x < 0 {
		return false
	}
	xStr := strconv.Itoa(x)
	length := len(xStr)
	fmt.Println("x:", xStr)

	for index := range xStr {
		if index == length/2 {
			break
		}
		if xStr[index] != xStr[length-1-index] {
			return false
		}
	}

	// left, right := 0, len(xStr)-1

	// for left < right {
	//     if xStr[left] != xStr[right] {
	//         return false
	//     }
	//     left++
	//     right--
	// }
	return true

}

// 优化方法，不使用字符串
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}

func ValidBrackets() {
	test1 := "()"
	test2 := "()[]{}"
	test3 := "(]"
	test4 := "([)]"
	test5 := "{[]}"
	t1s := isValidBrackets(test1)
	t2s := isValidBrackets(test2)
	t3s := isValidBrackets(test3)
	t4s := isValidBrackets(test4)
	t5s := isValidBrackets(test5)

	fmt.Println("test1 Reuslt:", t1s)
	fmt.Println("test2 Reuslt:", t2s)
	fmt.Println("test3 Reuslt:", t3s)
	fmt.Println("test4 Reuslt:", t4s)
	fmt.Println("test5 Reuslt:", t5s)

}

func isValidBrackets(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)
	for _, char := range s {
		value, exists := bracketMap[char]
		if !exists {
			stack = append(stack, char)

		} else {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0

}

//最长公共前缀

func testLongestCommonPrefix() {
	strs := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(strs)
	fmt.Println("最长公共前缀是:", result)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}

		}
	}
	return strs[0]

}

// 加1
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	newDigits := make([]int, len(digits)+1)
	fmt.Println("newDigits:", newDigits)
	newDigits[0] = 1
	return newDigits
}

func testPlusOne() {
	digits := []int{9, 9, 9}
	result := plusOne(digits)
	fmt.Println("加1结果是:", result)
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1

}

func testRemoveDuplicates() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	newLength := removeDuplicates(nums)
	fmt.Println("删除重复项后数组的新长度是:", newLength)
	fmt.Println("新的数组是:", nums[:newLength])
}

// 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		last := result[len(result)-1]

		if current[0] <= last[1] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			result = append(result, current)
		}
	}
	return result

}

func testMerge() {
	test1 := [][]int{{1, 4}, {0, 2}, {3, 5}}
	r1 := merge(test1)
	fmt.Println("merge r1:", r1)
}

// 两数之和
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

	// 优化方法，使用哈希表
	// numMap := make(map[int]int)

	// for i, num := range nums {
	// 	complement := target - num
	// 	index, exists := numMap[complement]
	// 	if exists {
	// 		return []int{index,i}
	// 	}
	// 	numMap[num] = i
	// }

}

func testTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("两数之和的结果是:", result)
}

func main() {
	OnlyOnceNumber()
	PalindromicNumber()
	ValidBrackets()
	testLongestCommonPrefix()
	testPlusOne()
	testRemoveDuplicates()
	testMerge()
	testTwoSum()
}
