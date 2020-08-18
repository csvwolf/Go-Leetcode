package TwoSum

func twoSumBetter(nums []int, target int) []int {
    m := make(map[int]int)

    // 建立 map 顺便匹配结果
    for index, num := range nums {
        index2, found := m[target - num]
        m[num] = index
        if found && index != index2 {
            return []int{index2, index}
        }
    }
    return []int{}
}

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    // 建立 map
    for index, num := range nums {
        m[num] = index
    }

    // 匹配结果
    for index, num := range nums {
        index2, found := m[target - num]
        if found && index != index2 {
            return []int{index, index2}
        }
    }
    return []int{}
}
