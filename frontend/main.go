package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var currentPath string
var mainPageHTMLBuffer bytes.Buffer

func main() {
	log.Println("(main) The Server starts.")
	// initialize logger
	logFile := logInit()
	defer logFile.Close()

	log.Println("(main) The Logger has been initialized.")

	err := generateMainPageHTML(&mainPageHTMLBuffer)
	if err != nil {
		panic(err)
	}

	log.Println("(main) The main page source code has been generated.")
	log.Println("(main) Waiting request... ")

	http.HandleFunc("/", handler)
	http.HandleFunc("/db/", dbHandler)
	http.HandleFunc("/db/getDataByMajor", sendDataByMajorHandler)
	http.HandleFunc("/db/getSubjectTable", sendSubjectTableHandler)
	http.ListenAndServe(":8080", nil)
}

func logInit() *os.File {
	currentPath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}

	// setting logger
	logDir := "log"
	_, err = os.Stat(logDir)
	if err != nil {
		// directory is not exist
		err = os.Mkdir("log", 0755)
		if err != nil {
			panic(err)
		}
	}

	logFile, err := os.Create(currentPath + "/log/" + time.Now().String())
	if err != nil {
		panic(err)
	}

	// print log to console and logFile simultaneously
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// set default logger
	log.SetOutput(multiWriter)

	return logFile
}

func generateMainPageHTML(w io.Writer) error {
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
		panic(err)
	}

	return tmpl.Execute(w, MajorList)
}
