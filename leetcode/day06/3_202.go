package day05

func isHappy(n int) bool {
	set := map[int]bool{}  // 模拟集合
	for {
		sum := 0
        //拆数求平方和
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
        //得到正确答案
		if sum == 1 {
			return true
		}
        
        //结果不在集合中则放入集合
		if _, ok := set[sum]; !ok {
			set[sum] = true
		} else {// 否则退出循环 return false
			break
		}
		n = sum // 新一轮的数
	}
	return false
}