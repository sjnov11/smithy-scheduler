package main

import (
	"bytes"
	"html/template"
	"io"
	"sort"
	"strconv"
)

func writeMainPageHTML(w io.Writer) error {
	MajorList := struct {
		Majors []string
	}{
		// TODO: DB에서 전공 정보 받아오게 코드 고쳐야겠다. 서버 로딩할 때 한 번만 하니까 DB에서 가져와도 상관없겠지
		Majors: []string{
			"STS(과학기술학) 전공",
			"STS(과학기술학)전공",
			"간호학과",
			"간호학과(야)",
			"건설환경공학과",
			"건축공학부",
			"건축학부",
			"경영학부",
			"경제금융학부",
			"고전읽기융합전공",
			"공공수행인문학전공",
			"관광학부",
			"관현악과",
			"교육공학과",
			"교육학과",
			"국악과",
			"국어교육과",
			"국어국문학과",
			"국제학부",
			"글로벌비즈니스문화전공(영어전용)",
			"기계공학부",
			"도시공학과",
			"독어독문학과",
			"무용학과",
			"물리학과",
			"미디어문화전공",
			"미디어커뮤니케이션학과",
			"미래인문학융합전공",
			"미래자동차공학과",
			"사학과",
			"사회학과",
			"산업공학과",
			"산업융합학부",
			"생명공학과",
			"생명공학전공",
			"생명과학과",
			"생체공학전공",
			"서울 대학",
			"성악과",
			"소프트웨어전공",
			"수학과",
			"수학교육과",
			"스포츠산업학과",
			"식품영양학과",
			"신소재공학부",
			"실내건축디자인학과",
			"에너지공학과",
			"연극영화학과",
			"영어교육과",
			"영어영문학과",
			"영어커뮤니케이션전공",
			"원자력공학과",
			"유기나노공학과",
			"융합전자공학부",
			"응용미술교육과",
			"응용시스템전공",
			"의류학과",
			"의예과",
			"인문공공행정전공",
			"인문소프트웨어융합전공",
			"자동차-SW융합전공",
			"자원환경공학과",
			"작곡과",
			"전기·생체공학부",
			"전기공학전공",
			"정보시스템학과",
			"정보융합전공",
			"정책학과",
			"정치외교학과",
			"중국경제통상전공",
			"중어중문학과",
			"창업융합전공",
			"철학과",
			"체육학과",
			"컴퓨터소프트웨어학부",
			"컴퓨터전공",
			"통상한국어커뮤니케이션전공",
			"파이낸스경영학과",
			"피아노과",
			"한중통번역전공",
			"행정학과",
			"화학공학과",
			"화학공학전공",
			"화학과",
		},
	}

	tmpl, err := template.ParseFiles("./template/index.html")
	if err != nil {
		return err
	}

	tmpl.Execute(w, MajorList)
	return nil
}

// This function returns HTML code of subjects table. HTML template package is used.
func drawSubjectTable(subjects []Subject) (string, error) {
	var data struct {
		GradeNames             []string
		SubjectsDividedByGrade map[int][]Subject

		SubjectsOrderedForTable [][]Subject
		IsGradeEmpty            []bool
	}
	// divide subjects by grade

	var err error
	data.SubjectsDividedByGrade, err = divideSubjectsByGrade(subjects)
	if err != nil {
		return "", err
	}

	// check which grade has no subjects
	for i := 0; i <= 5; i++ {
		if len(data.SubjectsDividedByGrade[i]) == 0 {
			data.IsGradeEmpty = append(data.IsGradeEmpty, true)
		} else {
			data.IsGradeEmpty = append(data.IsGradeEmpty, false)
		}
	}

	// we can check whether subjects has 5th grade or not
	// and there may be cases that does not have 1th, 2nd, 3th or 4rd grade
	// numberOfGrade := len(data.SubjectsDividedByGrade)

	// when 0th grade is exist (0th grade means that here is no information about grade)

	// Fill GradeNames
	data.GradeNames = append(data.GradeNames, "학년정보없음")
	// from 1st grade to 5th grade
	for i := 1; i <= 5; i++ {
		data.GradeNames = append(data.GradeNames, strconv.Itoa(i)+"학년")
	}

	// make each grade's subject count same to fill blank data.
	// get maximum count
	var subjectCounts []int
	for _, subjects := range data.SubjectsDividedByGrade {
		subjectCounts = append(subjectCounts, len(subjects))
	}
	// sort to find maximum count
	sort.Sort(sort.Reverse(sort.IntSlice(subjectCounts)))
	// fmt.Println(subjectCounts)
	maximumSubjectCount := subjectCounts[0]

	// fill the blank data
	// 일단 꽉채워놓자 여긴 map이라서 괜찮다
	for idx := 0; idx <= 5; idx++ {
		for len(data.SubjectsDividedByGrade[idx]) < maximumSubjectCount {
			data.SubjectsDividedByGrade[idx] = append(data.SubjectsDividedByGrade[idx], Subject{
				GwamokNm: " ",
			})
		}
	}

	// for check whether each grade have same number of subjects
	/* subjectCounts = nil
	 * for _, subjects := range data.subjectsDividedByGrade {
	 *   subjectCounts = append(subjectCounts, len(subjects))
	 * }
	 * // sort to find maximum count
	 * sort.Sort(sort.Reverse(sort.IntSlice(subjectCounts)))
	 * fmt.Println(subjectCounts) */

	// fill data for table
	// 건축학부 5학년 처리. 여기서 터지네
	/* for rowNumber := 0; rowNumber < maximumSubjectCount; rowNumber++ {
	 *   var row []Subject
	 *   for grade := 1; grade <= numberOfGrade; grade++ {
	 *     // for debugging
	 *     row = append(row, data.SubjectsDividedByGrade[grade][rowNumber])
	 *   }
	 *   data.SubjectsOrderedForTable = append(data.SubjectsOrderedForTable, row)
	 * } */

	// fill the last data
	for rowNumber := 0; rowNumber < maximumSubjectCount; rowNumber++ {
		var row []Subject
		for grade := 0; grade <= 5; grade++ {
			// for debugging
			if len(data.SubjectsDividedByGrade[grade]) == 0 {
				row = append(row, Subject{})
			} else {
				row = append(row, data.SubjectsDividedByGrade[grade][rowNumber])
			}
		}
		data.SubjectsOrderedForTable = append(data.SubjectsOrderedForTable, row)
	}

	// remove empty grade and subjects
	for grade, gradeIsEmpty := range data.IsGradeEmpty {
		if gradeIsEmpty {
			data.GradeNames[grade] = ""
			for row := 0; row < maximumSubjectCount; row++ {
				data.SubjectsOrderedForTable[row][grade] = Subject{}
			}
		}
	}

	// templating to draw the table
	tmpl, err := template.ParseFiles("./template/subjectTable.html")
	if err != nil {
		return "", err
	}

	// generate html code and save it to source variable
	var source bytes.Buffer
	err = tmpl.Execute(&source, data)
	if err != nil {
		return "", err
	}

	return source.String(), nil
}

func divideSubjectsByGrade(subjects []Subject) (map[int][]Subject, error) {
	var subjectsDividedByGrade map[int][]Subject = make(map[int][]Subject)
	for _, subject := range subjects {
		// TODO: Issue. If IsuGrade has empty value, it is set to first grade.
		var grade int
		var err error
		if subject.IsuGrade != "" {
			grade, err = strconv.Atoi(subject.IsuGrade)
			if err != nil {
				return nil, err
			}
		} else {
			// 0th grade means that there is no grade
			grade = 0
		}
		subjectsDividedByGrade[grade] = append(subjectsDividedByGrade[grade], subject)
	}

	return subjectsDividedByGrade, nil
}
