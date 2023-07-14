package stuff

import "strconv"

var Name string = "ladoblanco"

func IntSliceToStrSlice(intSlice []int) []string {
	var strSlice []string
	for _, i := range intSlice {
		strSlice = append(strSlice, strconv.Itoa(i))
	}
	return strSlice
}
