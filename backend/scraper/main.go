// http://github.com/hrzon
// Written by Mjae

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/fedesog/webdriver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run *.go sshLoginID")
		return
	}

	chromeDriver := webdriver.NewChromeDriver("./chromedriver_mac")
	htmlPages := 233 // It means the last number of html page

	err := chromeDriver.Start()
	if err != nil {
		fmt.Println("start():", err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)

	// Go to target page
	if err != nil {
		fmt.Println(err)
	}
	session.Url("https://portal.hanyang.ac.kr/sugang/sulg.do")

	time.Sleep(5 * time.Second)

	change_to_kr, err := session.FindElement(webdriver.ClassName, "lan_ko")
	change_to_kr.Click()

	time.Sleep(3 * time.Second)

	login_popup, err := session.FindElement(webdriver.ID, "btn-user2")
	login_popup.Click()

	// manual login
	fmt.Println("Waiting for login")
	time.Sleep(10 * time.Second)

	// click to get data
	sugangpyunram, err := session.FindElement(webdriver.LinkText, "수강편람")
	sugangpyunram.Click()

	time.Sleep(3 * time.Second)

	johoi, err := session.FindElement(webdriver.ID, "btn_Find")
	johoi.Click()

	time.Sleep(2 * time.Second)

	// make saving directory
	sourceSaveDirectory := "html_sources/"

	_, err = os.Stat(sourceSaveDirectory)
	if err != nil {
		// directory is not exist
		os.Mkdir(sourceSaveDirectory, 0777)
	}

	// time without space bar
	crawlingMoment := fmt.Sprint(time.Now())

	// replace space bars with under bar for convenience
	removeSpaces := []byte(crawlingMoment)
	for i := 0; i < len(removeSpaces); i++ {
		if removeSpaces[i] == ' ' {
			removeSpaces[i] = '_'
		}
	}
	crawlingMoment = string(removeSpaces)

	// current time will be the name of saving directory
	currentSourceSavePath := sourceSaveDirectory + crawlingMoment + "/"

	os.Mkdir(currentSourceSavePath, 0777)

	fmt.Println("currentSourceSavePath: ", currentSourceSavePath)

	// first page
	f, _ := os.Create(currentSourceSavePath + "1.html")
	source, _ := session.Source()

	err = WriteTo(f, source)
	if err == nil {
		fmt.Println("1.html has been saved.")
	} else {
		fmt.Println(err)
	}

	f.Close()

	// other pages
	for i := 2; i <= htmlPages; i++ {
		time.Sleep(1 * time.Second)

		// go to next page
		if i%10 == 1 {
			paging_panel, _ := session.FindElement(webdriver.ID, "pagingPanel")
			arrows, _ := paging_panel.FindElements(webdriver.TagName, "img")

			time.Sleep(2 * time.Second)

			arrows[2].Click()
		} else {
			time.Sleep(1 * time.Second)
			next_page, _ := session.FindElement(webdriver.LinkText, strconv.Itoa(i))
			next_page.Click()
		}

		// save html source of the pages
		filename := strconv.Itoa(i) + ".html"
		f, _ := os.Create(currentSourceSavePath + filename)

		source, _ := session.Source()

		err = WriteTo(f, source)
		if err == nil {
			fmt.Println(filename + " has been saved.")
		} else {
			fmt.Println(err)
		}

		f.Close()

		// wait for preventing network error
		if i%50 == 0 {
			fmt.Println("wait 20 seconds...")
			time.Sleep(20 * time.Second)
			fmt.Println("waiting is done")
		}
	}

	// copy html sources to latest directory
	latestDirectory := sourceSaveDirectory + "latest/"

	os.RemoveAll(latestDirectory)

	_, err = os.Stat(latestDirectory)
	if err != nil {
		// directory is not exist
		os.Mkdir(latestDirectory, 0777)
	}

	// copy html sources to lastest folder
	srcFolder := currentSourceSavePath
	destFolder := latestDirectory
	cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
	err = cpCmd.Run()

	// stop the chromedriver
	time.Sleep(3 * time.Second)
	session.Delete()
	chromeDriver.Stop()

	// start transfering
	fmt.Println("transfering saved html sources to server...")

	sshID := os.Args[1] + "@118.32.156.218"
	serverBackendFolder := "smithy-scheduler/backend/"

	// make directory at server
	makeServerHtmlSourceDirectory := exec.Command("ssh", sshID, "mkdir "+"~/"+serverBackendFolder+sourceSaveDirectory)
	makeServerHtmlSourceDirectory.Run()

	// copy sources to server
	copySourcesToServer := exec.Command("scp", "-r", currentSourceSavePath, sshID+":"+serverBackendFolder+currentSourceSavePath)
	copySourcesToServer.Run()

	// remove old sources at server
	removeLatest := exec.Command("ssh", sshID, "rm -rf "+serverBackendFolder+latestDirectory)
	removeLatest.Run()

	// copy latest sources to server
	copyLatestSourcesToServer := exec.Command("scp", "-r", latestDirectory, sshID+":"+serverBackendFolder+latestDirectory)
	copyLatestSourcesToServer.Run()

	fmt.Println("transfering has been done. Run parsing")

	// Run parsing program in server
	htmlPagesNumberString := strconv.Itoa(htmlPages)
	parsing := exec.Command("ssh", sshID, "cd "+serverBackendFolder+"parser; go run *.go "+htmlPagesNumberString)
	parsing.Run()

	// Parsing has been done.
	fmt.Println("save data to DB")

	// save bson file to database
	saveToDB := exec.Command("ssh", sshID, "mongorestore --drop -d smithy -c sugangInfo "+serverBackendFolder+"/parser/outputs/bson.bson")
	saveToDB.Run()

	// remove sources in local
	fmt.Println("remove sources in local")
	removeSourcesInLocal := exec.Command("rm", "-rf", sourceSaveDirectory)
	removeSourcesInLocal.Run()

	fmt.Println("Done.")
}

func WriteTo(w io.Writer, line string) error {
	if _, err := fmt.Fprintln(w, line); err != nil {
		return err
	}
	return nil
}
