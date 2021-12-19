package feed

import (
    "testing"
	// "fmt"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"net"
	"time"	
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(bufSize)
	u := Server{}
    Server := grpc.NewServer()
    RegisterFeedServiceServer(Server, &u)
    go func() {
        if err := Server.Serve(lis); err != nil {
            log.Fatalf("Server exited with error: %v", err)
        }
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestGetFeed(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewFeedServiceClient(conn)

	// positive test case
    resp, err := client.GetFeed(ctx, &FeedRequest{Username: "vihaha"})
    if err != nil {
        t.Fatalf("GetFeed failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)
	// if resp.Done != true {
	// 	t.Errorf("Error")
	// }


}



func TestPostToServer(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewFeedServiceClient(conn)

	// positive test case
    resp, err := client.PostToServer(ctx, &PostData{
	Postid:      0,
	Title:       "test title",
	Description: "test description",
	Author:      "vihaha",
	Timestamp:   time.Now().Format("01-02-2006 15:04:05"),
	})

    if err != nil {
        t.Fatalf("PostToServer failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)
	// if resp.Done != true {
	// 	t.Errorf("Error")
	// }


}