package plugin

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/mitchellh/devflow/sdk/component"
	"github.com/mitchellh/devflow/sdk/internal-shared/mapper"
	"github.com/mitchellh/devflow/sdk/internal/funcspec"
	"github.com/mitchellh/devflow/sdk/proto"
)

// PlatformPlugin implements plugin.Plugin (specifically GRPCPlugin) for
// the Platform component type.
type PlatformPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl    component.Platform // Impl is the concrete implementation
	Mappers []*mapper.Func     // Mappers
	Logger  hclog.Logger       // Logger
}

func (p *PlatformPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterPlatformServer(s, &platformServer{
		Impl:    p.Impl,
		Mappers: p.Mappers,
		Logger:  p.Logger,
	})
	return nil
}

func (p *PlatformPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	// Keep track of all our impl types
	var platformLog component.LogPlatform

	// Build our client to the platform service
	client := &platformClient{
		client: proto.NewPlatformClient(c),
		logger: p.Logger,
	}

	// Check if we also implement the LogPlatform
	resp, err := client.client.IsLogPlatform(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	if resp.Implements {
		raw, err := (&LogPlatformPlugin{
			Logger: p.Logger,
		}).GRPCClient(ctx, broker, c)
		if err != nil {
			return nil, err
		}

		platformLog = raw.(component.LogPlatform)
	}

	// Build our result
	var result interface{} = client
	if platformLog != nil {
		p.Logger.Info("platform plugin capable of logs")
		result = &platform_Log{
			ConfigurableNotify: client,
			Platform:           client,
			LogPlatform:        platformLog,
		}
	}

	return result, nil
}

// platformClient is an implementation of component.Platform over gRPC.
type platformClient struct {
	client proto.PlatformClient
	logger hclog.Logger
}

func (c *platformClient) Config() (interface{}, error) {
	return configStructCall(context.Background(), c.client)
}

func (c *platformClient) ConfigSet(v interface{}) error {
	return configureCall(context.Background(), c.client, v)
}

func (c *platformClient) DeployFunc() interface{} {
	// Get the spec
	spec, err := c.client.DeploySpec(context.Background(), &proto.Empty{})
	if err != nil {
		return funcErr(err)
	}

	return funcspec.Func(spec, c.push, funcspec.WithLogger(c.logger))
}

func (c *platformClient) push(
	ctx context.Context,
	args funcspec.Args,
) (interface{}, error) {
	// Call our function
	resp, err := c.client.Deploy(ctx, &proto.Deploy_Args{Args: args})
	if err != nil {
		return nil, err
	}

	// We return the *any.Any directly.
	return resp.Result, nil
}

// platformServer is a gRPC server that the client talks to and calls a
// real implementation of the component.
type platformServer struct {
	Impl    component.Platform
	Mappers []*mapper.Func
	Logger  hclog.Logger
}

func (s *platformServer) IsLogPlatform(
	ctx context.Context,
	empty *empty.Empty,
) (*proto.ImplementsResp, error) {
	_, ok := s.Impl.(component.LogPlatform)
	return &proto.ImplementsResp{Implements: ok}, nil
}

func (s *platformServer) ConfigStruct(
	ctx context.Context,
	empty *empty.Empty,
) (*proto.Config_StructResp, error) {
	return configStruct(s.Impl)
}

func (s *platformServer) Configure(
	ctx context.Context,
	req *proto.Config_ConfigureRequest,
) (*empty.Empty, error) {
	return configure(s.Impl, req)
}

func (s *platformServer) DeploySpec(
	ctx context.Context,
	args *proto.Empty,
) (*proto.FuncSpec, error) {
	return funcspec.Spec(s.Impl.DeployFunc(),
		funcspec.WithMappers(s.Mappers),
		funcspec.WithLogger(s.Logger))
}

func (s *platformServer) Deploy(
	ctx context.Context,
	args *proto.Deploy_Args,
) (*proto.Deploy_Resp, error) {
	encoded, err := callDynamicFuncAny(ctx, s.Logger, args.Args, s.Impl.DeployFunc(), s.Mappers)
	if err != nil {
		return nil, err
	}

	return &proto.Deploy_Resp{Result: encoded}, nil
}

var (
	_ plugin.Plugin                = (*PlatformPlugin)(nil)
	_ plugin.GRPCPlugin            = (*PlatformPlugin)(nil)
	_ proto.PlatformServer         = (*platformServer)(nil)
	_ component.Platform           = (*platformClient)(nil)
	_ component.Configurable       = (*platformClient)(nil)
	_ component.ConfigurableNotify = (*platformClient)(nil)
)