package main

import (
	"strconv"
)

// This function returns HTML code of subjects table. HTML template package is used.
func drawSubjectTable(subjects []Subject) (string, error) {
	// divide subjects by grade

	subjectsDividedByGrade, err := divideSubjectsByGrade(subjects)
	if err != nil {
		return "", nil
	}

	// we can check subjects has 5th grade or not
	numberOfGrade := len(subjectsDividedByGrade)

	// templating to draw the table
	// TODO: write a table template

	return "", nil
}

func divideSubjectsByGrade(subjects []Subject) (map[int][]Subject, error) {
	var subjectsDividedByGrade map[int][]Subject = make(map[int][]Subject)
	for _, subject := range subjects {
		grade, err := strconv.Atoi(subject.IsuGrade)
		if err != nil {
			return nil, err
		}
		subjectsDividedByGrade[grade] = append(subjectsDividedByGrade[grade], subject)
	}

	return subjectsDividedByGrade, nil
}
