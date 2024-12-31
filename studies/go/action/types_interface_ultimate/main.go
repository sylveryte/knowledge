package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (ds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	userName, ok := ds.userData[userId]
	return userName, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Ash",
			"2": "Resh",
			"3": "San",
		},
	}
}

//generic interfaces

type Logger interface {
	Log(message string)
}

type DataStore interface {
	UserNameForId(userId string) (string, bool)
}

type LogAdapter func(message string)

func (lg LogAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in SayHellow for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown error")
	}
	return "Hello, " + name, nil
}
func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown error")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds SimpleDataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}


type Logic interface {
  SayHello(userId string) (string, error)
}

type Controller struct{
  l Logger
  logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request){
  c.l.Log("In SayHello Controller")
  userId := r.URL.Query().Get("user_id")
  message, err := c.logic.SayHello(userId)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }
  w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
  return Controller{
    l: l,
    logic: logic,
  }
}

func main() {
  l := LogAdapter(LogOutput)
  ds := NewSimpleDataStore();

  logic := NewSimpleLogic(l,ds)
  c := NewController(l,logic)

  http.HandleFunc("/hello",c.SayHello)
  http.ListenAndServe(":8080",nil)


}
