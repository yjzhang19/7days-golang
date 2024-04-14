/**
* @Author: zhangyujin
* @Description:
* @Flie:gee
* @Date: 2024/4/14 14:10
 */
package myGee

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	route map[string]HandleFunc
}

func New() *Engine {
	return &Engine{route: make(map[string]HandleFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	log.Printf("route %4s - %s", method, pattern)
	engine.route[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}
func (engine *Engine) POST(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.route[key]; ok {
		handler(w, req)
	} else {
		fmt.Println(w, "404 NOT FOUND:%s", req.URL)
	}
}
