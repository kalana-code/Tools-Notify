package db

import (
	"container/list"
	"errors"
	"github.com/notify/disperser/model"
	"github.com/notify/disperser/utils"
	"log"
	"sync"
)

//JobQueue is hold current jobs
type JobQueue struct {
	List *list.List
}

var instance JobQueue

var totalJobCount int

var once sync.Once

//GetJobQueue Initiating list database
func GetJobsQueue() *JobQueue {
	once.Do(func() {
		instance.List = list.New()
		totalJobCount = 0
	})
	return &instance
}

//AddJob  is used for add JOBs
func (obj *JobQueue) AddJob(Job model.Job) error {
	if instance.List != nil {
		log.Println("INFO: [RD]: Add Job")
		instance.List.PushBack(Job)
		totalJobCount++
		return nil
	}
	return errors.New("No Data Base Initiate")

}

//Disperse  is used for Disperse JOBs
func (obj *JobQueue) Disperse() error {
	if instance.List != nil {
		if instance.List.Back() != nil {
			currentJob := instance.List.Front().Value
			cJob, valid := currentJob.(model.Job)
			instance.List.Remove(instance.List.Front())
			if !valid {
				return errors.New("Invalid JOB ")
			}
			utils.Disperse(&cJob)
		}
		return nil
	}
	return errors.New("No Data Base Initiate")

}
