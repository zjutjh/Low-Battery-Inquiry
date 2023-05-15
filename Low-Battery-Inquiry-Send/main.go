package main

import (
	"Low-Battery-Inquiry-Send/app/tasks"
	"Low-Battery-Inquiry-Send/config/database"
	"Low-Battery-Inquiry-Send/config/redis"
	"github.com/hibiken/asynq"
	"log"
)

func main() {
	database.Init()
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redis.RedisInfo.Host + ":" + redis.RedisInfo.Port},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.LowBatteryRemind, tasks.HandleBatteryRemindTask)
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
