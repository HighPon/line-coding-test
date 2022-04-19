package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type TaxiLog struct {
	recordTime string
	Mileage    string
}

const (
	BaseDistance               = 1052
	BaseMoney                  = 410
	PerDistance                = 237
	PerMoney                   = 80
	Magnification              = 1000
	MidnightAdditionalDistance = 125
)

func useIoutilReadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func stringToTaxiLog(readFile string) ([]TaxiLog, error) {
	splitNewLine := strings.Split(readFile, "\n")
	taxiLog := make([]TaxiLog, len(splitNewLine), len(splitNewLine))
	for i, v := range splitNewLine {
		tmp := strings.Fields(v)
		r1 := regexp.MustCompile(`\d{2}:\d{2}:\d{2}.\d{3}`)
		r2 := regexp.MustCompile(`\d{1,2}.\d`)
		if r1.MatchString(tmp[0]) && r2.MatchString(tmp[1]) {
			taxiLog[i] = TaxiLog{recordTime: tmp[0], Mileage: tmp[1]}
		} else {
			log.Println(i, tmp[1], r2.MatchString(tmp[1]))
			return nil, errors.New("invalid format")
		}
	}
	return taxiLog, nil
}

func checkMidnightTime(recordTime string) bool {
	tmp := strings.Split(recordTime, ":")[0]
	hour, _ := strconv.Atoi(tmp)
	return ((hour%24) >= 0 && (hour%24) <= 4) || ((hour%24) >= 22 && (hour%24) <= 23)
}

func calcWageFromDistance(distance int) int {
	if distance <= BaseDistance*Magnification {
		return BaseMoney
	} else {
		return BaseMoney + ((distance-BaseDistance*Magnification)/(PerDistance*Magnification))*PerMoney
	}
}

func calcTimeFromZeroWithoutDecimal(timeString string) (retval int) {
	pointSplit := strings.Split(timeString, ".")
	colonSplit := strings.Split(pointSplit[0], ":")
	var time int
	time, _ = strconv.Atoi(colonSplit[0])
	retval += time * 60 * 60 * 1000
	time, _ = strconv.Atoi(colonSplit[1])
	retval += time * 60 * 1000
	time, _ = strconv.Atoi(colonSplit[2])
	retval += time * 1000
	time, _ = strconv.Atoi(pointSplit[1])
	retval += time
	return
}

// secondDistance - firstDistance
func calcTimeBetweenTwo(firstDistance string, secondDistance string) (retval int) {
	firstPointSplit := strings.Split(firstDistance, ".")
	firstTime := calcTimeFromZeroWithoutDecimal(strings.Split(firstPointSplit[0], ":"))

	secondPointSplit := strings.Split(secondDistance, ".")
	secondTime := calcTimeFromZeroWithoutDecimal(strings.Split(secondPointSplit[0], ":"))

	retval = secondTime - firstTime

	firstDecimal, _ := strconv.Atoi(firstPointSplit[1])
	secondDecimal, _ := strconv.Atoi(secondPointSplit[1])

	if firstDecimal > secondDecimal && retval > 0 {
		retval -= 1
	}
	return
}

func calcMoney(taxiLog []TaxiLog) (retval int) {
	for i, _ := range taxiLog {
		var j int
		taxiLog[i].Mileage = strings.Replace(taxiLog[i].Mileage, ".", "", -1)
		j, _ = strconv.Atoi(taxiLog[i].Mileage)
		if i > 0 && checkMidnightTime(taxiLog[i-1].recordTime) && checkMidnightTime(taxiLog[i].recordTime) {
			j *= MidnightAdditionalDistance
		} else {
			j *= 100
		}
		retval += j
	}
	fmt.Println(taxiLog)
	return
}

func main() {
	readFile, err := useIoutilReadFile("test2")
	if err != nil {
		log.Println("err", err)
		return
	}
	taxiLog, err := stringToTaxiLog(readFile)
	if err != nil {
		return
	}
	fmt.Println(taxiLog)
	fmt.Println(calcMoney(taxiLog))
	fmt.Println(calcWageFromDistance(calcMoney(taxiLog)))

}
