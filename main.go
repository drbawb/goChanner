package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)


func main() {
	res, err := http.Get("http://ponychan.net/chan/chat")
	if err != nil {
		fmt.Printf("error @ http.Get")
		os.Exit(1)
	}

	parse, err := html.Parse(res.Body)
	if err != nil {
		fmt.Printf("error @ html.parse")
		os.Exit(1)
	}
  threads := buildThreads(parse)
  for ix := 0; ix < len(threads); ix++ {
    fmt.Printf("Author: %s / %s \n Subj: %s \n-----\n", threads[ix].Author.Name, threads[ix].Author.Trip, threads[ix].Subject)
  }  
}

func buildThreads(in *html.Node) []*Thread {
  var threadOut []*Thread

	for ix := 0; ix < len(in.Child); ix++ {
		if in.Child[ix].Type == html.ElementNode && in.Child[ix].Data == "div" {
			for aix := 0; aix < len(in.Child[ix].Attr); aix++ {
				if in.Child[ix].Attr[aix].Key == "class" && in.Child[ix].Attr[aix].Val == "thread" {
					t := new(Thread)
          t.Build(in.Child[ix])
          fmt.Printf("building thread ... \n")
					threadOut = append(threadOut, t)
          fmt.Printf("thread build, len is: %d", len(threadOut))
				}
			}
		}

		threadOut = append(threadOut, buildThreads(in.Child[ix])...)
	}
  
  return threadOut
}

