package models

type LowBatteryRemindPayload struct {
	UserId     int
	UserOpenID string
	Value      string
	Address    string
	Remark     string
}
