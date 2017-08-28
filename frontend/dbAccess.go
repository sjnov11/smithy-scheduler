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
	id             bson.ObjectId `json:"_id"`
	IsuGrade       string        `json:"isugrade"`
	BanNo          string        `json:"banno"`
	IsuGbNm        string        `json:"isugbnm"`
	YungyukNm      string        `json:"yungyuknm"`
	SuupNo2        string        `json:"suupno2"`
	HaksuNo        string        `json:"haksuno"`
	GwamokNm       string        `json:"gwamoknm"`
	SuupMsg        string        `json:"suupmsg"`
	AbekGb         string        `json:"abekgb"`
	HakwiNm        string        `json:"hakwinm"`
	DaepyoGangsaNm string        `json:"daepyogangsanm"`
	Hakjeom        string        `json:"hakjeom"`
	IronSigan      string        `json:"ironsigan"`
	SilsSigan      string        `json:"silssigan"`
	SuupTypeNm     string        `json:"suuptypenm"`
	JehanInwon     string        `json:"jehaninwon"`

	SuupTimes   []string `json:"suuptimes"`
	SuupRoomNms []string `json:"suuproomnms"`

	IsuJehanYn   string `json:"isujehanyn"`
	SuupTypeGb   string `json:"suuptypegb"`
	BanSosokNm   string `json:"bansosoknm"`
	GnjSosokNm   string `json:"gnjsosoknm"`
	Best_teacher string `json:"best_teacher"`

	SecondData struct {
		Times_number  string `json:"times_number"`
		TimesAndClass []struct {
			Classroom    string `json:"classroom"`
			Day          string `json:"day"`
			Start_time   string `json:"start_time"`
			Start_minute string `json:"start_minute"`
			End_time     string `json:"end_time"`
			End_minute   string `json:"end_minute"`
		} `json:"timesandclass"`
	} `json:"seconddata"`
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

	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	log.Printf("(getDataFromDBByMajor) Request Major: %s, Subjects Count: %d\n", major, count)

	// get data from query
	// var result []map[string]interface{}
	var result []Subject
	err = query.All(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
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
