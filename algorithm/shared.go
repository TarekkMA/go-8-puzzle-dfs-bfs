package algorithm

func CalcIntArrHashCode(arr []int) int {
	result := 1
	for i := 0; i < len(arr); i++ {
		result = 31*result + arr[i]
	}
	return result
}
