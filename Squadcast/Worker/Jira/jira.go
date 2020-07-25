package main

import (
	p "Squadcast/EventHandler/publisher"
	"Squadcast/model"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/nats-io/nats"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
)

// JiraWorker ... it is jira worker
type JiraWorker struct {
	connection *nats.Conn
	topic      string
}

const topic = "incident_created"

/*
func (worker *JiraWorker) actionOnEvent() {
	if len(worker.queue) > 0 {
		event := worker.queue[0]
		fmt.Println("Processing this id:", event.Name)
		incident := model.GetIncidentByID(cast.ToInt64(event.Name))

		fmt.Println(time.Now(), "-", incident.CreatedAt, "=", time.Now().Sub(incident.CreatedAt))
		if incident.Acknowlege == "true" {
			worker.queue = worker.queue[1:]
		} else {
			fmt.Println(" I am going to riase the jira ticket for this Id:", event.Name)

		}
	}
}
*/
const threshold = 10

func getTickerValue(incident *model.Incident) int {
	expiry := time.Now().Add(time.Minute * threshold)
	at := incident.CreatedAt
	diff := expiry.Sub(at)
	fmt.Println(diff.Minutes())
	return cast.ToInt(diff)
}

/*
type DataEvent struct {
	Data  interface{}
	Topic string
}

// DataChannel ...
type DataChannel chan DataEvent

type EventBus struct {
	subscribers map[int]DataEvent
	rm          sync.RWMutex
}
*/

// BufferQueue ...
func BufferQueue() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})
	out := make(chan interface{})
	go func() {
		var queue []interface{}
	loop:
		for {
			select {
			case v, ok := <-in:
				if !ok {
					break loop
				} else {
					queue = append(queue, v)
				}

			}
		}
		//close(out)
	}()
	return in, out
}

// DataRequest ...
type DataRequest struct {
	Data p.Request
	Time time.Time
}

// EventQueue ...
type EventQueue struct {
	queue []DataRequest
	rm    sync.RWMutex
}

// EnQueue ...
func (q *EventQueue) EnQueue(data *DataRequest) {
	q.rm.RLock()
	q.queue = append(q.queue, *data)
	q.rm.RUnlock()
}

// DeQueue ...
func (q *EventQueue) DeQueue() (data *DataRequest) {
	q.rm.Lock()
	if len(q.queue) > 0 {
		data = &q.queue[0]
		q.queue = q.queue[1:]
	}
	q.rm.Unlock()
	return
}

// Peek ...
func (q *EventQueue) Peek() (data *DataRequest) {
	q.rm.RLock()
	if len(q.queue) > 0 {
		data = &q.queue[0]
	}
	q.rm.RUnlock()
	return data
}

var tube = &EventQueue{}

func main() {
	var worker JiraWorker
	var err error
	worker.connection, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}

	worker.topic = topic
	worker.connection.Subscribe(worker.topic, func(msg *nats.Msg) {
		eventStore := p.Request{}
		err := proto.Unmarshal(msg.Data, &eventStore)

		if err == nil {
			log.Println("Received message in jira service id:", eventStore.Name, " type:", eventStore.Type)
			incident := model.GetIncidentByID(cast.ToInt64(eventStore.Name))

			if incident.Acknowlege == "false" {
				// push into queue
				data := DataRequest{Data: eventStore, Time: incident.CreatedAt.Add(time.Second * threshold)}
				fmt.Println("job is enqueued")
				tube.EnQueue(&data)
			}
			go func() {
				for now := range time.Tick(time.Second * 5) {
					//fmt.Println(now)
					temp := tube.Peek()
					if temp != nil {
						fmt.Println("..................................", now)
						fmt.Println(temp.Data, ":", temp.Time, ":", time.Now())
						if time.Now().After(temp.Time) {
							// we will check if ticket is ack or not
							fmt.Println("Jira Ticket Create for the ID:", temp.Data.Name)
							tube.DeQueue()
						}
					}
				}
			}()
		}
	})
	runtime.Goexit()
}
