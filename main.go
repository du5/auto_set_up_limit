package main

import (
	"log"
	"time"
)

func main() {
	Log_ENV()
	tc5s := time.NewTicker(5 * time.Second)
	tc24h := time.NewTicker(24 * time.Hour)
	for {
		<-tc5s.C
		for _, host := range env.Host {
			for k, v := range Get_List(host) {
				up_limit, hostname := Get_Limit(v.Tracker)
				if v.Up_Limit <= 0 || v.Up_Limit > up_limit {
					if err := Set_Limit(host, k, up_limit); err != nil {
						log.Printf("[%s] 种子: %s 限速更新失败, %s", hostname, k, err.Error())
					} else {
						log.Printf("[%s] 种子: %s 限速更新 %.2fMiB/s 成功!", hostname, k, float64(up_limit)/1024/1024)
					}
				}
			}
		}
		select {
		case <-tc24h.C:
			for _, host := range env.Host {
				Update_TK(host)
			}
		default:
		}
	}
}
