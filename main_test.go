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
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPackDef(t *testing.T) {
	pd := getPackDef()
	assert.Equal(t, "Ticker", pd.Name)
	assert.Equal(t, 0, len(pd.Labels))
	assert.Equal(t, []flyte.EventDef{{Name: "Tick"}}, pd.EventDefs)
	assert.Equal(t, 0, len(pd.Commands))
	assert.Equal(t, "https://github.com/ExpediaGroup/flyte-ticker/blob/master/README.md", pd.HelpURL.String())
}

func TestShouldSendTickEventsOnTickEvents(t *testing.T) {
	ch := make(chan time.Time)

	now := time.Now()

	var receivedEvents []flyte.Event

	pack := MockPack{}
	pack.sendEvent = func(event flyte.Event) error {
		receivedEvents = append(receivedEvents, event)

		assert.Equal(t, "Tick", event.EventDef.Name)
		assert.Nil(t, event.EventDef.HelpURL)
		assert.Equal(t, tickEvent{Time: now}, event.Payload)

		return nil
	}

	ticker := &time.Ticker{C: ch}
	go sendTickEvents(ticker, pack)

	ch <- now

	require.NotEmpty(t, receivedEvents)
	assert.Contains(t, receivedEvents, flyte.Event{EventDef: flyte.EventDef{Name: "Tick"}, Payload: tickEvent{Time: now}})
}

type MockPack struct {
	sendEvent func(flyte.Event) error
}

func (p MockPack) Start() {}

func (p MockPack) SendEvent(event flyte.Event) error {
	return p.sendEvent(event)
}
