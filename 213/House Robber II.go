package main

/*
思路一: 需要选择的点成环，如果正常顺序dp的话，前面选取的策略最后会影响第一个点的选取，也就是有后效性。
所以，需要一定选择的点，然后这样和它相邻的就一定是无法选的，排除掉这三个点跑一个dp就行了
时间复杂度 O(N^2)
*/

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func rob(nums []int) int {
	mx := 0
	for i := 0; i < len(nums); i++ {
		f1 := 0
		f0 := 0
		st := nums[i]
		for j := i + 2; j < i+len(nums)-1; j++ {
			nf1 := f0 + nums[j%len(nums)]
			nf0 := max(f0, f1)
			f0 = nf0
			f1 = nf1
		}
		mx = max(max(f0, f1)+st, mx)
	}
	return mx
}

/*

	思路2：对于相邻的三个，肯定有一个必选，所以上面n的范围变成3就行了
*/
