package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// data type
/* type Subject struct {
 *   isugrade       string
 *   banno          string
 *   isugbnm        string
 *   yungyuknm      string
 *   suupno2        string
 *   haksuno        string
 *   gwamoknm       string
 *   suupmsg        string
 *   abekgb         string
 *   hakwinm        string
 *   daepyogangsanm string
 *   hakjeom        string
 *   ironsigan      string
 *   silssigan      string
 *   suuptypenm     string
 *   jehaninwon     string
 *   suuptimes      []string
 *   suuproomnms    []string
 *   isujehanyn     string
 *   suuptypegb     string
 *   bansosoknm     string
 *   gnjsosoknm     string
 *   best_teacher   string
 *   seconddata     struct {
 *     times_number  string
 *     timesandclass []struct {
 *       classroom    string
 *       day          string
 *       start_time   string
 *       start_minute string
 *       end_time     string
 *       end_minute   string
 *     }
 *   }
 * } */

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
	fmt.Println(count)

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
	return majorList_string, nil
}

func createJSCodeWithTemplate(majorList []string, filename string) error {

	// prepare data type for filling the template
	type MajorName struct {
		MajorName string
	}

	var majorNameSlice []MajorName
	for _, name := range majorList {
		var tempMajorName MajorName
		// add syntax for javascript
		tempMajorName.MajorName = "'" + name + "',"

		majorNameSlice = append(majorNameSlice, tempMajorName)
	}

	tmpl, err := template.ParseFiles("./majorListTemplate.js")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl.Execute(f, majorNameSlice)
	return nil
}

func main() {
	// get major list from db
	majorList, err := getMajorListFromDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing majorList slice: ")
	fmt.Println(majorList)

	filename := "majorList.js"
	err = createJSCodeWithTemplate(majorList, filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("generating majorList.js has been completed.")
}
