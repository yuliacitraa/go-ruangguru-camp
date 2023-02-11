package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Study struct {
	StudyName string `json:"study_name"`
	StudyCredit int `json:"study_credit"`
	Grade string `json:"grade"`
}

type Report struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
	Semester int `json:"semester"`
	Studies []Study `json:"studies"`
}

func ReadJSON(filename string) (Report, error) {
	f, err := os.Open(filename)
    if err != nil {
        return Report{}, err
    }

    defer f.Close()

    jsonData, err := ioutil.ReadAll(f)
    if err != nil {
        return Report{}, err
    }

    var r Report

    err = json.Unmarshal(jsonData, &r)
    if err != nil {
        return Report{}, err
    }

    return r, nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0 {
        return 0.0
    }
	
	var totalCredit int
    var totalPoint, gpa float64

    var scale = map[string]float64{
        "A":  4.0,
        "AB": 3.5,
        "B":  3.0,
        "BC": 2.5,
        "C":  2.0,
        "CD": 1.5,
        "D":  1.0,
        "DE": 0.5,
        "E":  0.0,
    }

    for _, study := range report.Studies {
        totalCredit += study.StudyCredit
        totalPoint += float64(study.StudyCredit) * scale[study.Grade]
    }

	gpa = totalPoint / float64(totalCredit)

    return gpa

	// var s Study
	// var totalCredit int
	// var scale, jmlNilai, ip float64
	// switch {
	// case s.Grade == "A":
	// 	scale = 4
	// case s.Grade == "AB":
	// 	scale = 3.5
	// case s.Grade == "B":
	// 	scale = 3
	// case s.Grade == "BC":
	// 	scale = 2.5
	// case s.Grade == "C":
	// 	scale = 2
	// case s.Grade == "CD":
	// 	scale = 1.5
	// case s.Grade == "D":
	// 	scale = 1
	// case s.Grade == "DE":
	// 	scale = 0.5
	// case s.Grade == "E":
	// 	scale = 0
		
	// }

	// jmlNilai = scale * float64(s.StudyCredit)
	
	// for i := 0; i < s.StudyCredit; i++ {
	// 	totalCredit += s.StudyCredit
	// }

	// ip = jmlNilai/float64(totalCredit)
	
	// return ip // TODO: replace this
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
