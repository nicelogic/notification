package event

type ApplyAddContactEvent struct {
	Event     string
	UserId    string
	UserName  string
	ContactId string
}

func NewApplyAddContactEvent(userId string, userName string, contactId string) *ApplyAddContactEvent {
	return &ApplyAddContactEvent{
		Event:     "APPLY_ADD_CONTACT_EVENT",
		UserId:    userId,
		UserName:  userName,
		ContactId: contactId,
	}
}
