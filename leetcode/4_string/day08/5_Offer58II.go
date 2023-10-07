package day08

// 方法一
func dynamicPassword(password string, target int) string {
	return password[target:] + password[:target]
}


// 方法二
func reverse(chars []byte, left, right int) {
    for left < right {
        temp := chars[left]
        chars[left] = chars[right]
        chars[right] = temp
        left++
        right--
    }
}
func dynamicPassword2(password string, target int) string {
    chars := []byte(password)
    reverse(chars, 0, target - 1)
    reverse(chars, target, len(chars) - 1)
    reverse(chars, 0, len(chars) - 1)
    return string(chars)
}