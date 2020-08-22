package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"xbrute/api/handlers"
	"xbrute/pkg/task"
)

func main() {
	router := httprouter.New()
	taskSvc := task.NewTaskService()

	handlers.MakeTaskHandlers(router, &taskSvc)

	fmt.Println("Listening on port: 4444")

	err := http.ListenAndServe(fmt.Sprintf(":%s", "4444"), router)
	if err != nil {
		log.Fatal(err)
		return
	}

	//newTask := task.Task{
	//	Id: 0,
	//	PayloadData: task.Payload{
	//		Prefix: []byte("bbabbcbcbacacababacac"),
	//		Start:    []byte("aabbbabaaac"),
	//		Count:    100000000,
	//		Alphabet: []byte{'a', 'b', 'c', 'd'},
	//	},
	//	Target:        []byte{30, 226, 236, 2, 153, 130, 231, 161, 210, 80, 163, 80, 125, 33, 225, 48, 183, 106, 240, 70, 116, 163, 229, 251, 242, 139, 65, 95, 172, 181, 157, 102, 250, 58, 201, 252, 242, 211, 101, 133, 144, 227, 34, 37, 106, 3, 255, 37, 114, 149, 127, 199, 44, 97},
	//	PartialData:   []byte("Super"),
	//	AlgorithmUsed: task.AES,
	//}
	//
	//fmt.Println(s.ExecuteTask(newTask))
}
