package pkg

import (
	"context"
	"fmt"
	"sync"

	proto "github.com/Oik17/gRPC-game/gen"

	database "github.com/Oik17/gRPC-game/db/init"
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

func (p *Pool) CreateStream(req *proto.Connect, stream proto.GameService_CreateStreamServer) error {
	if req == nil || req.User == nil {
		return fmt.Errorf("invalid connection request: user data missing")
	}

	conn := &Connection{
		stream: stream,
		id:     req.User.Id,
		active: true,
		error:  make(chan error),
	}

	p.Connection = append(p.Connection, conn)

	fmt.Printf("User %v connected and listening for responses.\n", req.User.Id)

	<-stream.Context().Done()

	p.removeConnection(conn)

	fmt.Printf("User %v disconnected.\n", req.User.Id)
	return nil
}

func (p *Pool) removeConnection(conn *Connection) {
	for i, c := range p.Connection {
		if c.id == conn.id {
			p.Connection = append(p.Connection[:i], p.Connection[i+1:]...)
			break
		}
	}
}

func (s *Pool) SubmitAnswer(ctx context.Context, response *proto.Response) (*proto.Close, error) {

	queries, err := database.InitDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

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
				answer, err := queries.GetAnswer(ctx, response.Question)
				if err != nil {
					fmt.Printf("Error fetching answer: %v\n", err)
					conn.active = false
					select {
					case conn.error <- err:
					default:
						fmt.Println("Error channel full, dropping error")
					}
					return
				}

				correctAnswer := answer.String
				isCorrect := response.Answer == correctAnswer

				fmt.Printf("User %s answered: %s | Correct Answer: %s | Correct: %v\n",
					response.User.Id, response.Answer, correctAnswer, isCorrect)

				response.IsCorrect = isCorrect

				if err := conn.stream.Send(response); err != nil {
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
