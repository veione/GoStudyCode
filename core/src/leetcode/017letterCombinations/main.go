package main

//func letterCombinations(digits string) []string {
//	numToStr := map[string]string{}
//	numToStr["2"] = "abc"
//	numToStr["3"] = "def"
//	numToStr["4"] = "ghi"
//	numToStr["5"] = "jkl"
//	numToStr["6"] = "mno"
//	numToStr["7"] = "pqrs"
//	numToStr["8"] = "tuv"
//	numToStr["9"] = "wxyz"
//	result := make([]string, 0, 10)
//	var dfs func(res string, digits string)
//	dfs = func(res string, digits string) {
//		if len(digits) == 0 {
//			result = append(result, res)
//			return
//		}
//		for _, char := range numToStr[string(digits[0])] {
//			res += string(char)
//			dfs(res, digits[1:])
//			res = res[:len(res)-1]
//		}
//	}
//	dfs("", digits)
//	return result
//}
//
//func main() {
//	var str string
//	for {
//		_, err := fmt.Scan(&str)
//		if err == io.EOF {
//			break
//		}
//		rt := letterCombinations(str)
//		for _, v := range rt {
//			fmt.Printf("%s ", v)
//		}
//	}
//}

func letterCombinations(digits string) []string {
	ans := []string{}
	dMap := map[byte]string{}
	dMap['2'] = "abc"
	dMap['3'] = "def"
	dMap['4'] = "ghi"
	dMap['5'] = "jkl"
	dMap['6'] = "mno"
	dMap['7'] = "pqrs"
	dMap['8'] = "tuv"
	dMap['9'] = "wxyz"
	printLetters(digits, dMap, 0, 0, []byte{}, &ans)
	return ans
}

func printLetters(digits string, dMap map[byte]string, dcur int, dmcur int, tmp []byte, ans *[]string) {
	ld := len(digits)
	if (dcur >= ld) || (dmcur >= len(dMap[digits[dcur]])) {
		return
	}

	if dmcur != 0 {
		tmp[dcur] = dMap[digits[dcur]][dmcur]
	} else {
		tmp = append(tmp, dMap[digits[dcur]][dmcur])
	}

	if len(tmp) == len(digits) {
		*ans = append(*ans, string(tmp))
	}
	printLetters(digits, dMap, dcur+1, 0, tmp, ans)
	printLetters(digits, dMap, dcur, dmcur+1, tmp, ans)
}
