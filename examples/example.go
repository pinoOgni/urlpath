package main

import "github.com/pinoOgni/urlpath"
import "fmt"

var pathTempalte = urlpath.New("/polycube/v1/:ServiceName/")

var provaUrlTemplate string = "/polycube/v1/:ServiceName/"
type key struct {
	url  string
	verb string
}

var methods_map = make(map[key]func(string, string) string)

func main() {
  provaVerb := "GET"
  url_request := "/polycube/v1/helloworld/h1/"
  RegisterHandler("/polycube/v1/:ServiceName/","GET",first)
  RegisterHandler("/polycube/v1/:ServiceName/:CubeName/","GET",second)
  for key, method := range methods_map {
	  match, ok := urlpath.MatchPath(key.url,url_request); 
	  if ok == true && provaVerb == key.verb {
		fmt.Println("the url is present in the method_map ")
		fmt.Println(method(url_request,key.url))
		fmt.Println(match.Params["ServiceName"])
		fmt.Println(match.Params["CubeName"])
		for k,v := range match.Params {
			fmt.Println("k ", k, " --- v ", v)
		}
	  } else {
		fmt.Println("the url is NOT present in the method_map ")
		fmt.Println(method(url_request,key.url))
		fmt.Println(match.Params["ServiceName"])
		fmt.Println(match.Params["CubeName"])
		for k,v := range match.Params {
			fmt.Println("k ", k, " --- v ", v)
		}
	  }
  }

  // Output:
  //
  // foo
  // bar
  //fmt.Println(match.Params["ServiceName"])
}

func RegisterHandler(url string, verb string, handler func(string, string) string) {
	// m[key{url: template, verb: "GET"}] = first
	methods_map[key{url: url, verb: verb}] = handler
}



func first(a string, b string) string {
	return a + " --first-- " + b
}

func second(a string, b string) string {
	return a + " --second-- " + b
}


/*

match, ok := urlpath.MatchPath(provaUrlTemplate,url_request)
  if !ok {
	  fmt.Println("the url is not present in the method_map ")
	  fmt.Println("URL request ", url_request)
	  fmt.Println("URL template ", provaUrlTemplate)
    // handle the input not being valid
    return
  } else {
	fmt.Println("the url is present in the method_map ")
	fmt.Println("URL request ", url_request)
	fmt.Println("URL template ", provaUrlTemplate)
  }

  */