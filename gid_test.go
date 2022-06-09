package getgroup_test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/liucxer/getgroup"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewGroup(t *testing.T) {
	group, err := getgroup.NewGroup(10, 10*1024*1024)
	require.NoError(t, err)
	spew.Dump(group)
}
