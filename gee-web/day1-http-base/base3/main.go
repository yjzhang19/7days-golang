package main

// $ curl http://localhost:9999/
// URL.Path = "/"
// $ curl http://localhost:9999/hello
// Header["Accept"] = ["*/*"]
// Header["User-Agent"] = ["curl/7.54.0"]
// curl http://localhost:9999/world
// 404 NOT FOUND: /world

import (
	"fmt"
	"myGee"
	"net/http"
)

func main() {
	//r := myGee.New()
	//r.GET("/", func(w http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//})
	//
	//r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
	//	for k, v := range req.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	//
	//r.Run(":9999")
	r := myGee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {

		//格式化并输出到io.Writer
		fmt.Fprintf(w, "URL.Path=%s\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%s]=%s\n", k, v)
		}
	})
	r.Run(":8888")
}
