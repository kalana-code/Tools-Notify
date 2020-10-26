package model

//JobType used for define job type
type JobType int

const (
	//send mail job type is send mail
	SEND_MAIL JobType = 1
)

//Job is used for add Tesk for task queue
type Job struct {
	Type        JobType
	TaskDetails interface{}
}

// send Email jog contex
type SendEmail struct {
	Email string
}
