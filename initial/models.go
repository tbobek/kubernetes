package main

type Job struct {
	Id string
	Request
}

type Request struct {
	Id         string
	Calls      int `json:"calls" binding:"required"`      // number of calls
	Iterations int `json:"iterations" binding:"required"` // number of iterations
	WaitTime   int `json:"wait_time"`                     // number of milliseconds to wait
}

type Response struct {
	Id       string
	Result   []int
	Status   string
	Error    bool
	ErrorMsg string
	Input    string
	Duration string
}

type Result struct {
	Result int
}
