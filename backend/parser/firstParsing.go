package main

import "golang.org/x/net/html"

func BestTeacherCheck(z *html.Tokenizer, subject *Subject) {
	next := z.Next()
	if next == html.SelfClosingTagToken {
		value := z.Token()
		if value.Data == "img" {
			subject.Best_teacher = "true"
		}
	}
}

func ParsingSuupTimes(z *html.Tokenizer, subject *Subject, meaningValue *html.Token) {
	subject.SuupTimes = append(subject.SuupTimes, meaningValue.String())

	//check whehter suup is more than one time.
	for {
		next := z.Next()
		if next == html.SelfClosingTagToken {
			continue
		} else if next == html.TextToken {
			value := z.Token()
			subject.SuupTimes = append(subject.SuupTimes, value.String())
		} else {
			break
		}
	}
}
func ParsingClassrooms(z *html.Tokenizer, subject *Subject, meaningValue *html.Token) {
	cls := meaningValue.String()
	Classrooms := []rune(cls)
	// fmt.Println(Classrooms)

	startIdx := 0
	endIdx := 0
	for idx, c := range []rune(Classrooms) {
		if string(c) == "," {
			endIdx = idx
			subject.SuupRoomNms = append(subject.SuupRoomNms, string(Classrooms[startIdx:endIdx]))
			startIdx = endIdx + 1
		}
	}

	subject.SuupRoomNms = append(subject.SuupRoomNms, string(Classrooms[startIdx:]))
}

func SaveMeaningValueToSubjectStruct(z *html.Tokenizer, idName string, subject *Subject, meaningValue *html.Token, sendSubject *bool) {
	switch idName {
	case idNames[0]:
		subject.IsuGrade = meaningValue.String()
	case idNames[1]:
		subject.BanNo = meaningValue.String()
	case idNames[2]:
		subject.IsuGbNm = meaningValue.String()
	case idNames[3]:
		subject.YungyukNm = meaningValue.String()
	case idNames[4]:
		subject.SuupNo2 = meaningValue.String()
	case idNames[5]:
		subject.HaksuNo = meaningValue.String()
	case idNames[6]:
		subject.GwamokNm = meaningValue.String()
	case idNames[7]:
		subject.SuupMsg = meaningValue.String()
	case idNames[8]:
		subject.AbekGb = meaningValue.String()
	case idNames[9]:
		subject.HakwiNm = meaningValue.String()
	case idNames[10]:
		subject.DaepyoGangsaNm = meaningValue.String()
		BestTeacherCheck(z, subject)
	case idNames[11]:
		subject.Hakjeom = meaningValue.String()
	case idNames[12]:
		subject.IronSigan = meaningValue.String()
	case idNames[13]:
		subject.SilsSigan = meaningValue.String()
	case idNames[14]:
		subject.SuupTypeNm = meaningValue.String()
	case idNames[15]:
		subject.JehanInwon = meaningValue.String()
	case idNames[16]:
		ParsingSuupTimes(z, subject, meaningValue)
	case idNames[17]:
		ParsingClassrooms(z, subject, meaningValue)
	case idNames[18]:
		subject.IsuJehanYn = meaningValue.String()
	case idNames[19]:
		subject.SuupTypeGb = meaningValue.String()
	case idNames[20]:
		subject.BanSosokNm = meaningValue.String()
	case idNames[21]:
		subject.GnjSosokNm = meaningValue.String()

		// The end of parsing.
		// Exception. if Best_teacher has no info, insert "false".
		if subject.Best_teacher == "" {
			subject.Best_teacher = "false"
		}

		// parsing end. send Subject to channel
		*sendSubject = true

	default:
		panic("error in switch!")
	}

}
