package main

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"gim/config"
	"gim/internal/logic/device"
	"gim/internal/logic/friend"
	"gim/internal/logic/group"
	"gim/internal/logic/message"
	"gim/internal/logic/room"
	"gim/pkg/interceptor"
	"gim/pkg/logger"
	pb "gim/pkg/protocol/pb/logicpb"
	"gim/pkg/urlwhitelist"
)

func main() {
	logger.Init("logic")

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.NewInterceptor(urlwhitelist.Logic)))

	// 监听服务关闭信号，服务平滑重启
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		slog.Info("server stop", "signal", s)
		server.GracefulStop()
	}()

	pb.RegisterDeviceExtServiceServer(server, &device.DeviceExtService{})
	pb.RegisterDeviceIntServiceServer(server, &device.DeviceIntService{})
	pb.RegisterMessageIntServiceServer(server, &message.MessageIntService{})
	pb.RegisterFriendExtServiceServer(server, &friend.FriendExtService{})
	pb.RegisterGroupExtServiceServer(server, &group.GroupExtService{})
	pb.RegisterRoomExtServiceServer(server, &room.RoomExtService{})
	pb.RegisterRoomIntServiceServer(server, &room.RoomIntService{})

	listen, err := net.Listen("tcp", config.Config.LogicRPCListenAddr)
	if err != nil {
		panic(err)
	}

	slog.Info("rpc服务已经开启")
	err = server.Serve(listen)
	if err != nil {
		slog.Error("serve error", "error", err)
	}
}
