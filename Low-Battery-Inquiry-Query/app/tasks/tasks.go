package tasks

import (
	"Low-Battery-Inquiry-Query/app/models"
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	LowBatteryRemind = "LowBatteryRemind"
)

func NewLowBatteryRemindTask(userId int, userOpenID, value, address, remark string) (*asynq.Task, error) {
	payload, err := json.Marshal(models.LowBatteryRemindPayload{
		UserId:     userId,
		UserOpenID: userOpenID,
		Value:      value,
		Address:    address,
		Remark:     remark})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(LowBatteryRemind, payload), nil
}
