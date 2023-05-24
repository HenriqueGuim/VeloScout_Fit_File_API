package services

import (
	"github.com/tormoder/fit"
)

func GetBestHearthRateAverage(records []*fit.RecordMsg, time uint) uint8 {
	var currSum uint = 0
	var maxAverageHeartRate uint8 = 0

	if time > uint(len(records)) {
		return 0
	}

	for index := uint(0); index < time; index++ {
		if records[index].HeartRate == 255 {
			records[index].HeartRate = 0
		}
		currSum += uint(records[index].HeartRate)
	}

	maxAverageHeartRate = uint8(currSum / time)

	for index := time; index < uint(len(records)); index++ {

		if records[index].HeartRate == 255 {
			records[index].HeartRate = 0
		}

		firstIndex := index - time

		currSum -= uint(records[firstIndex].HeartRate)
		currSum += uint(records[index].HeartRate)
		if uint8(currSum/time) > maxAverageHeartRate {
			maxAverageHeartRate = uint8(currSum / time)
		}

	}

	return maxAverageHeartRate

}

func GetBestPowerAverage(records []*fit.RecordMsg, time uint) uint16 {
	var currSum uint = 0
	var maxAveragePower uint16 = 0

	if time > uint(len(records)) {
		return 0
	}

	for index := uint(0); index < time; index++ {
		if records[index].Power == 65535 {
			records[index].Power = 0
		}
		currSum += uint(records[index].Power)
	}

	maxAveragePower = uint16(currSum / time)

	for index := time; index < uint(len(records)); index++ {

		firstIndex := index - time

		currSum -= uint(records[firstIndex].Power)
		currSum += uint(records[index].Power)

		if uint16(currSum/time) > maxAveragePower {
			maxAveragePower = uint16(currSum / time)
		}

	}

	return maxAveragePower

}
