package gRPC

import(
	"log"
	"golang.org/x/net/context"
	"time"
)

type Server struct{

}

func(s *Server) WhoDat(ctx context.Context, timerequest *timeRequest) (*time , error){
	log.Printf("received time request from client: %s" timerequest.body)
	return &time{timeNow : time.Now()}
}