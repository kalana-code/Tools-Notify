package subscriber

const (
	component_name = "SC" // subscribe service component
)

type SubscribeRequest struct {
	Name       string
	Email      string
	LevelCount int
}

type ApplicationContext struct {
	ApplicationID   string
	ApplicationName string
}

type Subscriber struct {
	Id         int    `json:"subscriber_id"`
	Email      string `json:"subscriber_email"`
	LevelCount int    `json:"subscriber_level_count"`
}
