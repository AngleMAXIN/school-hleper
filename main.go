package main

import (
	"fmt"
	"school-hleper/spider"
)

func main() {
	fetch := spider.CreateNewFetch()
	fetch.Prepare()
	fetch.Login("","")

	result := fetch.PullScore("")

	for i,v :=  range result{
		fmt.Println(i,v)
	}
}