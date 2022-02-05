package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

type ENV struct {
	Host         []string
	Tracker      map[string]int64
	Default      int64
	TrackersList string
}

func Get_ENV() ENV {
	var env ENV
	jf, _ := os.Open("env.json")
	defer jf.Close()
	byteValue, _ := ioutil.ReadAll(jf)
	_ = json.Unmarshal([]byte(byteValue), &env)
	return env
}

func Get_Limit(tracker string) (int64, string) {
	uri, _ := url.Parse(tracker)
	for k, v := range env.Tracker {
		if strings.Contains(strings.ToLower(uri.Hostname()), k) {
			return v * 1024, uri.Hostname()
		}
	}
	return env.Default * 1024, uri.Hostname()
}

func Log_ENV() {
	log.Printf("站点列表: [%s]", strings.Join(env.Host, ", "))
	limits := []string{}
	for k, v := range env.Tracker {
		limits = append(limits, fmt.Sprintf("%s: %.2fMiB/s", k, float64(v)/1024))
	}
	log.Printf("限速规则: [%s]", strings.Join(limits, ", "))
	default_limit := fmt.Sprintf("%.2fMiB/s", float64(env.Default)/1024)
	if env.Default == 0 {
		default_limit = "不限速"
	}
	log.Printf("默认限速: [%s]", default_limit)
	if env.TrackersList != "" {
		log.Printf("自动更新: [%s]", env.TrackersList)
	}
}
