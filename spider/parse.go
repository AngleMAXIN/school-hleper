package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

//  1 2016-2017-1 01100110 地球科学概论 72 4 64 考试 必修 学科基础课程

type itemScore struct {
	ID               string
	OpenDate         string
	CourseID         string
	CourseName       string
	Score            string
	ScoreCredit      string
	ScoreTime        string
	AssessmentMethod string
	AssessmentAttr   string
	CourseNature     string
}

func parseHtml(htmlText string) ([]string, error) {
	//println(htmlText)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlText))
	if err != nil {
		return nil, err
	}

	infoList := make([]string, 0)
	dom.Find("tbody>tr>td").Each(func(i int, selection *goquery.Selection) {
		str := strings.TrimSpace(selection.Text())
		if str != "" && str != " " {
			infoList = append(infoList, str)
		}
	})
	return infoList, nil
}


func buildComplete(source []string) []*itemScore {
	length := len(source)
	if length == 0 {
		log.Println("source list is nil")
		return nil
	}
	if length%10 != 0 {
		fmt.Println("end", length, source[0])
		return nil
	}

	buildResult := make([]*itemScore, 0)
	for i := 0; i < length; i += 10 {
		item := &itemScore{
			ID:               source[i],
			OpenDate:         source[i+1],
			CourseID:         source[i+2],
			CourseName:       source[i+3],
			Score:            source[i+4],
			ScoreCredit:      source[i+5],
			ScoreTime:        source[i+6],
			AssessmentMethod: source[i+7],
			AssessmentAttr:   source[i+8],
			CourseNature:     source[i+9],
		}
		buildResult = append(buildResult, item)
	}
	return buildResult
}
