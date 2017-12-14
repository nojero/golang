package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nojero/computer"
	"io/ioutil"
	"net/http"
	"log"
	"strconv"
)

// App variables
var computer_count = 0
var computers []computer.Computer

func parseBody(r *http.Request, t interface{}) error {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	json.Unmarshal([]byte(body), t)

    return nil
}
func createComputer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	type Args struct {
		Stack int `json:"stack"`
	}
	var args Args

    err := parseBody(r, &args)
    if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer r.Body.Close()

	size := args.Stack
	if size <= 0 {
		http.Error(w, "Error: stack value invalid or not present", http.StatusInternalServerError)
		return
	}

	comp := computer.New(100)
	computers = append(computers, comp)
	fmt.Fprintf(w, "%d\n", computer_count)
	computer_count += 1
}

func setAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    strId := ps.ByName("id")

    id, err := strconv.Atoi(strId)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // valid id
	type Args struct {
		Addr int `json:"addr"`
	}
	var args Args

    err = parseBody(r, &args)
    if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer r.Body.Close()
    if id < len(computers) {
        comp := &computers[id]
        comp.SetAddress(args.Addr)
    } else {
		http.Error(w, "Computer not existant", http.StatusInternalServerError)
    }
}

func insert(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    strId := ps.ByName("id")
    ins := ps.ByName("type")

    id, err := strconv.Atoi(strId)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // valid id
	type Args struct {
		Arg int `json:"arg"`
	}
	var args Args

    err = parseBody(r, &args)
    if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    defer r.Body.Close()

    if id < len(computers) {
        comp := &computers[id]
        comp.Insert(ins, args.Arg)
    } else {
		http.Error(w, "Computer not existant", http.StatusInternalServerError)
    }
}

func execute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    strId := ps.ByName("id")

    id, err := strconv.Atoi(strId)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if id < len(computers) {
        comp := &computers[id]
        ret, err := comp.Execute()
        if err != nil {
	        fmt.Fprintf(w, ret)
        } else {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    } else {
		http.Error(w, "Computer not existant", http.StatusInternalServerError)
    }
}

func main() {
	router := httprouter.New()
	router.POST("/v1/computers", createComputer)
    router.PATCH("/v1/computers/:id/stack/pointer", setAddress)
    router.POST("/v1/computers/:id/stack/insert/:type", insert)
    router.POST("/v1/computers/:id/exec", execute)

	log.Fatal(http.ListenAndServe(":8080", router))
}
