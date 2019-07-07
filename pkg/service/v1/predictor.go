package v1

import (
	"context"
	"fmt"
	"time"

	// "database/sql"
	// "fmt"
	// "time"

	// "github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// gw "bitbucket.org/shchukin_a/go-predictor/api"
	core "bitbucket.org/shchukin_a/go-predictor/internal/core"
	pb "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"

	// "bitbucket.org/shchukin_a/go-predictor/pkg/logger"

	v1 "bitbucket.org/shchukin_a/go-predictor/pkg/api/v1"

	// "github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/struct"
	// "go.uber.org/zap"
	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type predictorServer struct {
	// db *sql.DB
}

// NewPredictorServer creates ToDo service
func NewPredictorServer() v1.PredictorServer {
	return &predictorServer{}
}

func (s *predictorServer) GetBaseMatrix(ctx context.Context, e *empty.Empty) (*pb.Matrix, error) {
	values := []*structpb.Value{}

	for _, v := range core.OriginList {
		values = append(values, &structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: float64(v)}})
	}

	return &pb.Matrix{
		M: &structpb.ListValue{
			Values: values,
		},
	}, nil
}

func (s *predictorServer) FindCardByBirthday(ctx context.Context, date *pb.Date) (*pb.Card, error) {
	d := time.Date(int(date.Year), time.Month(int(date.Month)), int(date.Day), 0, 0, 0, 0, time.UTC)
	fmt.Printf("got birthday: %v\n", d)

	if d == time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) {
		return nil, status.Error(codes.Unknown, "index out of range")
	}

	card := core.FindCardByBirthday(&d)

	return &pb.Card{
		Number: uint32(card.Order),
		Suite:  card.Suit,
		Rank:   card.Rank,
	}, nil
}

/*
 *
 * // checkAPI checks if the API version requested by client is supported by server
 * func (s *toDoServiceServer) checkAPI(api string) error {
 *     // API version is "" means use current version of the service
 *     if len(api) > 0 {
 *         if apiVersion != api {
 *             return status.Errorf(codes.Unimplemented,
 *                 "unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
 *         }
 *     }
 *     return nil
 * }
 *
 * // connect returns SQL database connection from the pool
 * func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
 *     c, err := s.db.Conn(ctx)
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
 *     }
 *     return c, nil
 * }
 *
 * // Create new todo task
 * func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
 *     // check if the API version requested by client is supported by server
 *     if err := s.checkAPI(req.Api); err != nil {
 *         return nil, err
 *     }
 *
 *     // get SQL connection from pool
 *     c, err := s.connect(ctx)
 *     if err != nil {
 *         return nil, err
 *     }
 *     defer c.Close()
 *
 *     reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
 *     if err != nil {
 *         return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
 *     }
 *
 *     // insert ToDo entity data
 *     res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
 *         req.ToDo.Title, req.ToDo.Description, reminder)
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
 *     }
 *
 *     // get ID of creates ToDo
 *     id, err := res.LastInsertId()
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
 *     }
 *
 *     return &v1.CreateResponse{
 *         Api: apiVersion,
 *         Id:  id,
 *     }, nil
 * }
 *
 * // Read todo task
 * func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
 *     // check if the API version requested by client is supported by server
 *     if err := s.checkAPI(req.Api); err != nil {
 *         return nil, err
 *     }
 *
 *     // get SQL connection from pool
 *     c, err := s.connect(ctx)
 *     if err != nil {
 *         return nil, err
 *     }
 *     defer c.Close()
 *
 *     // query ToDo by ID
 *     rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",
 *         req.Id)
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
 *     }
 *     defer rows.Close()
 *
 *     if !rows.Next() {
 *         if err := rows.Err(); err != nil {
 *             return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
 *         }
 *         return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
 *             req.Id))
 *     }
 *
 *     // get ToDo data
 *     var td v1.ToDo
 *     var reminder time.Time
 *     if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
 *     }
 *     td.Reminder, err = ptypes.TimestampProto(reminder)
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
 *     }
 *
 *     if rows.Next() {
 *         return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%d'",
 *             req.Id))
 *     }
 *
 *     return &v1.ReadResponse{
 *         Api:  apiVersion,
 *         ToDo: &td,
 *     }, nil
 *
 * }
 *
 * // Update todo task
 * func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
 *     // check if the API version requested by client is supported by server
 *     if err := s.checkAPI(req.Api); err != nil {
 *         return nil, err
 *     }
 *
 *     // get SQL connection from pool
 *     c, err := s.connect(ctx)
 *     if err != nil {
 *         return nil, err
 *     }
 *     defer c.Close()
 *
 *     reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
 *     if err != nil {
 *         return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
 *     }
 *
 *     // update ToDo
 *     res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
 *         req.ToDo.Title, req.ToDo.Description, reminder, req.ToDo.Id)
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to update ToDo-> "+err.Error())
 *     }
 *
 *     rows, err := res.RowsAffected()
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
 *     }
 *
 *     if rows == 0 {
 *         return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
 *             req.ToDo.Id))
 *     }
 *
 *     return &v1.UpdateResponse{
 *         Api:     apiVersion,
 *         Updated: rows,
 *     }, nil
 * }
 *
 * // Delete todo task
 * func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
 *     // check if the API version requested by client is supported by server
 *     if err := s.checkAPI(req.Api); err != nil {
 *         return nil, err
 *     }
 *
 *     // get SQL connection from pool
 *     c, err := s.connect(ctx)
 *     if err != nil {
 *         return nil, err
 *     }
 *     defer c.Close()
 *
 *     // delete ToDo
 *     res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", req.Id)
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to delete ToDo-> "+err.Error())
 *     }
 *
 *     rows, err := res.RowsAffected()
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
 *     }
 *
 *     if rows == 0 {
 *         return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
 *             req.Id))
 *     }
 *
 *     return &v1.DeleteResponse{
 *         Api:     apiVersion,
 *         Deleted: rows,
 *     }, nil
 * }
 *
 * // Read all todo tasks
 * func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
 *     // check if the API version requested by client is supported by server
 *     if err := s.checkAPI(req.Api); err != nil {
 *         return nil, err
 *     }
 *
 *     // get SQL connection from pool
 *     c, err := s.connect(ctx)
 *     if err != nil {
 *         return nil, err
 *     }
 *     defer c.Close()
 *
 *     // get ToDo list
 *     rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
 *     if err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
 *     }
 *     defer rows.Close()
 *
 *     var reminder time.Time
 *     list := []*v1.ToDo{}
 *     for rows.Next() {
 *         td := new(v1.ToDo)
 *         if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
 *             return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
 *         }
 *         td.Reminder, err = ptypes.TimestampProto(reminder)
 *         if err != nil {
 *             return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
 *         }
 *         list = append(list, td)
 *     }
 *
 *     if err := rows.Err(); err != nil {
 *         return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
 *     }
 *
 *     return &v1.ReadAllResponse{
 *         Api:   apiVersion,
 *         ToDos: list,
 *     }, nil
 * }
 */
