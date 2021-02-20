package utils

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	TIMEZONE                 = "VN"
	LAYOUTTIMEDDMMYYYY       = "02-01-2006"
	LAYOUTTIMEDDMMYYYYHHMMSS = "02-01-2006 15:04:05"
)

func ParseStringToTime(date string) (time.Time, error) {
	layout := "02-01-2006"
	logrus.WithFields(logrus.Fields{}).Infof("[ParseStringToTime] input date %s", date)
	myDate, err := time.Parse(layout, date)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStringToTime] parse error %v", err)
		return time.Now(), errors.New("Thời gian không hợp lệ")
	}
	return myDate, nil
}

var countryTz = map[string]string{
	"VN": "Asia/Ho_Chi_Minh",
}

func TimeIn(inputTime time.Time, name, layout string) string {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[TimeIn] parse error %v", err)
		return time.Now().In(loc).Format("02-01-2006 15:04:05")
	}
	return inputTime.In(loc).Format(layout)
}
