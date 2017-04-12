package main

import (
	"cookeem.com/ch2/common"
	"fmt"
	"time"
)

func main() {
	//类的继承
	jokeFish := common.Fish{TypeName: "jokeFish", Age: 5}
	jokeFish.Swimming()

	sharkFish := common.Shark{}
	sharkFish.Fish.TypeName = "shark"
	sharkFish.Fish.Age = 10
	sharkFish.NumOfTooth = 20

	sharkFish.Swimming()
	//含义一样
	sharkFish.Fish.Swimming()

	ptrJob := common.NewJob("printenv")
	fmt.Println("ptrJob:", ptrJob.Command, ptrJob.Time)
	ptrJob.Start()
	time.Sleep(2 * 1e9)
	ptrJob.SetCommand("tail -f /etc/hosts")
	ptrJob.Start()

	ptrJob.CloseJob()
}
