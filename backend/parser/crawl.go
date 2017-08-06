package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// return true and idName if the attributes
func attrIsExist(t html.Token) (bool, string) {
	for _, a := range t.Attr {
		for _, idName := range idNames {
			if a.Key == "id" && a.Val == idName {
				return true, idName
			}
		}
	}
	return false, ""
}

// Extract all http** links from a given webpage
func crawl(filepath string, ch chan Subject, chFinished chan bool) {

	f, err := os.Open(filepath)
	defer f.Close()

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + filepath + "\"")
		return
	}

	z := html.NewTokenizer(f)

	var subject Subject
	var sendSubject bool
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <td> tag
			isTd := t.Data == "td"
			if !isTd {
				continue
			}

			isExist, idName := attrIsExist(t)

			// idName(like "gwanmokNm", "banNo") is exist!
			if isExist {
				next := z.Next()

				// start parsing
				if next == html.TextToken {
					meaningValue := z.Token()
					SaveMeaningValueToSubjectStruct(z, idName, &subject, &meaningValue, &sendSubject)
				} else {
					// when it is not a TextToken
					// do nothing
				}
			} else {
				// idName is not exist
				// do nothing
			}

			// The end of parsing
			if sendSubject == true {
				// before sending the data, second parsing is needed.

				secondParsing(&subject)

				// send subject to channel
				ch <- subject

				// reinitialize
				sendSubject = false
				subject = Subject{}
			} else {
				// parsing has not been completed.
				// do nothing. Iterative for loop
			}
		}
	}
}
