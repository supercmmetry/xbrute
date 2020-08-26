package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"xbrute/pkg/task"
	"xbrute/utils"
)

func addTask(taskSvc *task.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newTask := task.Task{}

		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			utils.ErrWrap(w, err.Error())
			return
		}

		taskSvc.AddTask(&newTask)
		utils.RespWrap(w, http.StatusOK, "Added task successfully")
	}
}

func getTasks(taskSvc *task.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks := taskSvc.GetTasks()

		respTasks := make([]map[string]interface{}, 0)

		for _, t := range tasks {
			rt := map[string]interface{} {
				"id": t.Id,
				"attack_count": t.AttackCount,
				"algorithm": t.AlgorithmUsed,
				"payload": map[string]interface{} {
					"start": utils.Bytes2IntSlice(t.PayloadData.Start),
					"count": t.PayloadData.Count,
					"prefix": utils.Bytes2IntSlice(t.PayloadData.Prefix),
					"alphabet": utils.Bytes2IntSlice(t.PayloadData.Alphabet),
				},
				"target": utils.Bytes2IntSlice(t.Target),
				"partial_data": utils.Bytes2IntSlice(t.PartialData),
			}

			respTasks = append(respTasks, rt)
		}

		utils.Wrap(w, map[string]interface{}{"tasks": respTasks})
	}
}

func executeLocalTask(taskSvc *task.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newTask := task.Task{}

		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			utils.ErrWrap(w, err.Error())
			return
		}

		result := taskSvc.ExecuteTask(&newTask)
		if result == nil {
			utils.RespWrap(w, 404, "No key was found in the given payload")
			return
		}

		outputSlice := make([]uint16, 0)
		for _, v := range result.Output {
			outputSlice = append(outputSlice, uint16(v))
		}

		utils.Wrap(w, map[string]interface{}{"id": result.Id, "output": outputSlice})
	}
}

func executeTask(taskSvc *task.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newTask := task.Task{}

		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			utils.ErrWrap(w, err.Error())
			return
		}

		result := taskSvc.BruteForce(&newTask)
		if result == nil {
			utils.RespWrap(w, 404, "No key was found in the given payload")
			return
		}

		outputSlice := make([]uint16, 0)
		for _, v := range result.Output {
			outputSlice = append(outputSlice, uint16(v))
		}

		utils.Wrap(w, map[string]interface{}{"id": result.Id, "output": outputSlice})
	}
}

func feedResult(taskSvc *task.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := task.Result{}
		err := json.NewDecoder(r.Body).Decode(&result)
		if err != nil {
			utils.ErrWrap(w, err.Error())
			return
		}

		if !taskSvc.VerifyResult(result) {
			utils.RespWrap(w, http.StatusConflict, "Result validation failed")
			return
		}

		taskSvc.SetSolution(result.Output)
		utils.RespWrap(w, http.StatusOK, "Result validated successfully")
	}
}

func MakeTaskHandlers(router *httprouter.Router, taskSvc *task.Service) {
	router.HandlerFunc("POST", "/api/v1/tasks/add", addTask(taskSvc))
	router.HandlerFunc("GET", "/api/v1/tasks/all", getTasks(taskSvc))
	router.HandlerFunc("POST", "/api/v1/tasks/execute-local", executeLocalTask(taskSvc))
	router.HandlerFunc("POST", "/api/v1/tasks/execute", executeTask(taskSvc))
	router.HandlerFunc("POST", "/api/v1/tasks/feed", feedResult(taskSvc))
}

