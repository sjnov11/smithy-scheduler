package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func PrintSubjectsToTextFile(foundSubjects []Subject, filename string) {
	f, _ := os.Create(filename)
	defer f.Close()

	// writing
	for _, subject := range foundSubjects {
		WriteTo(f, subject.String()+"\n")
	}
}

func PrintSubjectsToJsonFile(foundSubjects []Subject, filename string) {
	// json
	j, _ := os.Create(filename)
	defer j.Close()

	// writing
	for _, subject := range foundSubjects {
		jsonString, err := json.Marshal(subject)
		if err != nil {
			panic("json error")
		}
		WriteTo(j, jsonPrettyPrint(string(jsonString)))
	}
}

func PrintSubjectsToBsonFile(foundSubjects []Subject, filename string) {
	// bson
	b, _ := os.Create(filename)
	defer b.Close()

	// writing
	for _, subject := range foundSubjects {
		bsonBytes, _ := bson.Marshal(subject)
		b.Write(bsonBytes)
	}
}

func WriteTo(w io.Writer, line string) error {
	if _, err := fmt.Fprintln(w, line); err != nil {
		return err
	}
	return nil
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
