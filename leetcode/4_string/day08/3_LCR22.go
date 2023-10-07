package day08

func pathEncryption(path string) string {
	newString := ""
	for _, v:= range path{
		if v == '.' {
			newString += " "
		} else {
			newString += string(v)
		}
	}
    return newString
}