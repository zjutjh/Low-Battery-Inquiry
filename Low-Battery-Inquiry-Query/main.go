package main

import (
	"Low-Battery-Inquiry-Query/app/service"
	"Low-Battery-Inquiry-Query/app/tasks"
	"Low-Battery-Inquiry-Query/config/database"
	"Low-Battery-Inquiry-Query/config/redis"
	"github.com/hibiken/asynq"
	"log"
	"strconv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	database.Init()
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redis.RedisInfo.Host + ":" + redis.RedisInfo.Port})
	defer client.Close()
	users, err := service.QueryRecords()
	if err != nil {
		log.Println("query users error")
		log.Println(err)
		return
	}
	for _, user := range users {
		token, err := service.Auth(user.YxyUid)
		if err != nil {
			err := service.DeleteRecord(user.ID)
			if err != nil {
				log.Println("delete record error")
				return
			}
			log.Println("auth yxy error")
			log.Println(err)
			continue
		}
		electricityBalance, err := service.ElectricityBalance(*token)
		if err != nil {
			err := service.DeleteRecord(user.ID)
			if err != nil {
				log.Println("delete record error")
				return
			}
			log.Println("query electricity balance error")
			log.Println(err)
			continue
		}
		if electricityBalance.Soc < 200.0 {
			task, err := tasks.NewLowBatteryRemindTask(user.ID, user.WechatOpenID,
				strconv.FormatFloat(electricityBalance.Soc, 'f', 2, 64),
				electricityBalance.DisplayRoomName,
				"请尽快充值")
			if err != nil {
				log.Println("new task error")
				log.Println(err)
				continue
			}
			info, err := client.Enqueue(task)
			if err != nil {
				log.Println("enqueue task error")
				log.Println(err)
				continue
			}
			log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
		}
	}
}
