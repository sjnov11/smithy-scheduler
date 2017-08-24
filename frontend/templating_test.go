package main

import (
	"fmt"
)

func ExampleDivideSubjectsByGrade() {
	sampleDataNumber := 9
	sampleData := make([]Subject, sampleDataNumber)

	sampleData[0] = Subject{
		IsuGrade: "3",
		GwamokNm: "3학년과목1",
	}
	sampleData[1] = Subject{
		IsuGrade: "3",
		GwamokNm: "3학년과목2",
	}
	sampleData[2] = Subject{
		IsuGrade: "4",
		GwamokNm: "4학년과목1",
	}
	sampleData[3] = Subject{
		IsuGrade: "4",
		GwamokNm: "4학년과목2",
	}
	sampleData[4] = Subject{
		IsuGrade: "3",
		GwamokNm: "3학년과목3",
	}
	sampleData[5] = Subject{
		IsuGrade: "1",
		GwamokNm: "1학년과목1",
	}
	sampleData[6] = Subject{
		IsuGrade: "1",
		GwamokNm: "1학년과목2",
	}
	sampleData[7] = Subject{
		IsuGrade: "2",
		GwamokNm: "2학년과목1",
	}
	sampleData[8] = Subject{
		IsuGrade: "2",
		GwamokNm: "2학년과목2",
	}

	subjectsDividedByGrade, err := divideSubjectsByGrade(sampleData)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= 4; i++ {
		fmt.Printf("%d%s\n", i, "학년:")
		for _, subject := range subjectsDividedByGrade[i] {
			fmt.Println(subject.GwamokNm)
		}
	}

	// Output:
	// 1학년:
	// 1학년과목1
	// 1학년과목2
	// 2학년:
	// 2학년과목1
	// 2학년과목2
	// 3학년:
	// 3학년과목1
	// 3학년과목2
	// 3학년과목3
	// 4학년:
	// 4학년과목1
	// 4학년과목2
}
