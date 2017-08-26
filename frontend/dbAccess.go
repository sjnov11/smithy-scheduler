package main

import (
	"log"

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
