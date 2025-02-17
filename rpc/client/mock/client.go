package mock

/*
package mock returns a Client implementation that
accepts various (mock) implementations of the various methods.

This implementation is useful for using in tests, when you don't
need a real server, but want a high-level of control about
the server response you want to mock (eg. error handling),
or if you just want to record the calls to verify in your tests.

For real clients, you probably want the "http" package.  If you
want to directly call a tendermint node in process, you can use the
"local" package.
*/

import (
	"context"
	"reflect"

	"github.com/tendermint/tendermint/internal/rpc/core"
	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/coretypes"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"github.com/tendermint/tendermint/types"
)

// Client wraps arbitrary implementations of the various interfaces.
type Client struct {
	client.Client
	env *core.Environment
}

func New() Client {
	return Client{
		env: &core.Environment{},
	}
}

var _ client.Client = Client{}

// Call is used by recorders to save a call and response.
// It can also be used to configure mock responses.
//
type Call struct {
	Name     string
	Args     interface{}
	Response interface{}
	Error    error
}

// GetResponse will generate the apporiate response for us, when
// using the Call struct to configure a Mock handler.
//
// When configuring a response, if only one of Response or Error is
// set then that will always be returned. If both are set, then
// we return Response if the Args match the set args, Error otherwise.
func (c Call) GetResponse(args interface{}) (interface{}, error) {
	// handle the case with no response
	if c.Response == nil {
		if c.Error == nil {
			panic("Misconfigured call, you must set either Response or Error")
		}
		return nil, c.Error
	}
	// response without error
	if c.Error == nil {
		return c.Response, nil
	}
	// have both, we must check args....
	if reflect.DeepEqual(args, c.Args) {
		return c.Response, nil
	}
	return nil, c.Error
}

func (c Client) Status(ctx context.Context) (*coretypes.ResultStatus, error) {
	return c.env.Status(&rpctypes.Context{})
}

func (c Client) ABCIInfo(ctx context.Context) (*coretypes.ResultABCIInfo, error) {
	return c.env.ABCIInfo(&rpctypes.Context{})
}

func (c Client) ABCIQuery(ctx context.Context, path string, data bytes.HexBytes) (*coretypes.ResultABCIQuery, error) {
	return c.ABCIQueryWithOptions(ctx, path, data, client.DefaultABCIQueryOptions)
}

func (c Client) ABCIQueryWithOptions(
	ctx context.Context,
	path string,
	data bytes.HexBytes,
	opts client.ABCIQueryOptions) (*coretypes.ResultABCIQuery, error) {
	return c.env.ABCIQuery(&rpctypes.Context{}, path, data, opts.Height, opts.Prove)
}

func (c Client) BroadcastTxCommit(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTxCommit, error) {
	return c.env.BroadcastTxCommit(&rpctypes.Context{}, tx)
}

func (c Client) BroadcastTxAsync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	return c.env.BroadcastTxAsync(&rpctypes.Context{}, tx)
}

func (c Client) BroadcastTxSync(ctx context.Context, tx types.Tx) (*coretypes.ResultBroadcastTx, error) {
	return c.env.BroadcastTxSync(&rpctypes.Context{}, tx)
}

func (c Client) CheckTx(ctx context.Context, tx types.Tx) (*coretypes.ResultCheckTx, error) {
	return c.env.CheckTx(&rpctypes.Context{}, tx)
}

func (c Client) NetInfo(ctx context.Context) (*coretypes.ResultNetInfo, error) {
	return c.env.NetInfo(&rpctypes.Context{})
}

func (c Client) ConsensusState(ctx context.Context) (*coretypes.ResultConsensusState, error) {
	return c.env.GetConsensusState(&rpctypes.Context{})
}

func (c Client) DumpConsensusState(ctx context.Context) (*coretypes.ResultDumpConsensusState, error) {
	return c.env.DumpConsensusState(&rpctypes.Context{})
}

func (c Client) ConsensusParams(ctx context.Context, height *int64) (*coretypes.ResultConsensusParams, error) {
	return c.env.ConsensusParams(&rpctypes.Context{}, height)
}

func (c Client) Health(ctx context.Context) (*coretypes.ResultHealth, error) {
	return c.env.Health(&rpctypes.Context{})
}

func (c Client) DialSeeds(ctx context.Context, seeds []string) (*coretypes.ResultDialSeeds, error) {
	return c.env.UnsafeDialSeeds(&rpctypes.Context{}, seeds)
}

func (c Client) DialPeers(
	ctx context.Context,
	peers []string,
	persistent,
	unconditional,
	private bool,
) (*coretypes.ResultDialPeers, error) {
	return c.env.UnsafeDialPeers(&rpctypes.Context{}, peers, persistent, unconditional, private)
}

func (c Client) BlockchainInfo(ctx context.Context, minHeight, maxHeight int64) (*coretypes.ResultBlockchainInfo, error) { //nolint:lll
	return c.env.BlockchainInfo(&rpctypes.Context{}, minHeight, maxHeight)
}

func (c Client) Genesis(ctx context.Context) (*coretypes.ResultGenesis, error) {
	return c.env.Genesis(&rpctypes.Context{})
}

func (c Client) Block(ctx context.Context, height *int64) (*coretypes.ResultBlock, error) {
	return c.env.Block(&rpctypes.Context{}, height)
}

func (c Client) BlockByHash(ctx context.Context, hash bytes.HexBytes) (*coretypes.ResultBlock, error) {
	return c.env.BlockByHash(&rpctypes.Context{}, hash)
}

func (c Client) Commit(ctx context.Context, height *int64) (*coretypes.ResultCommit, error) {
	return c.env.Commit(&rpctypes.Context{}, height)
}

func (c Client) Validators(ctx context.Context, height *int64, page, perPage *int) (*coretypes.ResultValidators, error) { //nolint:lll
	return c.env.Validators(&rpctypes.Context{}, height, page, perPage)
}

func (c Client) BroadcastEvidence(ctx context.Context, ev types.Evidence) (*coretypes.ResultBroadcastEvidence, error) {
	return c.env.BroadcastEvidence(&rpctypes.Context{}, ev)
}
