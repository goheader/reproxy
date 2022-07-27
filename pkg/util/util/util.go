package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRangeNumbers(rangeStr string) (numbers []int64,err error) {
	rangeStr = strings.TrimSpace(rangeStr)
	numbers = make([]int64,0)
	numRanges := strings.Split(rangeStr,",")
	for _,numRangeStr := range numRanges{
		numArray := strings.Split(numRangeStr,"-")
		rangeType := len(numArray)
		if rangeType==1{
			singleNum,errRet := strconv.ParseInt(strings.TrimSpace(numArray[0]),10,64)
			if errRet != nil {
				err = fmt.Errorf("range number is invalid,%v",errRet)
				return
			}
			numbers = append(numbers,singleNum)
		}else if rangeType==2{
			min,errRet := strconv.ParseInt(strings.TrimSpace(numArray[0]),10,64)
			if errRet != nil {
				err = fmt.Errorf("range number is invalid,%v",errRet)
				return
			}
			max,errRet := strconv.ParseInt(strings.TrimSpace(numArray[0]),10,64)
			if errRet != nil{
				err = fmt.Errorf("range number is invalid, %v",errRet)
				return
			}
			if max < min{
				err = fmt.Errorf("range number is invalid")
				return
			}
			for i:=min;i<=max;i++{
				numbers = append(numbers,i)
			}

		}else{
			err = fmt.Errorf("range number is invalid")
			return
		}
	}
	return
}