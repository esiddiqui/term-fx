package progress

import (
	"fmt"
	"time"

	"github.com/esiddiqui/tfx/cursor"
)

var (
	WaiterFrames1 []string = []string{"/", "-", "\\", "|", "-", "\\", "|"}
	WaiterFrames2 []string = []string{".", "..", "...", "....", " ...", "  ..", "   .", "", "   .", "  ..", " ...", "....", "... ", "..  ", ".  ", ""}
	WaiterFrames3 []string = []string{"¯", "¯\\", "¯\\_", "¯\\_(", "¯\\_(ツ", "¯\\_(ツ)", "¯\\_(ツ)_", "¯\\_(ツ)_/", "¯\\_(ツ)_/¯", "          "}
)

// a simple wrapper func type alias that takes in a `chan any` to indicate the worker finished processing & produced a result..
type SimpleWaiterFnc func(chan any)

// a simple post processor func type alias, which takes an argument of type any to process
type SimpleWaiterPostFnc func(any)

// SimpleWaiter displays a progress waiter text animation on the screen until the
// supplied worker aysnchronously completes the job & return a value over `channel any`
type SimpleWaiter struct {
	Fps    int
	Frames []string
}

// NewSimpleWaiter creates & returns a SimpleWaiter
func NewSimpleWaiter(fps int) SimpleWaiter {
	if fps > 1000 || fps < 0 {
		fps = 20
	}
	return SimpleWaiter{
		Fps:    fps,
		Frames: WaiterFrames1,
	}
}

// WaitWithPost calls the wrapped works `fn` & starts a visual Waiter using underlying Wait()
// method. After the method returns, the results of the underlying worker call returned via the
// internal quit channel is passed on to the post fnc for any post processing.
func (w SimpleWaiter) WaitWithPost(fn SimpleWaiterFnc, post SimpleWaiterPostFnc) error {
	rc, err := w.Wait(fn)
	if err != nil {
		return err
	}
	post(rc)
	return nil
}

// Wait calls the wrapped worker `fn` & starts a visual Wwaiter. When the fn is done
// it passes the return data received over the channel to the post fn for processing
func (w SimpleWaiter) Wait(fn SimpleWaiterFnc) (any, error) {

	var idx int
	duration := time.Duration(1000/w.Fps) * time.Millisecond

	// build a ticker that ticks every xms based on
	ticker := time.After(duration)
	rc := make(chan any) // make a quick channel

	// start the waiter func as a go routing
	cursor.Off()
	go fn(rc)

	for {

		select {
		case val := <-rc:
			cursor.ClearToStartOfLine() // clear this ln
			cursor.Col(1)               // move cursor to beginning of ln
			cursor.On()                 // set cursor visible
			return val, nil

		case <-ticker:
			cursor.Col(1)                   // move cursor to beginning of ln
			fmt.Print(w.Frames[idx])        // paint the next frame
			idx = (idx + 1) % len(w.Frames) // move idx
			ticker = time.After(duration)   // set timer
		}
	}
}
