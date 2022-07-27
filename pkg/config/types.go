package config

import (
	"errors"
	"strconv"
	"strings"
)

const (
	MB = 1024 * 1024
	KB = 1024
)

type BandWidthQuantity struct {
	s string
	i int64
}


func (q *BandWidthQuantity) Equal(u *BandWidthQuantity) bool{
	if q == nil && u == nil{
		return true
	}
	if q != nil && u != nil{
		return q.i == u.i
	}
	return false
}

func NewBandwidthQuantity(s string) (BandWidthQuantity,error){
	q := BandWidthQuantity{}
	err := q.UnmarshalString(s)
	if err != nil {
		return q,err
	}
	return q,nil
}


func (q *BandWidthQuantity) UnmarshalString(s string) error{
	s = strings.TrimSpace(s)
	if s == ""{
		return nil
	}

	var (
		base int64
		f float64
		err error
	)
	if strings.HasSuffix(s,"MB"){
		base = MB
		fstr := strings.TrimSuffix(s,"MB")
		f,err = strconv.ParseFloat(fstr,64)
		if err != nil {
			return err
		}
	}else if strings.HasSuffix(s,"KB") {
		base = KB
		fstr := strings.TrimSuffix(s,"KB")
		f,err = strconv.ParseFloat(fstr,64)
		if err != nil {
			return err
		}
	}else {
		return errors.New("unit not support")
	}
	q.s = s
	q.i = int64(f * float64(base))
	return nil
}