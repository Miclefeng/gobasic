package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2018/11/15 上午10:04
 */

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		cmd    *exec.Cmd
		ctx    context.Context
		cancel context.CancelFunc
		resu   *result
	)

	resChan := make(chan *result)

	ctx, cancel = context.WithCancel(context.Background())

	go func() {
		var (
			err    error
			output []byte
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")
		output, err = cmd.CombinedOutput()
		resChan <- &result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(3 * time.Second)

	cancel()

	resu = <-resChan
	fmt.Println(resu.err, string(resu.output))
}
