package main

import (
	"net/http"
	"time"

	"github.com/zhaowalilangka/go_blog/internal/routers"
)

type User struct {
	Name string
}

type Bin struct {
	User
	Age int
}

func (u *User) GetName() string {
	return u.Name
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	// a := &Bin{
	// 	User{"ddd"}, 3,
	// }
	// fmt.Println(a.GetName())

	// fmt.Println(a)

}
