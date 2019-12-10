package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Middleware func(w http.ResponseWriter, r *http.Request) error

var RouteMap = make(map[string]http.HandlerFunc)
var MiddlewareMap = make(map[string]Middleware)
var NoAuthNeedMap = make(map[string]bool)

func init() {
	RouteMap["/"] = HomeHandler

	// user
	MiddlewareMap["user"] = userMiddleware
	RouteMap["/user"] = UserHandler
	RouteMap["/user/get_all"] = UserGetAllHandler

	// book
	NoAuthNeedMap["book"] = true
	MiddlewareMap["book"] = bookMiddleware
	RouteMap["/book"] = BookHandler
	RouteMap["/book/get_all"] = bookGetAllHandler
}

func main() {
	http.HandleFunc("/", MainHandler)
	log.Println(http.ListenAndServe("localhost:8080", nil))
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	err := runMiddlewares(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	handlerFunc, ok := RouteMap[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
	} else {
		handlerFunc(w, r)
	}
}

func runMiddlewares(w http.ResponseWriter, r *http.Request) error {
	groups := getGroups(r.URL.Path)
	auth := false

	for _, group := range groups {
		middleware := getMiddleware(group)
		if middleware != nil {
			if NeedAuth(group) && !auth {
				err := Auth(w, r)
				if err != nil {
					return err
				}
				auth = true
			}
			err := middleware(w, r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getGroups(url string) []string {
	return strings.Split(url, "/")
}

func getMiddleware(group string) Middleware {
	handlerFunc, ok := MiddlewareMap[group]

	if !ok {
		return nil
	} else {
		return handlerFunc
	}
}

func Auth(w http.ResponseWriter, r *http.Request) error {
	response := &http.Response{
		StatusCode: http.StatusNetworkAuthenticationRequired,
		Status:     "Nooooooooooo Way",
	}

	response.Write(w)

	return fmt.Errorf("Access Denid %v", r.URL.Host)
}

func NeedAuth(key string) bool {
	if _, ok := NoAuthNeedMap[key]; ok {
		return false
	}
	return true
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func UserGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func bookGetAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func bookMiddleware(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(r.URL.Path + "bookMiddleware\n"))
	return nil
}

func userMiddleware(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte(r.URL.Path + "userMiddleware\n"))
	return nil
}
