package main

import (
	//"fmt"
	"html"
  "strings"
)

// Very simple tree structure of posts
// We will eventually analyze posts for `>>` reply forwards
// These reply forwards will determine the ID of the parent
// Then we can insert the post as a child by looking for the parent in the tree.
type Thread struct {
	Node          *html.Node
	Author        Author
	Subject, Body string
	ThreadNo      string
}

type Author struct {
	Name, Trip string
}

//Attempts to build the thread from a DOM tree
func (t *Thread) Build(in *html.Node) {
	t.Node = in
	t.ExtractMeta() //make functional later
}

//Extracts meta-data from a thread that has an underlying DOM tree, returns err. otherwise.
func (t *Thread) ExtractMeta() {
	for ix := 0; ix < len(t.Node.Child); ix++ {
		c := t.Node.Child[ix]

		if c.Type == html.ElementNode && c.Data == "label" {
			t.extractMetaSpans(c)
		}
	}
}

func (t *Thread) extractMetaSpans(in *html.Node) {
  for ix := 0; ix < len(in.Child); ix++ {
    c := in.Child[ix]
    if c.Type == html.ElementNode && c.Data == "span" {
      for aix := 0; aix < len(c.Attr); aix++ {
        if c.Attr[aix].Key == "class" {
          switch c.Attr[aix].Val {
          case "postername":
            t.Author.Name = t.extAuthor(c)
            //fmt.Printf("author name: %s \n", t.Author.Name)
          case "filetitle":
            t.Subject = t.extSubj(c)
            //fmt.Printf("subj: %s \n", t.Subject)
          case "postertrip":
            //fmt.Printf("getting trip")
            t.Author.Trip = t.extTrip(c)
          }
        }
      }
    }     
  }

}

func (t *Thread) extTrip(in *html.Node) string {
  out := "DEBUG-default"

  out = strings.TrimLeft(in.Child[0].Data, "\r\n")

  return out
}

func (t *Thread) extSubj(in *html.Node) string {
	out := "DEBUG-default"

	//go over all nodes looking for
	out = strings.TrimLeft(in.Child[0].Data, "\r\n")

	return out
}

//Gets the author from a subnode of the DOM tree [implementation specific]
//For pChan, the subnode is the first <label> tree under the <div class='thread'>
//The <label> tree has a <span> named postername that is the authors name
func (t *Thread) extAuthor(in *html.Node) string {
	out := "DEBUG-default"

	for ix := 0; ix < len(in.Child); ix++ {
		c := in.Child[ix]
		if c.Type == html.ElementNode && c.Data == "a" {
			out = c.Child[0].Data
		} else {
			out = c.Data
		}
	}

	return out
}
