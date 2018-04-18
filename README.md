# flyte-ticker

![Build Status](https://travis-ci.org/HotelsDotCom/flyte-ticker.svg?branch=master)

A simple ticker pack for Flyte. Emits a "Tick" event every minute which can be
used in flows for cron like behaviour. The event looks like this:

    {
        "time": "2018-02-14T17:14:37.765525Z"
    }

That is, the time is in the [ISO-8601 format](https://en.wikipedia.org/wiki/ISO_8601).

## Build

Pack requires go version 1.9+ and uses dep to manage dependencies (install dep
and run `dep ensure` before build/test).

- to build: `go build`
- to test: `go test`
- docker build `docker build -t <name>:<version> .`

Or, if you have make installed, you can use that. See the [Makefile](Makefile)
for available targets.

## Configuration

The plugin is configured using environment variables:

ENV VAR                          | Default  |  Description                               | Example               
 ------------------------------- |  ------- |  ----------------------------------------- |  ---------------------
`FLYTE_API`                      | -        | The API endpoint to use                    | http://localhost:8080
`FLYTE_API_TIMEOUT`              | 10 secs  | Combined timeout for accessing the API     | 20
`FLYTE_LABELS`                   | -        | Labels to disambiguate this instance       | env=staging,bar=foo


## Example Flow

This flow sends a friendly greeting at 9am every morning!:
```json
{
  "name": "tick_demo",
  "description": "Demo echoing tick events to Slack room",
  "steps": [
    {
      "id": "get_tick",
      "event": {
        "packName": "Ticker",
        "name": "Tick"
      },
      "criteria": "{{ Event.Payload.time | match: 'T09:' }}",
      "context": {
        "Time": "{{ Event.Payload.time }}"
      },
      "command": {
        "packName": "Slack",
        "name": "SendMessage",
        "input": {
          "channelId":"<SLACK CHANNEL ID>",
          "message":"Good morning everyone!"
        }
      }
    }
  ]
}
```
