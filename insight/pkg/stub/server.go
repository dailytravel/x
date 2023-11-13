package stub

import (
	"github.com/dailytravel/x/proto/activity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing access metadata")
)

type Server struct {
	activity.UnimplementedActivityServer
}
