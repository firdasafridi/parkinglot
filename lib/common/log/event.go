package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Event struct {
	Type    string                 `json:"type,omitempty"`
	Time    time.Time              `json:"timestamp,omitempty"`
	Tag     map[string]interface{} `json:"fields,omitempty"`
	Message string                 `json:"msg,omitempty"`
	config  *Config                `json:"-"`
}

// NewEvent will create new structure event that created
func NewEvent(typeValue string) *Event {
	return &Event{Type: typeValue, config: defaultLog}
}

func (event *Event) Timestamp() *Event {
	if event == nil {
		return event
	}
	event.Time = time.Now()
	return event
}

func (event *Event) Msgf(format string, v ...interface{}) {
	if event == nil {
		return
	}
	event.Message = fmt.Sprintf(format, v...)
	event.run()
}

func (event *Event) Msg(msg string) {
	if event == nil {
		return
	}
	event.Message = msg
	event.run()
}

func (event *Event) MsgFatalln(msg string) {
	if event == nil {
		return
	}
	event.Message = msg
	event.runFatalln()
}

func (event *Event) MsgFatalf(format string, v ...interface{}) {
	if event == nil {
		return
	}
	event.Message = fmt.Sprintf(format, v...)
	event.runFatalln()
}

func (event *Event) Fields(tag map[string]interface{}) *Event {
	if event == nil {
		return event
	}
	event.Tag = tag
	return event
}

func (event *Event) run() {
	if !event.config.IsJson {
		log.Println(event)
		return
	}

	eventb, _ := json.Marshal(event)
	fmt.Printf("%s\n", eventb)

	f, err := os.OpenFile("parking_lot.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%v\n", event)); err != nil {
		fmt.Printf("%s\n", err)
	}
}
func (event *Event) runFatalln() {
	if !event.config.IsJson {
		log.Println(event)
		return
	}
	f, err := os.OpenFile("parking_lot.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%v\n", event)); err != nil {
		fmt.Printf("%s\n", err)
	}
	eventb, _ := json.Marshal(event)
	panic(fmt.Sprintf("%s", eventb))
}
