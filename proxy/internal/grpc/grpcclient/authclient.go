package grpcclient

/*
type ClientAuth struct{}

func (c *ClientAuth) CallRegister(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pbauth.NewAuthServiceClient(conn)

	res, err := client.Register(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}

	log.Printf("Ответ от сервера: %s", res.Token)
	return res, nil
}


*/
