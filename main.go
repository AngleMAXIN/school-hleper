package main

import (
	"fmt"
	"school-hleper/spider"
)

func main() {
	fetch := spider.CreateNewFetch()
	fetch.Prepare()
	fetch.Login("201601013109","maxin123")

	result := fetch.PullScore("")

	for i,v :=  range result{
		fmt.Println(i,v)
	}
}