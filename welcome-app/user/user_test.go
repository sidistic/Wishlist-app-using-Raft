package user

import (
    "testing"
	"fmt"
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

func Equal(a []string, b [2]string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
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
    log.Printf("\nResponse: %v", resp)
	log.Printf("\nResponse: %v", resp.Followers)
	log.Printf("\nResponse: %v", resp.NotFollowers)

	trueFollowers := [...]string{"karanimal", "vihaha"}
	trueNotFollowers := [...]string{"sidistic", "bappi"}

	fmt.Println(Equal(resp.Followers, trueFollowers))
	fmt.Println(Equal(resp.NotFollowers, trueNotFollowers))

	if Equal(resp.Followers, trueFollowers) != true{
		t.Errorf("Error")
	}

	if Equal(resp.NotFollowers, trueNotFollowers) != true{
		t.Errorf("Error")
	}
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

	if resp.Success != true {
		t.Errorf("Error")
	}

}



func TestUpdateFollower(t *testing.T) {
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        t.Fatalf("Failed to dial bufnet: %v", err)
    }
    defer conn.Close()
    client := NewUserServiceClient(conn)

	// Test Case 1: Successfully follow previously unfollowed
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

	// Test Case 2: Successfully unfollow previously followed
	resp, err = client.UpdateFollower(ctx, &UpdateFollowersRequest{
		Username: "vihaha",
		Newuser:  "sidistic",
		IsFollow: true,
	})
    if err != nil {
        t.Fatalf("UpdateFollower failed: %v", err)
    }
    log.Printf("\nResponse: %+v", resp)

	if resp.Success != true {
		t.Errorf("Error")
	}
}

