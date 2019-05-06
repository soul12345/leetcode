package 38
func getAns(res string) string {
	res1 := ""
	cnt := 0
	for i < len(res) {
		if i == 0 || res[i] == res[i-1] {
			cnt++
		} else {
			res1 += string('0' + cnt)
			res1 += res[i-1]
			cnt = 1
		}
	}
	return res1
}
func countAndSay(n int) string {
	res := "1"
	for i := 2; i <= n; i++ {
		res = getAns(res)
	}
	return res
}