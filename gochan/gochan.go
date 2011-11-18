/*
main.go: This is the entry point. It will eventually switch between the CLI toolkit and browser UI.
Currently, though, it's pretty much a testbed for thread.go

See doc/LICENSE for licensing restrictions.

"It's not the way you plan it, it's how you make it happen."
*/

package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
  "gochanner/thread"
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
    fmt.Printf("id: %s Author: %s / %s \n Subj: %s \n-----\n", threads[ix].ThreadNo, threads[ix].Author.Name, threads[ix].Author.Trip, threads[ix].Subject)
  }
}

func buildThreads(in *html.Node) []*thread.Thread {
  var threadOut []*thread.Thread

  //outer loop over in (usually a full HTML document parse tree)
  //looking for <div class="thread"> or equivalent; this should be handled by the board driver.
  for ix := 0; ix < len(in.Child); ix++ {
    if in.Child[ix].Type == html.ElementNode && in.Child[ix].Data == "div" {
			//loop over the div's attribute array to determine it's CSS class
      for aix := 0; aix < len(in.Child[ix].Attr); aix++ {
				if in.Child[ix].Attr[aix].Key == "class" && in.Child[ix].Attr[aix].Val == "thread" {
					t := new(thread.Thread)
          t.Build(in.Child[ix])
					threadOut = append(threadOut, t)
				}
			}
		}

		threadOut = append(threadOut, buildThreads(in.Child[ix])...) //do one more append so we don't lose the results on a long forgotten stack
	}

  return threadOut
}

