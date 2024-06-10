package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmosregistry/example"
	"github.com/stretchr/testify/require"
)

func TestQueryParams(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Params(f.ctx, &example.QueryParamsRequest{})
	require.NoError(err)
	require.Equal(example.Params{}, resp.Params)
}

func TestQueryCounter(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Counter(f.ctx, &example.QueryCounterRequest{Address: f.addrs[0].String()})
	require.NoError(err)
	require.Equal(uint64(0), resp.Counter)

	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[0].String()})
	require.NoError(err)

	resp, err = f.queryServer.Counter(f.ctx, &example.QueryCounterRequest{Address: f.addrs[0].String()})
	require.NoError(err)
	require.Equal(uint64(1), resp.Counter)
}

func TestQueryCounters(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{})
	require.NoError(err)
	require.Equal(0, len(resp.Counters))

	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[0].String()})
	require.NoError(err)

	resp, err = f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{})
	require.NoError(err)
	require.Equal(1, len(resp.Counters))
	require.Equal(uint64(1), resp.Counters[0].Count)
	require.Equal(f.addrs[0].String(), resp.Counters[0].Address)
}

func TestQueryCountersPaginated(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{Pagination: &query.PageRequest{Limit: 1}})
	require.NoError(err)
	require.Equal(0, len(resp.Counters))

	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[0].String()})
	require.NoError(err)
	_, err = f.msgServer.IncrementCounter(f.ctx, &example.MsgIncrementCounter{Sender: f.addrs[1].String()})
	require.NoError(err)

	resp, err = f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{Pagination: &query.PageRequest{Limit: 1}})
	require.NoError(err)
	require.Equal(1, len(resp.Counters))
	require.Equal(uint64(1), resp.Counters[0].Count)
	require.Equal(f.addrs[1].String(), resp.Counters[0].Address)

	resp, err = f.queryServer.Counters(f.ctx, &example.QueryCountersRequest{})
	require.NoError(err)
	require.Equal(2, len(resp.Counters))
}
