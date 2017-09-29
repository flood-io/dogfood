package dogfood

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"github.com/Pallinder/sillyname-go"
)

type Dog struct {
	Name string
  Body string
  Time int32
}

func Adddog(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusCreated)
}

func Deletedog(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNoContent)
}

func FinddogsByStatus(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusOK)
}

func FinddogsByTags(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusOK)
}

func GetdogById(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusOK)
}

func Updatedog(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusOK)
}

func UpdatedogWithForm(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusOK)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		m := Dog{sillyname.GenerateStupidName(), "Hello", int32(time.Now().Unix())}
		b, err := json.Marshal(m)
		if err != nil {
      panic(err)
    }
		fmt.Fprintf(w, string(b))
		w.WriteHeader(http.StatusOK)
}

