package dailycodingproblem

func sumNonAdjacent(arr []int) int {
	maxi := 0 //max value including previous item in array, doubles as running max value
	maxe := 0 //max value excluding previous item in array
	for _, i := range arr {
		// add item to maxe if it will increase it's value
		if i > 0 {
			maxe += i
		}

		//grab mai for latter, jic
		tmp := maxi
		//only update maxi if maxe is larger
		if maxi < maxe {
			maxi = maxe
		}

		//set maxe to old maxi
		maxe = tmp

	}
	return maxi
}
