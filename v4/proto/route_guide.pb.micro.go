// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/route_guide.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RouteGuide service

func NewRouteGuideEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RouteGuide service

type RouteGuideService interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(ctx context.Context, in *Point, opts ...client.CallOption) (*Feature, error)
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(ctx context.Context, in *Rectangle, opts ...client.CallOption) (RouteGuide_ListFeaturesService, error)
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(ctx context.Context, opts ...client.CallOption) (RouteGuide_RecordRouteService, error)
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(ctx context.Context, opts ...client.CallOption) (RouteGuide_RouteChatService, error)
}

type routeGuideService struct {
	c    client.Client
	name string
}

func NewRouteGuideService(name string, c client.Client) RouteGuideService {
	return &routeGuideService{
		c:    c,
		name: name,
	}
}

func (c *routeGuideService) GetFeature(ctx context.Context, in *Point, opts ...client.CallOption) (*Feature, error) {
	req := c.c.NewRequest(c.name, "RouteGuide.GetFeature", in)
	out := new(Feature)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeGuideService) ListFeatures(ctx context.Context, in *Rectangle, opts ...client.CallOption) (RouteGuide_ListFeaturesService, error) {
	req := c.c.NewRequest(c.name, "RouteGuide.ListFeatures", &Rectangle{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &routeGuideServiceListFeatures{stream}, nil
}

type RouteGuide_ListFeaturesService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*Feature, error)
}

type routeGuideServiceListFeatures struct {
	stream client.Stream
}

func (x *routeGuideServiceListFeatures) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *routeGuideServiceListFeatures) Close() error {
	return x.stream.Close()
}

func (x *routeGuideServiceListFeatures) Context() context.Context {
	return x.stream.Context()
}

func (x *routeGuideServiceListFeatures) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routeGuideServiceListFeatures) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routeGuideServiceListFeatures) Recv() (*Feature, error) {
	m := new(Feature)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideService) RecordRoute(ctx context.Context, opts ...client.CallOption) (RouteGuide_RecordRouteService, error) {
	req := c.c.NewRequest(c.name, "RouteGuide.RecordRoute", &Point{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &routeGuideServiceRecordRoute{stream}, nil
}

type RouteGuide_RecordRouteService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*Point) error
}

type routeGuideServiceRecordRoute struct {
	stream client.Stream
}

func (x *routeGuideServiceRecordRoute) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *routeGuideServiceRecordRoute) Close() error {
	return x.stream.Close()
}

func (x *routeGuideServiceRecordRoute) Context() context.Context {
	return x.stream.Context()
}

func (x *routeGuideServiceRecordRoute) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routeGuideServiceRecordRoute) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routeGuideServiceRecordRoute) Send(m *Point) error {
	return x.stream.Send(m)
}

func (c *routeGuideService) RouteChat(ctx context.Context, opts ...client.CallOption) (RouteGuide_RouteChatService, error) {
	req := c.c.NewRequest(c.name, "RouteGuide.RouteChat", &RouteNote{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &routeGuideServiceRouteChat{stream}, nil
}

type RouteGuide_RouteChatService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
}

type routeGuideServiceRouteChat struct {
	stream client.Stream
}

func (x *routeGuideServiceRouteChat) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *routeGuideServiceRouteChat) Close() error {
	return x.stream.Close()
}

func (x *routeGuideServiceRouteChat) Context() context.Context {
	return x.stream.Context()
}

func (x *routeGuideServiceRouteChat) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routeGuideServiceRouteChat) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routeGuideServiceRouteChat) Send(m *RouteNote) error {
	return x.stream.Send(m)
}

func (x *routeGuideServiceRouteChat) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RouteGuide service

type RouteGuideHandler interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	GetFeature(context.Context, *Point, *Feature) error
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	ListFeatures(context.Context, *Rectangle, RouteGuide_ListFeaturesStream) error
	// A client-to-server streaming RPC.
	//
	// Accepts a stream of Points on a route being traversed, returning a
	// RouteSummary when traversal is completed.
	RecordRoute(context.Context, RouteGuide_RecordRouteStream) error
	// A Bidirectional streaming RPC.
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	RouteChat(context.Context, RouteGuide_RouteChatStream) error
}

func RegisterRouteGuideHandler(s server.Server, hdlr RouteGuideHandler, opts ...server.HandlerOption) error {
	type routeGuide interface {
		GetFeature(ctx context.Context, in *Point, out *Feature) error
		ListFeatures(ctx context.Context, stream server.Stream) error
		RecordRoute(ctx context.Context, stream server.Stream) error
		RouteChat(ctx context.Context, stream server.Stream) error
	}
	type RouteGuide struct {
		routeGuide
	}
	h := &routeGuideHandler{hdlr}
	return s.Handle(s.NewHandler(&RouteGuide{h}, opts...))
}

type routeGuideHandler struct {
	RouteGuideHandler
}

func (h *routeGuideHandler) GetFeature(ctx context.Context, in *Point, out *Feature) error {
	return h.RouteGuideHandler.GetFeature(ctx, in, out)
}

func (h *routeGuideHandler) ListFeatures(ctx context.Context, stream server.Stream) error {
	m := new(Rectangle)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RouteGuideHandler.ListFeatures(ctx, m, &routeGuideListFeaturesStream{stream})
}

type RouteGuide_ListFeaturesStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Feature) error
}

type routeGuideListFeaturesStream struct {
	stream server.Stream
}

func (x *routeGuideListFeaturesStream) Close() error {
	return x.stream.Close()
}

func (x *routeGuideListFeaturesStream) Context() context.Context {
	return x.stream.Context()
}

func (x *routeGuideListFeaturesStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routeGuideListFeaturesStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routeGuideListFeaturesStream) Send(m *Feature) error {
	return x.stream.Send(m)
}

func (h *routeGuideHandler) RecordRoute(ctx context.Context, stream server.Stream) error {
	return h.RouteGuideHandler.RecordRoute(ctx, &routeGuideRecordRouteStream{stream})
}

type RouteGuide_RecordRouteStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Point, error)
}

type routeGuideRecordRouteStream struct {
	stream server.Stream
}

func (x *routeGuideRecordRouteStream) Close() error {
	return x.stream.Close()
}

func (x *routeGuideRecordRouteStream) Context() context.Context {
	return x.stream.Context()
}

func (x *routeGuideRecordRouteStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routeGuideRecordRouteStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routeGuideRecordRouteStream) Recv() (*Point, error) {
	m := new(Point)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *routeGuideHandler) RouteChat(ctx context.Context, stream server.Stream) error {
	return h.RouteGuideHandler.RouteChat(ctx, &routeGuideRouteChatStream{stream})
}

type RouteGuide_RouteChatStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*RouteNote) error
	Recv() (*RouteNote, error)
}

type routeGuideRouteChatStream struct {
	stream server.Stream
}

func (x *routeGuideRouteChatStream) Close() error {
	return x.stream.Close()
}

func (x *routeGuideRouteChatStream) Context() context.Context {
	return x.stream.Context()
}

func (x *routeGuideRouteChatStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *routeGuideRouteChatStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *routeGuideRouteChatStream) Send(m *RouteNote) error {
	return x.stream.Send(m)
}

func (x *routeGuideRouteChatStream) Recv() (*RouteNote, error) {
	m := new(RouteNote)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
