package controller

import (
	"log"

	"github.com/RedHatInsights/platform-receptor-controller/internal/receptor/protocol"

	"github.com/google/uuid"
)

type ReceptorService struct {
	Account    string
	NodeID     string
	PeerNodeID string

	Metadata interface{}

	// FIXME:  Move the channels into a Transport object/struct
	SendChannel    chan<- Message
	ControlChannel chan<- protocol.Message
	ErrorChannel   chan<- error

	/*
	   edges
	   seen
	*/
}

func (r *ReceptorService) RegisterConnection(peerNodeID string, metadata interface{}) error {
	log.Printf("Registering a connection to node %s", peerNodeID)

	r.PeerNodeID = peerNodeID
	r.Metadata = metadata

	return nil
}

func (r *ReceptorService) UpdateRoutingTable(edges string, seen string) error {
	log.Println("edges:", edges)
	log.Println("seen:", seen)

	return nil
}

func (r *ReceptorService) SendMessage(recipient string, route []string, payload interface{}, directive string) (*uuid.UUID, error) {

	jobID, err := uuid.NewRandom()
	if err != nil {
		log.Println("Unable to generate UUID for routing the job...cannot proceed")
		return nil, err
	}

	msg := Message{MessageID: jobID,
		Recipient: recipient,
		RouteList: route,
		Payload:   payload,
		Directive: directive}

	r.SendChannel <- msg

	return &jobID, nil
}

func (r *ReceptorService) Close() {
}

func (r *ReceptorService) DisconnectReceptorNetwork() {
}
