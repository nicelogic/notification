package contactevent

type Event struct {
	Event     string
	UserId    string
	UserName  string
	ContactId string
}

func NewEvent(userId string, userName string, contactId string) *Event {
	return &Event{
		Event:     "EVENT_APPLY_ADD_CONTACT",
		UserId:    userId,
		UserName:  userName,
		ContactId: contactId,
	}
}
