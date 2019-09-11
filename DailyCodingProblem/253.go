package dailycodingproblem
import (
	"fmt"
	"strings"
)

func problem253(InputString string, k int) string{
	lineCount := len(InputString)

	if lineCount > k {
		lineCount = k 
	}

	lines := make([]string, lineCount)

	for index := 0; index < len(InputString); index++ {
		placement := index%(lineCount*2)
		if placement > lineCount {
			placement = lineCount-placement+lineCount
		}
		fmt.Printf("Af: Placement: %v, Linecount: %v\n",placement, lineCount)
		lines[placement] += string(InputString[index])
	}
	return strings.Join(lines,"\n")
}
