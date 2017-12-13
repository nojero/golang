package main

import (
	//"github.com/nojero/computer"
	"fmt";
	"encoding/json";
    "io/ioutil";
	"net/http"
)

func main() {
	http.HandleFunc("/v1/computers", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

        defer r.Body.Close()

        type Args struct {
            Stack int `json:"stack"`
        }

        var args Args
        err = json.Unmarshal([]byte(body), &args)
        if err != nil {
            fmt.Println("error:", err)
        }
        size := args.Stack
        if size <= 0 {
            http.Error(w, "Error: stack value invalid or not present", http.StatusInternalServerError)
        }
	})

	http.ListenAndServe(":8080", nil)

	/*
	   comp := computer.New(100)
	   comp.SetAddress(50)
	   comp.Insert("MULT", 0)
	   comp.Insert("PRINT", 0)
	   comp.Insert("RET", 0)
	   comp.SetAddress(0)
	   comp.Insert("PUSH", 1009)
	   comp.Insert("PRINT", 0)
	   comp.Insert("PUSH", 6)
	   comp.Insert("PUSH", 101)
	   comp.Insert("PUSH", 10)
	   comp.Insert("CALL", 50)
	   comp.Insert("STOP", 0)
	   comp.SetAddress(0)
	   comp.Execute()
	*/
}
