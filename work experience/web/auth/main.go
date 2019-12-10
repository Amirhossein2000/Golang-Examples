package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var OnlineUsers = make(map[int]User)

type User struct {
	name string
}

func main() {
	http.HandleFunc("/home", isAuth)
	http.HandleFunc("/login", login)

	go http.ListenAndServe("localhost:8080", nil)
	testRequest()

	time.Sleep(time.Second * 10)
}

func isAuth(responseWriter http.ResponseWriter, request *http.Request) {
	session, err := request.Cookie("session")

	if session != nil && err == nil {
		userUUID, err := strconv.Atoi(session.Value)

		if user, ok := OnlineUsers[userUUID]; ok && err == nil {
			_, err = fmt.Fprintln(responseWriter, "<h1>", user.name, "</h1>")
		}

	} else {
		_, err = fmt.Fprintln(responseWriter, "<h1>Bye</h1>")
	}
}

func login(responseWriter http.ResponseWriter, request *http.Request) {
	newUUDI := newUserSession()
	http.SetCookie(responseWriter, &http.Cookie{Name: "session", Value: strconv.Itoa(newUUDI)})
	OnlineUsers[newUUDI] = User{name: "USR"}

	responseWriter.Write([]byte("now you are in"))

	ctx, _ := context.WithTimeout(context.Background(), time.Minute*10)
	go expireLogin(newUUDI, ctx)
}

func newUserSession() int {
	for {
		uuid := rand.Intn(987654321-123456789) + 123456789
		if _, ok := OnlineUsers[uuid]; !ok {
			return uuid
		}
	}

}

func expireLogin(uuid int, ctx context.Context) {
	<-ctx.Done()
	delete(OnlineUsers, uuid)
}

func testRequest() {
	for i := 0; i < 500; i++ {
		go func() {
			r, err := http.Get("http://localhost:8080/login")
			if err != nil {
				fmt.Println()
			} else {
				fmt.Println(r.Cookies())
			}
		}()
	}
}
