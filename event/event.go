package event

type Event struct {
	UserIds []string
	Payload any
}

type ApplyAddContactEvent struct {
	Event     string
	UserId    string
	UserName  string
	ContactId string
	Message   string
}
func (event *ApplyAddContactEvent)SetDefaultValue() *ApplyAddContactEvent{
	event.Event = "APPLY_ADD_CONTACT_EVENT"
	return event
}

type MessageEvent struct {
	Event string 
	Type string;
	Content string;
	CreaterName string;
}
func (event *MessageEvent)SetDefaultValue() *MessageEvent{
	event.Event = "MESSAGE_EVENT"
	return event
}