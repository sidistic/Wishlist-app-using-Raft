package login

import (
    "testing"
	// "fmt"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"net"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(bufSize)
	u := Server{}
    Server := grpc.NewServer()
    RegisterAuthServiceServer(Server, &u)
    go func() {
        if err := Server.Serve(lis); err != nil {
            log.Fatalf("Server exited with error: %v", err)
        }
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestAuthenticate(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewAuthServiceClient(conn)

	// Test Case 1: Correct username and password - should pass authentication
    resp, err := client.Authenticate(ctx, &LoginDetails{Username: "vihaha", Password: "v"})
    if err != nil {
        t.Fatalf("Authenticate failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)
	if resp.Done != true {
		t.Errorf("Error")
	}

	// Test Case: Wrong username and password - should fail to authenticate
	resp, err = client.Authenticate(ctx, &LoginDetails{Username: "vihaha", Password: "wrongpass"})
    if err != nil {
        t.Fatalf("Authenticate failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)
	if resp.Done == true {
		t.Errorf("Error")
	}

}