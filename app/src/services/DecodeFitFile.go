package services

import (
	"fitAPI/app/src/models"
	"fitAPI/app/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tormoder/fit"
)

func DecodeFitFile(context *gin.Context) {
	fileHeader, err := context.FormFile("FitFile")
	if err != nil {
		context.JSON(400, gin.H{
			"message": "No file",
		})
		return
	}

	file, _ := fileHeader.Open()

	defer file.Close()

	fitObj, err := fit.Decode(file)
	if err != nil {
		context.JSON(500, gin.H{
			"message":  "Error decoding file",
			"message2": err.Error(),
		})
		return
	}

	activity, err := fitObj.Activity()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := getData(activity)

	context.JSON(200, data)
}

func getData(activity *fit.ActivityFile) *models.DataFormat {
	data := models.NewDataFormat()

	data.Date = activity.Sessions[0].Timestamp
	data.HeartRate = getHeartRateData(activity)
	data.Power = getPowerData(activity)
	data.Duration = activity.Sessions[0].TotalElapsedTime / 1000
	data.Distance = activity.Sessions[0].GetTotalDistanceScaled()
	data.ElevationGain = activity.Sessions[0].TotalAscent
	data.TSS = activity.Sessions[0].TrainingStressScore
	data.IF = float32(activity.Sessions[0].IntensityFactor) / 1000
	data.AverageCadence = activity.Sessions[0].AvgCadence

	return data
}

func getPowerData(activity *fit.ActivityFile) map[string]uint16 {
	data := make(map[string]uint16)

	data[utils.Avg] = activity.Sessions[0].AvgPower
	data[utils.Max] = activity.Sessions[0].MaxPower
	data[utils.Normalized] = activity.Sessions[0].NormalizedPower
	data[utils.Sec5Str] = GetBestPowerAverage(activity.Records, utils.Sec5)
	data[utils.Sec10Str] = GetBestPowerAverage(activity.Records, utils.Sec10)
	data[utils.Sec12Str] = GetBestPowerAverage(activity.Records, utils.Sec12)
	data[utils.Sec20Str] = GetBestPowerAverage(activity.Records, utils.Sec20)
	data[utils.Sec30Str] = GetBestPowerAverage(activity.Records, utils.Sec30)
	data[utils.Min1Str] = GetBestPowerAverage(activity.Records, utils.Min1)
	data[utils.Min2Str] = GetBestPowerAverage(activity.Records, utils.Min2)
	data[utils.Min5Str] = GetBestPowerAverage(activity.Records, utils.Min5)
	data[utils.Min10Str] = GetBestPowerAverage(activity.Records, utils.Min10)
	data[utils.Min12Str] = GetBestPowerAverage(activity.Records, utils.Min12)
	data[utils.Min20Str] = GetBestPowerAverage(activity.Records, utils.Min20)
	data[utils.Min30Str] = GetBestPowerAverage(activity.Records, utils.Min30)
	data[utils.Min60Str] = GetBestPowerAverage(activity.Records, utils.Min60)
	data[utils.Min90Str] = GetBestPowerAverage(activity.Records, utils.Min90)

	return data

}

func getHeartRateData(activity *fit.ActivityFile) map[string]uint8 {
	data := make(map[string]uint8)

	data[utils.Avg] = activity.Sessions[0].AvgHeartRate
	data[utils.Max] = activity.Sessions[0].MaxHeartRate
	data[utils.Sec5Str] = GetBestHearthRateAverage(activity.Records, utils.Sec5)
	data[utils.Sec10Str] = GetBestHearthRateAverage(activity.Records, utils.Sec10)
	data[utils.Sec12Str] = GetBestHearthRateAverage(activity.Records, utils.Sec12)
	data[utils.Sec20Str] = GetBestHearthRateAverage(activity.Records, utils.Sec20)
	data[utils.Sec30Str] = GetBestHearthRateAverage(activity.Records, utils.Sec30)
	data[utils.Min1Str] = GetBestHearthRateAverage(activity.Records, utils.Min1)
	data[utils.Min2Str] = GetBestHearthRateAverage(activity.Records, utils.Min2)
	data[utils.Min5Str] = GetBestHearthRateAverage(activity.Records, utils.Min5)
	data[utils.Min6Str] = GetBestHearthRateAverage(activity.Records, utils.Min6)
	data[utils.Min10Str] = GetBestHearthRateAverage(activity.Records, utils.Min10)
	data[utils.Min12Str] = GetBestHearthRateAverage(activity.Records, utils.Min12)
	data[utils.Min20Str] = GetBestHearthRateAverage(activity.Records, utils.Min20)
	data[utils.Min30Str] = GetBestHearthRateAverage(activity.Records, utils.Min30)
	data[utils.Min60Str] = GetBestHearthRateAverage(activity.Records, utils.Min60)
	data[utils.Min90Str] = GetBestHearthRateAverage(activity.Records, utils.Min90)

	return data
}
