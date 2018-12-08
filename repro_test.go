package grpc_http2_bug

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)


func createGRPCTestServer(t *testing.T) net.Listener {
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "failed to bind listen socket")

	grpcServer := grpc.NewServer()
	RegisterTestServer(grpcServer, newGRPCTestServer())

	server := newHTTP2Server(listener, grpcServer)
	go server.Serve(listener)

	return listener
}

func TestGRPCLargePayload(t *testing.T) {
	ln := createGRPCTestServer(t)
	defer ln.Close()

	testClient := newGRPCTestClient(t, ln.Addr().String())

	var largePayload = strings.Repeat("ABCDEFGH", 1024*1024)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := testClient.Echo(ctx, &EchoRequest{Payload: largePayload})
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok, "Expect gRPC status for gRPC errors")
	assert.Equal(t, codes.ResourceExhausted, status.Code(), "Unexpected status code")

	for i := 0; i < 1; i++ {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			_, err := testClient.Echo(ctx, &EchoRequest{Payload: "test"})
			require.NoError(t, err, "Expect following requests to succeed")
		}()
	}
}


func newHTTP2Server(ln net.Listener, delegate http.Handler) *http.Server {
	return &http.Server{
		Addr: ln.Addr().String(),
		Handler: &http2Handler{
			delegate: delegate,
			h2Server: &http2.Server{},
		},
	}
}

type http2Handler struct {
	delegate http.Handler
	h2Server *http2.Server
}

func (h *http2Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.ProtoMajor != 2 {
		http.Error(w, fmt.Sprintf("request to HTTP/2 endpoint had proto %v", r.Proto), http.StatusBadRequest)
		return
	}

	h2c.NewHandler(h.delegate, h.h2Server).ServeHTTP(w, r)
}


func newGRPCTestClient(t *testing.T, addr string) TestClient {
	clientConn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err, "Failed to dial grpc server")
	return NewTestClient(clientConn)
}


type grpcTestServer struct {}

func newGRPCTestServer() *grpcTestServer {
	return &grpcTestServer{}
}

func (s *grpcTestServer) Echo(ctx context.Context, request *EchoRequest) (response *EchoResponse, err error) {
	return &EchoResponse{
		Payload: request.Payload,
	}, nil
}

