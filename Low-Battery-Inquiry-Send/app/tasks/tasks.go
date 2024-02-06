package tasks

import (
	"Low-Battery-Inquiry-Send/app/config"
	"Low-Battery-Inquiry-Send/app/models"
	"Low-Battery-Inquiry-Send/app/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

const (
	LowBatteryRemind = "LowBatteryRemind"
)

func HandleBatteryRemindTask(ctx context.Context, t *asynq.Task) error {
	var p models.LowBatteryRemindPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	if !config.CheckAccessTokenKey() {
		accessToken, err := service.GetAccessToken()
		if err != nil {
			log.Println("query access token error")
			log.Println(err)
			return err
		}
		config.SetAccessTokenKey(accessToken.AccessToken, accessToken.ExpiresIn)
	}
	key, err := config.GetAccessTokenKey()
	if err != nil {
		log.Println("redis error")
		log.Println(err)
		return err
	}
	err = service.SendInquiry(p, key)
	if err != nil {
		log.Println("sendPart message error")
		log.Println(err)
		return err
	}
	record, err := service.QueryRecordByUserId(p.UserId)
	if err != nil {
		log.Println("query record error")
		log.Println(err)
		return err
	}
	err = service.DeleteRecord(record.ID)
	if err != nil {
		log.Println("delete record error")
		log.Println(err)
		return err
	}
	return nil
}
