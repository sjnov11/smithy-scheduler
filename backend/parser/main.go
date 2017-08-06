package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

var idNames []string

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run *.go pageNumber")
		return
	}

	// global variable init
	idNames = []string{
		"isuGrade",
		"banNo",
		"isuGbNm",
		"yungyukNm",
		"suupNo2",
		"haksuNo",
		"gwamokNm",
		"suupMsg",
		"abekGb",
		"hakwiNm",
		"daepyoGangsaNm",
		"hakjeom",
		"ironSigan",
		"silsSigan",
		"suupTypeNm",
		"jehanInwon",
		"suupTimes",
		"suupRoomNms",
		"isuJehanYn",
		"suupTypeGb",
		"banSosokNm",
		"gnjSosokNm",
	}

	htmlFileNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	filepaths := make([]string, htmlFileNumber)

	html_sources := "../html_sources/"
	for i := 0; i < htmlFileNumber; i++ {
		filepaths[i] = html_sources + "latest/" + strconv.Itoa(i+1) + ".html"
	}

	// Channels
	chSubjects := make(chan Subject)
	chFinished := make(chan bool)

	// Kick off the crawl process (concurrently)
	for _, filepath := range filepaths {
		go crawl(filepath, chSubjects, chFinished)
	}

	// Subscribe to all channels
	var foundSubjects []Subject
	for c := 0; c < len(filepaths); {
		select {
		case subject := <-chSubjects:
			foundSubjects = append(foundSubjects, subject)
		case <-chFinished:
			c++
		}
	}

	// We're done! Print the results...
	fmt.Println("\nFound", len(foundSubjects), "unique subjects:\n")

	// Sort results by GwamokNm
	sort.Sort(ByGwamokNm(foundSubjects))

	// Print
	PrintSubjectsToTextFile(foundSubjects, "outputs/text")
	PrintSubjectsToJsonFile(foundSubjects, "outputs/json")
	PrintSubjectsToBsonFile(foundSubjects, "outputs/bson.bson")

	// Move source to home directory
	makeSourceDirectoryAtHome := exec.Command("mkdir", "~/html_sources")
	makeSourceDirectoryAtHome.Run()

	moveSourcesToHomeDirectory := exec.Command("cp", "-r", html_sources, "~")
	moveSourcesToHomeDirectory.Run()

	// Remove sources
	removeHtmlSources := exec.Command("rm", "-rf", html_sources)
	removeHtmlSources.Run()

	close(chSubjects)
}
