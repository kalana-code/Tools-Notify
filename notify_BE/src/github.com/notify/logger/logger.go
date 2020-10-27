package logger

import "log"

func ErrorPrint(component string, message string , err error){
	log.Println("ERRO: ["+component+"]: "+message+" | ",err)
}


func InfoPrint(component string, message string){
	log.Println("INFO: ["+component+"]: "+message)
}