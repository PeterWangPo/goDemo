package main

import (
	"fmt"
	"sync"
)

//面试题：
//某工厂有A，B，C三辆厨房工程车，A车上能清洗材料，B车上能加工材料，C车上能装载材料；三辆工程车能边行驶边清洗/加工/装载材料
//每辆车上有三个工人；最初的原始材料有D1，D2，D3三种，每种材料的清洗耗时比例为6:3:3，每种材料的加工/装载耗时皆为1:1:1，每种
//材料的数量一致；材料的处理顺序为：清洗->加工->装载；车辆之间材料进行交互，需要保持比较近的相对距离
//要求：这三辆车需要将处理完的原材料，尽快送达到商家手上，请问如何分配比较好？

//为什么要声明成全局变量?
//可以避免重复申请变量，节省额外的性能和内存交互时间
var (
	wg sync.WaitGroup
	ca = make(chan []int, 1) //A处理完往ca写数据
	cb = make(chan []int, 1) //B从ca读取数据，处理完后写入cb
	cc = make(chan struct{}) //C从cb读取数据，处理完后写入cc, main中等待完成
)

func main() {
	elements := []int{1, 3, 3}
	go A(elements)
	fmt.Println("A done...")
	go B()
	fmt.Println("B done...")
	go C()
	fmt.Println("C done...")
	<-cc
	fmt.Println("all done...")
}

func A(elements []int) {
	//存储切分后的任务
	var tasks = make([][]int, 3)
	//三个工人干活
	for i := 0; i < 3; i++ {
		//用来分配任务
		task := []int{}
		for _, value := range elements {
			task = append(task, value/3)
		}
		wg.Add(1)
		go func(task []int, i int) {
			//完成清洗任务
			tasks[i] = clean(task)
			//tasks[i] = []int{2}
			wg.Done()
		}(task, i)
	}
	wg.Wait()

	//合并回elements

	for index, _ := range elements {
		elements[index] = 0
		for _, task := range tasks {
			elements[index] += task[index]
		}
	}
	fmt.Println("elements A:", elements)
	ca <- elements
	fmt.Println("Groutine A done...")
}

func clean(task []int) []int {
	return task
}

func B() {
	var elements []int
	c := <-ca
	fmt.Printf("Goroutine B Received:%v  type: %T\n", c, c)
	for index, element := range elements {
		wg.Add(1)
		go func(element, index int) {
			//完成加工任务
			elements[index] = cure(element)
			wg.Done()
		}(element, index)
		fmt.Println("Job B ...")
	}
	wg.Wait()
	cb <- elements
	fmt.Println("Groutine B done...")
}

func cure(e int) int {
	return e
}

func C() {
	elements := <-cb
	fmt.Printf("Goroutine C Received:%v  type: %T\n", elements, elements)
	for index, element := range elements {
		wg.Add(1)
		go func(element, index int) {
			//完成装载任务
			elements[index] = carry(element)
			//elements[index] = 3
			wg.Done()
		}(element, index)
	}
	wg.Wait()
	cc <- struct{}{}
	fmt.Println("Groutine C done...")
}

func carry(e int) int {
	return e
}
