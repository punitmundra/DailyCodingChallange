package main

import (
	p "Squadcast/EventHandler/publisher"
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats"
	"google.golang.org/protobuf/proto"
)

// RedditWorker ... it is reddit worker
type RedditWorker struct {
	connection *nats.Conn
	topic      string
}

const topic = "post_created"

func main() {

	var worker RedditWorker
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
			// Handle the message
			//log.Println("Received message in service:")
			log.Println("Received message in reddit service id:", eventStore.Name, " type:", eventStore.Type)
		}
	})
	runtime.Goexit()
}
