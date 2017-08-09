package main

import (
	"strconv"
	"unicode"
)

func secondParsing(subject *Subject) {
	// Day and time.
	for idx, s := range subject.SuupTimes {
		var DayAndTime parsedTimeAndClass

		suupTime := []rune(s)
		DayNames := []rune("월화수목금토일")

		isDay := false
		for _, d := range DayNames {
			if suupTime[0] == d {
				isDay = true
			}
		}

		if isDay == false {
			// 시간미지정강좌 혹은 집중강좌
			break
		} else {
			// Day
			switch suupTime[0] {
			case '월':
				DayAndTime.Day = "월"
			case '화':
				DayAndTime.Day = "화"
			case '수':
				DayAndTime.Day = "수"
			case '목':
				DayAndTime.Day = "목"
			case '금':
				DayAndTime.Day = "금"
			case '토':
				DayAndTime.Day = "토"
			case '일':
				DayAndTime.Day = "일"

			default:
				panic("time parsing error. no Day")
			}

			ParseStartAndEndTime(suupTime, &DayAndTime)
		}

		//Classroom parsing
		DayAndTime.Classroom = subject.SuupRoomNms[idx]

		subject.SecondData.TimesAndClass = append(subject.SecondData.TimesAndClass, DayAndTime)
		DayAndTime = parsedTimeAndClass{}
	}

	subject.SecondData.Times_number = strconv.Itoa(len(subject.SecondData.TimesAndClass))

}

func ParsingTime(suupTime []rune, target_time *string, currentIdx *int) {
	for {
		if unicode.IsDigit(suupTime[*currentIdx]) {
			*target_time = string(suupTime[*currentIdx : *currentIdx+2])
			*currentIdx += 2
			break
		} else {
			*currentIdx++
		}
	}
}

func ParseStartAndEndTime(suupTime []rune, DayAndTime *parsedTimeAndClass) {
	currentIdx := 1
	ParsingTime(suupTime, &DayAndTime.Start_time, &currentIdx)
	ParsingTime(suupTime, &DayAndTime.Start_minute, &currentIdx)
	ParsingTime(suupTime, &DayAndTime.End_time, &currentIdx)
	ParsingTime(suupTime, &DayAndTime.End_minute, &currentIdx)
}
