package handler

import (
	"Squadcast/EventHandler/publisher"
	"Squadcast/model"
	u "Squadcast/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cast"
	"google.golang.org/grpc"
)

// RequestToEventHandler ...
func RequestToEventHandler(incident *model.Incident) {
	// GRPC server
	conn, err := grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := publisher.NewPublisherServiceClient(conn)

	res, err := c.Publish(context.Background(), &publisher.Request{
		Name:   cast.ToString(incident.ID),
		Type:   incident.Type,
		Status: incident.Status,
	})
	fmt.Println("call from api handler:", res)
}

// HandledCreateRequest ... handle the new requext
var HandledCreateRequest = func(w http.ResponseWriter, r *http.Request) {

	f := &model.Incident{}
	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		fmt.Println("error:", err, " acc:", f)
		u.Respond(w, u.Message(false, "Invalid request", 400))
		return
	}
	model.InsertIncident(f)
	resp := u.Message(true, "success", 200)
	incident := model.GetLastIncident()

	resp["data"] = "please store the incident Id:" + cast.ToString(incident.ID)
	RequestToEventHandler(incident)
	u.Respond(w, resp)
}

// TicketID ...
type TicketID struct {
	ID int ` json:"id"`
}

// HandledAckRequest ...
var HandledAckRequest = func(w http.ResponseWriter, r *http.Request) {
	f := &TicketID{}
	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		fmt.Println("error:", err, " acc:", f)
		u.Respond(w, u.Message(false, "Invalid request", 400))
		return
	}
	incident := model.GetIncidentByID(cast.ToInt64(f.ID))
	if incident == nil {
		resp := u.Message(true, "success", 400)
		resp["data"] = "invalid request"
		return
	}

	incident.Acknowlege = "true"
	model.UpdateIncidentByID(incident, cast.ToInt64(f.ID))
	resp := u.Message(true, "success", 200)
	resp["data"] = incident
	u.Respond(w, resp)
}

// HandledResolveRequest ...
var HandledResolveRequest = func(w http.ResponseWriter, r *http.Request) {
	f := &TicketID{}
	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		fmt.Println("error:", err, " acc:", f)
		u.Respond(w, u.Message(false, "Invalid request", 400))
		return
	}
	incident := model.GetIncidentByID(cast.ToInt64(f.ID))
	if incident == nil {
		resp := u.Message(true, "success", 400)
		resp["data"] = "invalid request"
		return
	}
	incident.Status = "close"
	model.UpdateIncidentByID(incident, cast.ToInt64(f.ID))
	resp := u.Message(true, "success", 200)
	resp["data"] = incident
	u.Respond(w, resp)

}

// HandledCommentRequest ...
var HandledCommentRequest = func(w http.ResponseWriter, r *http.Request) {

}
