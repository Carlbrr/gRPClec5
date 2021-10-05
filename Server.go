package serverLec5GO

import(
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp",":9000")
	if err != nil{
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := time.Server{}

	grpcServer := grpc.NewServer()

	time.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("failed to serve...")
	}


}
