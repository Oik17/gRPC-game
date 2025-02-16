package pkg

import (
	"context"
	"fmt"
	"sync"

	proto "github.com/Oik17/gRPC-game/gen"
)

type Connection struct {
	proto.UnimplementedGameServiceServer
	stream proto.GameService_CreateStreamServer
	id     string
	active bool
	error  chan error
}

type Pool struct {
	proto.UnimplementedGameServiceServer
	Connection []*Connection
}

func (p *Pool) CreateStream(pconn *proto.Connect, stream proto.GameService_CreateStreamServer) error {
	if pconn == nil || pconn.User == nil {
		fmt.Println("Error: Received nil pconn or pconn.User")
		return fmt.Errorf("invalid connection request: user information is missing")
	}

	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}

	if p.Connection == nil {
		p.Connection = []*Connection{}
	}
	p.Connection = append(p.Connection, conn)

	fmt.Printf("User %v connected successfully\n", conn.id)

	return nil
}

func (s *Pool) SubmitAnswer(ctx context.Context, response *proto.Response) (*proto.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan struct{})

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(conn *Connection, response *proto.Response) {

			defer wait.Done()

			if conn == nil {
				fmt.Println("Skipping nil connection")
				return
			}
			if conn.active {
				fmt.Printf("User %s answered: %s | Correct: %v\n", response.User.Id, response.Answer, response.IsCorrect)

				err := conn.stream.Send(response)
				if err != nil {
					fmt.Printf("Error with Stream: %v - Error: %v\n", conn.stream, err)
					conn.active = false
					select {
					case conn.error <- err:
					default:
						fmt.Println("Error channel full, dropping error")
					}
				}
			}
		}(conn, response)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &proto.Close{}, nil
}
