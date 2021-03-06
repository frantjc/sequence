package sqnc

import (
	"context"
	"io"

	"github.com/bufbuild/connect-go"
	"github.com/frantjc/sequence/internal/rpcio"
	"github.com/frantjc/sequence/runtime"
)

func (c *sqncContainer) Exec(ctx context.Context, exec *runtime.Exec, streams *runtime.Streams) error {
	var (
		stdout = io.Discard
		stderr = stdout
	)
	if streams.Out != nil {
		stdout = streams.Out
	}
	if streams.Err != nil {
		stderr = streams.Err
	}

	stream, err := c.client.ExecContainer(ctx, connect.NewRequest(&ExecContainerRequest{
		Id:   c.GetID(),
		Exec: exec,
	}))
	if err != nil {
		return err
	}

	for stream.Receive() {
		var (
			msg = stream.Msg()
			b   = msg.GetLog().Data
		)
		switch rpcio.Stream(msg.Log.Stream) {
		case rpcio.StreamStdout:
			if _, err := stdout.Write(b); err != nil {
				return err
			}
		case rpcio.StreamStderr:
			if _, err := stderr.Write(b); err != nil {
				return err
			}
		}
	}

	if err = stream.Err(); err != nil {
		return err
	}

	return stream.Close()
}
