package utils

import "time"

func StartOfDay(dateInMillSec int64) int64 {
	var st time.Time

	if dateInMillSec > 0 {
		st = time.UnixMilli(dateInMillSec)
	} else {
		st = time.Now()
	}

	today := st
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.Local)

	return startOfDay.UnixMilli()
}

func EndOfDay(dateInMillSec int64) int64 {
	var st time.Time

	if dateInMillSec > 0 {
		st = time.UnixMilli(dateInMillSec)
	} else {
		st = time.Now()
	}

	today := st
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 23, 59, 59, 999, time.Local)

	return startOfDay.UnixMilli()
}

type DateStamp struct {
	baseTimeStamp int64
}

func NewDateStampWithBaseTime(baseTimeStamp int64) *DateStamp {
	return &DateStamp{
		baseTimeStamp: baseTimeStamp,
	}
}

func NewDateStamp() *DateStamp {
	return &DateStamp{
		baseTimeStamp: time.Now().UnixMilli(),
	}

}

// func (dst *DateStamp) StartOfDay() (int64, string) {
// 	sod := StartOfDay(dst.baseTimeStamp)
// 	return sod, time.UnixMilli(sod).String()
// }

// func (dst *DateStamp) EndOfDay() (int64, string) {
// 	eod := EndOfDay(dst.baseTimeStamp)
// 	return eod, time.UnixMilli(eod).String()
// }

func (dst *DateStamp) StartOfDay() *DateStamp {
	sod := StartOfDay(dst.baseTimeStamp)
	dst.baseTimeStamp = sod
	return dst
}

func (dst *DateStamp) EndOfDay() *DateStamp {
	eod := EndOfDay(dst.baseTimeStamp)
	dst.baseTimeStamp = eod
	return dst
}

func (dst *DateStamp) AddDays(numberOfDays int32) *DateStamp {
	dst.baseTimeStamp += 1000 * 60 * 60 * 24 * int64(numberOfDays)

	return dst
}

func (dst *DateStamp) AddHours(numberOfHours int32) *DateStamp {
	dst.baseTimeStamp += 1000 * 60 * 60 * int64(numberOfHours)

	return dst
}

func (dst *DateStamp) AddMinutes(numberOfMinutes int32) *DateStamp {
	dst.baseTimeStamp += 1000 * 60 * int64(numberOfMinutes)

	return dst
}

func (dst *DateStamp) AddSeconds(numberOfSeconds int32) *DateStamp {
	dst.baseTimeStamp += 1000 * 60 * 60 * int64(numberOfSeconds)

	return dst
}

func (dst *DateStamp) Value() (int64, string) {
	return dst.baseTimeStamp, time.UnixMilli(dst.baseTimeStamp).String()
}
