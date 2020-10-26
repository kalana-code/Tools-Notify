package services


type Event struct{
	ApplicationId string
	//following this used for event handle
	EventName  string
	EventData  interface{}
}
