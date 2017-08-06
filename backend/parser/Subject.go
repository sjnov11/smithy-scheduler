package main

import "fmt"

// ByGwamokNm implements sort.Interface for []Subject based on
// the GwamokNm field.
type ByGwamokNm []Subject

func (a ByGwamokNm) Len() int           { return len(a) }
func (a ByGwamokNm) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByGwamokNm) Less(i, j int) bool { return a[i].GwamokNm < a[j].GwamokNm }

type parsedTimeAndClass struct {
	Classroom    string `json:"classroom"`
	Day          string `json:"day"`
	Start_time   string `json:"start_time"`
	Start_minute string `json:"start_minute"`
	End_time     string `json:"end_time"`
	End_minute   string `json:"end_minute"`
}

type SubParsingData struct {
	Times_number  string               `json:"times_number"`
	TimesAndClass []parsedTimeAndClass `json:"timesAndClass"`
}

type Subject struct {
	IsuGrade       string         `json:"isuGrade"`
	BanNo          string         `json:"banNo"`
	IsuGbNm        string         `json:"isuGbNm"`
	YungyukNm      string         `json:"yungyukNm"`
	SuupNo2        string         `json:"suupNo2"`
	HaksuNo        string         `json:"haksuNo"`
	GwamokNm       string         `json:"gwamokNm"`
	SuupMsg        string         `json:"suupMsg"`
	AbekGb         string         `json:"abekGb"`
	HakwiNm        string         `json:"hakwiNm"`
	DaepyoGangsaNm string         `json:"daepyoGangsaNm"`
	Hakjeom        string         `json:"hakjeom"`
	IronSigan      string         `json:"ironSigan"`
	SilsSigan      string         `json:"silsSigan"`
	SuupTypeNm     string         `json:"suupTypeNm"`
	JehanInwon     string         `json:"jehanInwon"`
	SuupTimes      []string       `json:"suupTimes"`
	SuupRoomNms    []string       `json:"suupRoomNms"`
	IsuJehanYn     string         `json:"isuJehanYn"`
	SuupTypeGb     string         `json:"suupTypeGb"`
	BanSosokNm     string         `json:"banSosokNm"`
	GnjSosokNm     string         `json:"gnjSosokNm"`
	Best_teacher   string         `json:"best_teacher"`
	SecondData     SubParsingData `json:"secondData"`
}

func (subject Subject) String() string {
	var rt string
	rt += fmt.Sprintln("isuGrade:", subject.IsuGrade)
	rt += fmt.Sprintln("banNo:", subject.BanNo)
	rt += fmt.Sprintln("isuGbNm:", subject.IsuGbNm)
	rt += fmt.Sprintln("yungyukNm:", subject.YungyukNm)
	rt += fmt.Sprintln("suupNo2:", subject.SuupNo2)
	rt += fmt.Sprintln("haksuNo:", subject.HaksuNo)
	rt += fmt.Sprintln("gwamokNm:", subject.GwamokNm)
	rt += fmt.Sprintln("suupMsg:", subject.SuupMsg)
	rt += fmt.Sprintln("abekGb:", subject.AbekGb)
	rt += fmt.Sprintln("hakwiNm:", subject.HakwiNm)
	rt += fmt.Sprintln("daepyoGangsaNm:", subject.DaepyoGangsaNm)
	rt += fmt.Sprintln("hakjeom:", subject.Hakjeom)
	rt += fmt.Sprintln("ironSigan:", subject.IronSigan)
	rt += fmt.Sprintln("silsSigan:", subject.SilsSigan)
	rt += fmt.Sprintln("suupTypeNm:", subject.SuupTypeNm)
	rt += fmt.Sprintln("jehanInwon:", subject.JehanInwon)
	rt += fmt.Sprintln("suupTimes:", subject.SuupTimes)
	rt += fmt.Sprintln("suupRoomNms:", subject.SuupRoomNms)
	for idx, s := range subject.SuupRoomNms {
		rt += fmt.Sprintf("suupRoomNms[%d]:%s\n", idx, s)
	}
	rt += fmt.Sprintln("isuJehanYn:", subject.IsuJehanYn)
	rt += fmt.Sprintln("suupTypeGb:", subject.SuupTypeGb)
	rt += fmt.Sprintln("banSosokNm:", subject.BanSosokNm)
	rt += fmt.Sprintln("gnjSosokNm:", subject.GnjSosokNm)
	rt += fmt.Sprintln("best_teacher:", subject.Best_teacher)

	rt += fmt.Sprintln("secondData.times_number", subject.SecondData.Times_number)
	rt += fmt.Sprintln("secondData.timesAndClass", subject.SecondData.TimesAndClass)
	for idx, time := range subject.SecondData.TimesAndClass {
		rt += fmt.Sprintln("secondData.timesAndClass[", idx, "].classroom:", time.Classroom)
		rt += fmt.Sprintln("secondData.timesAndClass[", idx, "].day:", time.Day)
		rt += fmt.Sprintln("secondData.timesAndClass[", idx, "].start_time:", time.Start_time)
		rt += fmt.Sprintln("secondData.timesAndClass[", idx, "].start_minute:", time.Start_minute)
		rt += fmt.Sprintln("secondData.timesAndClass[", idx, "].end_time:", time.End_time)
		rt += fmt.Sprintln("secondData.timesAndClass[", idx, "].end_minute:", time.End_minute)
	}

	return rt
}
