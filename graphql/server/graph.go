package server

type Server struct {
	// accountClient *account.Client
	// catalogClient *catalog.Client
	// orderClient *order.Client
}

func NewGraphQlServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	// accountClient, err := account.NewClient(accountUrl)
	// defer accountClient.Close()
	// if err != nil {
	// 	return nil,err
	// }
	// catalogClient, err := catalog.NewClient(catalogUrl)
	// defer catalogClient.Close()
	// if err != nil {
	// 	return nil,err
	// }
	// orderClient, err := order.NewClient(orderUrl)
	// defer orderClient.Close()
	// if err != nil {
	// 	return nil,err
	// }

	return &Server{
		// accountClient,
		// catalogClient,
		// orderClient,
	}, nil
}

func (s *Server) Mutation() {
	// return &Resolver{

	// }
}
