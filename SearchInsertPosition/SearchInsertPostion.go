package SerchInsertPosition

// https://leetcode-cn.com/problems/search-insert-position/
// 普通方法 O(n)
func searchInsert(nums []int, target int) int {
    for index, num := range nums {
        if target <= num {
            return index
        }
    }
    return len(nums)
}

// 二分法 O(logN)
func searchInsert2(nums []int, target int) int {
    left, right, mid := 0, len(nums) - 1, 0
    result := len(nums) // 如果找不到应该是插入最后

    for left <= right {
        // 普通写法，容易溢出
        // mid = (left + right) / 2
        // 一般写法
        // mid = left + (right - left) / 2
        // 更快的写法
        mid = (right - left) >> 1 + left
        // 只有当小的时候才代表 result 的插入值应该在当前位置
        if target <= nums[mid] {
            right = mid - 1
            result = mid
        } else {
            left = mid + 1
        }
    }
    return result
}
