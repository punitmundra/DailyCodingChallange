package main

import (
	"Squadcast/EventHandler/publisher"
	p "Squadcast/EventHandler/publisher"
	"context"
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// EventHandler ...
type EventHandler struct {
}

// PublishToWorker ...
func PublishToWorker(req *p.Request) {
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()
	//eventData, _ := json.Marshal(req)

}

// Publish ...this is method is called by api handler to push the data further to the service worker
func (s *EventHandler) Publish(ctx context.Context, req *p.Request) (res *p.Response, err error) {
	log.Printf("Receive message body from client: %s", req.GetName())
	log.Printf("Receive message body from client: %s", req.GetType())

	// we can put below logic to config file
	var subject string
	if req.GetType() == "issue" { // jira
		subject = "incident_created"
	}
	if req.GetType() == "query" { // slack
		subject = "query_created"
	}
	if req.GetType() == "post" { // reddit
		subject = "post_created"
	}
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	defer natsConnection.Close()
	event := p.Request{
		Name:   req.GetName(),
		Type:   req.GetType(),
		Status: req.GetStatus(),
	}
	data, _ := proto.Marshal(&event)
	// Publish message on subject
	natsConnection.Publish(subject, data)
	//natsConnection.
	log.Println("Published message on subject " + subject)
	return &p.Response{Status: "Got the request from API Server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("not able to listen on %s: %v", lis.Addr().String(), err)
	}

	grpcServer := grpc.NewServer()
	publisher.RegisterPublisherServiceServer(grpcServer, &EventHandler{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf(" failed to start grpc server: %v", err)
	}
}
