package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Subject struct {
	id             bson.ObjectId `bson:"_id"`
	IsuGrade       string        `bson:"isugrade"`
	BanNo          string        `bson:"banno"`
	IsuGbNm        string        `bson:"isugbnm"`
	YungyukNm      string        `bson:"yungyuknm"`
	SuupNo2        string        `bson:"suupno2"`
	HaksuNo        string        `bson:"haksuno"`
	GwamokNm       string        `bson:"gwamoknm"`
	SuupMsg        string        `bson:"suupmsg"`
	AbekGb         string        `bson:"abekgb"`
	HakwiNm        string        `bson:"hakwinm"`
	DaepyoGangsaNm string        `bson:"daepyogangsanm"`
	Hakjeom        string        `bson:"hakjeom"`
	IronSigan      string        `bson:"ironsigan"`
	SilsSigan      string        `bson:"silssigan"`
	SuupTypeNm     string        `bson:"suuptypenm"`
	JehanInwon     string        `bson:"jehaninwon"`

	SuupTimes   []string `bson:"suuptimes"`
	SuupRoomNms []string `bson:"suuproomnms"`

	IsuJehanYn          string `bson:"isujehanyn"`
	SuupTypeGb          string `bson:"suuptypegb"`
	BanSosokNm          string `bson:"bansosoknm"`
	GnjSosokNm          string `bson:"gnjsosoknm"`
	Best_teacher_string string `bson:"best_teacher"`
	Best_teacher        bool   `bson:"-"`

	SecondData struct {
		Times_number  string `bson:"times_number"`
		TimesAndClass []struct {
			Classroom    string `bson:"classroom"`
			Day          string `bson:"day"`
			Start_time   string `bson:"start_time"`
			Start_minute string `bson:"start_minute"`
			End_time     string `bson:"end_time"`
			End_minute   string `bson:"end_minute"`
		} `bson:"timesandclass"`
	} `bson:"seconddata"`
}

// methods for sorting
type BySubjectName []Subject

func (s BySubjectName) Len() int {
	return len(s)
}

func (s BySubjectName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s BySubjectName) Less(i, j int) bool {
	if strings.Compare(s[i].GwamokNm, s[j].GwamokNm) < 0 {
		return true
	} else {
		return false
	}
}

type ByProfessorName []Subject

func (s ByProfessorName) Len() int {
	return len(s)
}

func (s ByProfessorName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByProfessorName) Less(i, j int) bool {
	if strings.Compare(s[i].DaepyoGangsaNm, s[j].DaepyoGangsaNm) < 0 {
		return true
	} else {
		return false
	}
}

type BySuupNo2 []Subject

func (s BySuupNo2) Len() int {
	return len(s)
}

func (s BySuupNo2) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s BySuupNo2) Less(i, j int) bool {
	if strings.Compare(s[i].SuupNo2, s[j].SuupNo2) < 0 {
		return true
	} else {
		return false
	}
}

func getDataFromDBByMajor(major string) ([]Subject, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// get query from db
	sugangInfo := session.DB("smithy").C("sugangInfo")
	query := sugangInfo.Find(bson.M{"bansosoknm": major})

	/* count, err := query.Count()
	 * if err != nil {
	 *   return nil, err
	 * }
	 * log.Printf("(getDataFromDBByMajor) Request Major: %s, Subjects Count: %d\n", major, count) */

	// get data from query
	// var result []map[string]interface{}
	var result []Subject
	err = query.All(&result)
	if err != nil {
		return nil, err
	}

	// remove same lecture (same SuupNo2)
	sort.Sort(BySuupNo2(result))
	var nonDuplicatedResult []Subject
	formerSuupNo2 := result[0].SuupNo2
	nonDuplicatedResult = append(nonDuplicatedResult, result[0])

	for _, subject := range result {
		if formerSuupNo2 == subject.SuupNo2 {
			// skip when number is same
			continue
		} else {
			//number is not same
			formerSuupNo2 = subject.SuupNo2
			nonDuplicatedResult = append(nonDuplicatedResult, subject)
		}
	}

	// check Best_teacher
	for idx := range nonDuplicatedResult {
		if nonDuplicatedResult[idx].Best_teacher_string[0] == 't' {
			nonDuplicatedResult[idx].Best_teacher = true
		} else {
			nonDuplicatedResult[idx].Best_teacher = false
		}
	}

	// log.Printf("(getDataFromDBByMajor) Subjects Count after duplication is removed: %d\n", len(nonDuplicatedResult))
	log.Printf("(getDataFromDBByMajor) Request Major: %s, Subjects Count: %d\n", major, len(nonDuplicatedResult))

	return nonDuplicatedResult, nil
}

func getMajorListFromDB() ([]string, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// get query from db
	sugangInfo := session.DB("smithy").C("sugangInfo")
	query := sugangInfo.Find(bson.M{})

	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	// fmt.Println(count)

	// get data from query
	var result []map[string]interface{}
	err = query.All(&result)
	if err != nil {
		return nil, err
	}

	// extract majors from data
	var majorList map[string]struct{}
	majorList = make(map[string]struct{}, count)

	// use map to remove duplication
	for _, data := range result {
		major := data["bansosoknm"]
		majorList[fmt.Sprint(major)] = struct{}{}
	}

	// transform non-duplicated-major-names to string type
	var majorList_string []string
	for major := range majorList {
		majorList_string = append(majorList_string, major)
	}

	// sorted major list
	sort.Strings(majorList_string)

	log.Println("(getMajorListFromDB) The subjects list has been generated.")
	return majorList_string, nil
}
