package utils

import (
	"fmt"
	"strconv"
)

func CastStrToFloat(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}

func CastFloatToStr(val float64) string {
	return fmt.Sprintf("%.2f", val)
}
