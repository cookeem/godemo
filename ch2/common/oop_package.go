package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

//匿名组合
type Fish struct {
	TypeName string
	Age      int
}

type Shark struct {
	Fish
	NumOfTooth int
}

func (fish Fish) Swimming() {
	fmt.Println("fish ", fish, "swimming")
}

type Job struct {
	Command     string
	time.Time            //值类型
	*log.Logger          //引用类型
	logFile     *os.File //指定名字的引用类型，包外不可见
}

func NewJob(command string) Job {
	job := Job{}
	job.Command = command
	job.Time = time.Now()
	fileName := "test.log"
	var logFile, err = os.Open(fileName)
	if err != nil {
		logFile, err = os.Create(fileName)
		if err != nil {
			fmt.Println("NewJob error:", err)
			return job
		}
	}
	job.logFile = logFile
	job.Logger = log.New(job.logFile, "[Debug]", log.LstdFlags) //这是一个引用
	return job
}

func (job *Job) Start() {
	job.Time = time.Now()
	job.Logger.SetPrefix("[Debug]")
	fmt.Println("[", job.Time.Format("2006-01-02 15:04:05"), "] start job: ", job.Command)
	job.Logger.Println("[", job.Time.Format("2006-01-02 15:04:05"), "] start job: ", job.Command)
}

//返回结果Job需要改变原Job
func (job *Job) SetCommand(command string) {
	job.Command = command
}

func (job *Job) CloseJob() {
	defer job.logFile.Close()
}
