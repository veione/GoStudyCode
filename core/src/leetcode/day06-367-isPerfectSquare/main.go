package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left, right := 2, num/2
	for left <= right{
		mid := left +(right-left)/2
		if mid * mid == num {
			return true
		}else if mid * mid < num{
			left = mid + 1
		}else {
			right = mid - 1
		}
	}
	return false
}

