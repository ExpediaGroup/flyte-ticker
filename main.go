/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	api "github.com/HotelsDotCom/flyte-client/client"
	"github.com/HotelsDotCom/flyte-client/config"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/go-logger"
	"net/url"
	"time"
)

var tickEventDef = flyte.EventDef{Name: "Tick"}

func main() {
	conf := config.FromEnvironment()
	pack := flyte.NewPack(getPackDef(), api.NewClient(conf.FlyteApiUrl, conf.Timeout))
	pack.Start()

	ticker := time.NewTicker(1 * time.Minute)
	sendTickEvents(ticker, pack)
}

func getPackDef() flyte.PackDef {
	hu, _ := url.Parse("https://github.com/ExpediaGroup/flyte-ticker/blob/master/README.md")
	return flyte.PackDef{
		Name:      "Ticker",
		HelpURL:   hu,
		EventDefs: []flyte.EventDef{tickEventDef},
	}
}

func sendTickEvents(ticker *time.Ticker, pack flyte.Pack) {
	for tick := range ticker.C {
		te := toTickEvent(tick)
		pack.SendEvent(te)
		logger.Debugf("sent tick event: %+v", te)
	}
}

type tickEvent struct {
	Time time.Time `json:"time"`
}

func toTickEvent(t time.Time) flyte.Event {
	return flyte.Event{
		EventDef: tickEventDef,
		Payload:  tickEvent{t},
	}
}
