package helpers

import (
	"fin/common"
	"strconv"
	"strings"
)

func ToInt(s string) int {
	r, e := strconv.Atoi(s)
	common.Error{}.CheckErr(e)
	return r
}

func ToBool(s string) bool {
	r, e := strconv.ParseBool(s)
	common.Error{}.CheckErr(e)
	return r
}

func ToValue(s string) float32 {
	i := strings.Replace(s, "R$", "", -1)
	var y string
	if strings.Count(s, ".") > 0 {
		x := strings.Replace(i, ".", "", -1)
		y = strings.Replace(x, ",", ".", 1)
	} else {
		y = strings.Replace(i, ",", ".", 1)
	}
	j := strings.Replace(y, " ", "", -1)
	v, err := strconv.ParseFloat(j, 32)
	common.Error{}.CheckErr(err)
	return float32(v)
}
