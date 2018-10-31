package dailycodingproblem

func maxSubArray(arr []int, k int) []int {
	var ret []int
	var max *int
	for i := 0; i+k <= len(arr); i++ {
		if max == nil { //Did we lose the last max to the window shift?
			subArr := arr[i : i+k]
			for i2 := range subArr {
				if max == nil || subArr[i2] > *max {
					max = &subArr[i2]
				}
			}
		} else if *max < arr[i+k-1] { //Nope? compare to the new blood
			max = &arr[i+k-1]
		}
		ret = append(ret, *max)

		// assume current max is on the way out
		if arr[i] == *max {
			max = nil
		}
	}
	return ret
}
