package main

import (
	"fmt"
)

type Runner struct {
	num int64
}

func (run *Runner) SetId(num int64) {
	run.num = num
	fmt.Println(run.num)
}

func (run *Runner) GetId() int64 {
	return run.num
}

func main() {
	var demo interface{} = new(Runner)

	demo.(interface {
		SetId(int64)
	}).SetId(134)
}
