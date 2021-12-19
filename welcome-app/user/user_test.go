package user

import (
    "testing"
	// "fmt"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"net"
)

// just for testing testing lol
func TestDummyFunc(t *testing.T) {
    if DummyFunc(2) != 4 {
        t.Error("Expected 2 + 2 to equal 4")
    }
}



const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
    lis = bufconn.Listen(bufSize)
	u := Server{}
    Server := grpc.NewServer()
    RegisterUserServiceServer(Server, &u)
    go func() {
        if err := Server.Serve(lis); err != nil {
            log.Fatalf("Server exited with error: %v", err)
        }
    }()
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestGetFollowing(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewUserServiceClient(conn)
    resp, err := client.GetFollowing(ctx, &FollowerRequest{Username: "vihaha"})
    if err != nil {
        t.Fatalf("GetFollowing failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)

	// if resp.NotFollowers != "bappi" {
	// 	t.Errorf("Error")
	// }

}



func TestSignUpUser(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewUserServiceClient(conn)
    resp, err := client.SignUpUser(ctx, &SignUpRequest{
		Username:        "testuser",
		Password:        "testpass",
		Name:            "testname",
		ConfirmPassword: "testpass",
	})
    if err != nil {
        t.Fatalf("SignUpUser failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)

	// if resp.NotFollowers != "bappi" {
	// 	t.Errorf("Error")
	// }

}



func TestUpdateFollower(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewUserServiceClient(conn)
    resp, err := client.UpdateFollower(ctx, &UpdateFollowersRequest{
		Username: "vihaha",
		Newuser:  "sidistic",
		IsFollow: false,
	})
    if err != nil {
        t.Fatalf("UpdateFollower failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)

	if resp.Success != true {
		t.Errorf("Error")
	}

}



// func TestCheckIfFollowed(t *testing.T) {

// 	follows := [...]string{"viha", "sid", "bharath"}

// 	// Positive test case
//     if CheckIfFollowed("viha", follows) != true{
//         t.Error("checkIfFollowed Failed")
//     }

// 	// Negative test case
// 	if CheckIfFollowed("karan", follows) == true{
//         t.Error("checkIfFollowed Failed")
//     }

// }