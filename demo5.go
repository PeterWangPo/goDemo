package main

import (
	"fmt"
)

type PluginInfo struct {
	Id   int
	Name string
}

func main() {
	info := make([]*PluginInfo, 0)

	pluginMap := make(map[int]*PluginInfo)
	pluginMap[1] = &PluginInfo{
		Id:   1,
		Name: "aa",
	}
	pluginMap[2] = &PluginInfo{
		Id:   2,
		Name: "bb",
	}

	for i, tmp := range pluginMap {
		fmt.Println(i, *tmp)     //打印出得数据没有问题
		info = append(info, tmp) //info中的地址数据竟然是相同的
	}
	fmt.Println(info)

	for _, tmp := range info {
		fmt.Println(*tmp)
	}

}
