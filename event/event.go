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

func NewApplyAddContactEvent(userId string,
	userName string,
	contactId string,
	message string) *ApplyAddContactEvent {
	return &ApplyAddContactEvent{
		Event:     "APPLY_ADD_CONTACT_EVENT",
		UserId:    userId,
		UserName:  userName,
		ContactId: contactId,
		Message:   message,
	}
}
