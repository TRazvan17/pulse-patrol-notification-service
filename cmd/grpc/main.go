package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/TRazvan17/pulse-patrol-notification-service/proto"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *server) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "ok"}, nil
}

func (s *server) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	if req.GetTo() == "" || req.GetMessage() == "" {
		return &pb.SendNotificationResponse{Id: "", Status: "invalid"}, nil
	}
	return &pb.SendNotificationResponse{Id: "notif-1", Status: "queued"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, &server{})

	// Enable reflection so tools like grpcurl can discover services
	reflection.Register(s)

	log.Println("gRPC listening on :9090")
	log.Fatal(s.Serve(lis))
}
