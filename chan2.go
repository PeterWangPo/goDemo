package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(values []int, c chan int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	c <- sum
}
func main() {
	randNum := generateRandomNumber(1, 15, 10)
	c := make(chan int, 10)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	for i := 0; i < 10; i++ {
		go sum(arr[:randNum[i]], c)
	}
	s := [10]int{0}
	for i := 0; i < 10; i++ {
		if v, ok := <-c; ok {
			s[i] = v
		} else {
			s[i] = 9999
		}
	}
	fmt.Println(s)
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// randPerm := r.Perm(15)//生成1-15随机切片,然后可以取前10个
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		// //查重
		// exist := false
		// for _, v := range nums {
		// 	if v == num {
		// 		exist = true
		// 		break
		// 	}
		// }

		// if !exist {
		// 	nums = append(nums, num)
		// }
		nums = append(nums, num)
	}

	return nums
}
