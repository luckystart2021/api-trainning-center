package utils

import (
	"time"

	"github.com/sirupsen/logrus"
)

func ParseStringToTime(date string) (time.Time, error) {
	layout := "02-01-2006"
	logrus.WithFields(logrus.Fields{}).Infof("[ParseStringToTime] input date %s", date)
	myDate, err := time.Parse(layout, date)
	if err != nil {
		logrus.WithFields(logrus.Fields{}).Errorf("[ParseStringToTime] parse error %v", err)
		return time.Now(), err
	}
	return myDate, nil
}
