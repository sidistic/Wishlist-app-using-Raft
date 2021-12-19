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

	// Test Case 1: Valid user
    resp, err := client.GetFeed(ctx, &FeedRequest{Username: "karanimal"})
    if err != nil {
        t.Fatalf("GetFeed failed: %v", err)
    }
    // log.Printf("\nResponse: %+v", resp)
	// log.Printf("\nResponse: %+v", resp.Postid[0])
	// log.Printf("\nResponse: %+v", resp.Title[0])
	// log.Printf("\nResponse: %+v", resp.Author[0])
	// log.Printf("\nResponse: %+v", resp.Description[0])
	// log.Printf("\nResponse: %+v", resp.Timestamp[0])
	if (resp.Postid[0] != 0) ||
		(resp.Title[0] != "Robovacuum") ||
		(resp.Author[0] != "vihaha") ||
		(resp.Description[0] != "a splendid vacuum for my house") ||
		(resp.Timestamp[0] != "72834492") {
		t.Errorf("Error")
	}

	// Test Case 1: Invalid user
    resp, err = client.GetFeed(ctx, &FeedRequest{Username: "gustavo"})
    if err != nil {
        t.Fatalf("GetFeed failed: %v", err)
    }
	// log.Printf("\nResponse: %+v", resp)
	// log.Printf("\nResponse: %+v", resp.Postid)
	// log.Printf("\nResponse: %+v", len(resp.Postid))
	if len(resp.Postid) != 0 {
		t.Errorf("Error")
	}
}



func TestPostToServer(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewFeedServiceClient(conn)

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
	if resp.Success != true {
		t.Errorf("Error")
	}


}