package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Json go의 형식을 맞추는 것이 중요!
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	// Io Reader Interface를 포함하고 있다.
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request : ", err)
		return
	}
	user.CreatedAt = time.Now()
	// Interface를 받아서 Json형식으로 엔코딩
	data, _ := json.Marshal(user)

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(data))
}

func BarHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	// Router instance 생성
	mux := http.NewServeMux()

	// index page에 대한 request가 왔을때 Hello World를 반환해라
	// Handler 등록 방법
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", BarHandler)
	mux.Handle("/foo", &fooHandler{})

	return mux
}
