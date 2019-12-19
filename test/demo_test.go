package test

import (
	"sort"
	"testing"
)

// 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// 先排序再去重
func RemoveRepeatedElement(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 利用了map的多返回值特性
func removeDuplicateElement(addrs []int) []int {
	result := make([]int, 0, len(addrs))
	temp := map[int]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// 去重
func removeDuplicate(arr []int) []int {
	resArr := make([]int, 0)
	tmpMap := make(map[int]interface{})
	for _, val := range arr {
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = struct{}{}
		}
	}
	return resArr
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop2(arr []int) []int {
	result := make([]int, 0) // 存放结果
	for i := 0; i < len(arr); i++ {
		exist := false
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop3(arr []int) []int {
	result := make([]int, 0) // 存放结果
	for i := 0; i < len(arr); i++ {
		exist := true
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = false
				break
			}
		}
		if exist {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop4(arr []int) []int {
	result := make([]int, 0, len(arr)) // 存放结果
	for i := 0; i < len(arr); i++ {
		exist := true
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = false
				break
			}
		}
		if exist {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop5(arr []int) []int {
	result := make([]int, 0) // 存放结果
	for i := 0; i < len(arr); i++ {
		j := 0
		for ; j < len(result); j++ {
			if result[j] == arr[i] {
				break
			}
		}
		if j == len(result) {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop6(arr []int) []int {
	result := make([]int, 0) // 存放结果
	for i := range arr {
		flag := true
		for j := range result {
			if arr[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop7(arr []int) []int {
	// 存放结果，空初始化比make定义要快
	result := []int{}
	// 外层循环准备添加到结果的切片
	for i := 0; i < len(arr); i++ {
		// 初始定义该元素不存在，很奇怪，初始在循环里面比先声明在外面之后再赋值要快
		exist := false
		// 这里根据当前切片的长度进行循环，直接使用 len 比初始一个 count变量 记数要快
		for j := 0; j < len(result); j++ {
			// 如果遇到重复提前退出
			if result[j] == arr[i] {
				// 并且说明已存在
				exist = true
				break
			}
		}
		// 如果在 result切片都没有遍历到此元素
		if !exist {
			// 那么就追加到 result
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop8(arr []int) []int {
	result := []int{} // 存放结果
	for i := 0; i < len(arr); i++ {
		j := 0
		for ; j < len(result); j++ {
			if result[j] == arr[i] {
				break
			}
		}
		if j == len(result) {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop9(arr []int) []int {
	result := []int{} // 存放结果
	count := 0
	for i := 0; i < len(arr); i++ {
		exist := false
		for j := 0; j < count; j++ {
			if result[j] == arr[i] {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, arr[i])
			count++
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop10(arr []int) []int {
	result := []int{} // 存放结果
	var exist bool
	for i := 0; i < len(arr); i++ {
		exist = false
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop11(arr []int) []int {
	result := []int{} // 存放结果
	for i := 0; i < len(arr); i++ {
		exist := true
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = false
				break
			}
		}
		if exist {
			result = append(result, arr[i])
		}
	}
	return result
}

// 通过两重循环过滤重复元素
func RemoveRepByLoop12(arr []int) []int {
	result := []int{} // 存放结果
	for i := range arr {
		exist := false
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, arr[i])
		}
	}
	return result
}

func Benchmark_Demo(b *testing.B) {
	// 0.226 ns/op           0 B/op
	arr := []int{1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6, 7, 4, 1, 1, 2, 3, 5, 3, 1, 5, 6}

	for i := 0; i < b.N; i++ { //use b.N for looping
		//RemoveRepByLoop(arr) // 10625 ns/op             120 B/op
		//RemoveRepByMap(arr) // 85435 ns/op             120 B/op
		//RemoveRepeatedElement(arr)  // 115283 ns/op             158 B/op
		//removeDuplicateElement(arr)  // 46070 ns/op           65536 B/op
		//unique(arr)  // 36345 ns/op             120 B/op
		//removeDuplicate(arr)  // 36633 ns/op             120 B/op

		//RemoveRepByLoop2(arr) // 14299 ns/op             120 B/op
		//RemoveRepByLoop3(arr)  // 14226 ns/op             120 B/op
		//RemoveRepByLoop4(arr)  // 16065 ns/op           65536 B/op
		//RemoveRepByLoop5(arr)  // 19612 ns/op             120 B/op
		//RemoveRepByLoop6(arr)  //  14283 ns/op             120 B/op
		//RemoveRepByLoop7(arr) // 10120 ns/op             120 B/op   最强
		//RemoveRepByLoop8(arr) // 18173 ns/op             120 B/op
		//RemoveRepByLoop9(arr)  // 14655 ns/op             120 B/op
		RemoveRepByLoop10(arr)  // 10146 ns/op             120 B/op
		//RemoveRepByLoop11(arr)  //  10130 ns/op             120 B/op
		//RemoveRepByLoop12(arr)  // 10146 ns/op             120 B/op

	}
	//fmt.Println(RemoveRepByLoop5(arr))
}
