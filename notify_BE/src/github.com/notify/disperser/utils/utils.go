package utils

// called endpoint service in node base on type of job as following

import (

	"fmt"
	"github.com/notify/disperser/model"
	"log"
)

//Disperse used for disperse jobs in job list
func Disperse(job *model.Job) {
	log.Println("INFO: [RD]: Dispursing JOB is being Initiated")
	switch job.Type {
	case model.SEND_MAIL:
		//Executing a job
		fmt.Println("Kalana")
		break
	case 4:
		break
	}
}
