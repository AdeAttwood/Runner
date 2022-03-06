// Copyright 2022 Practically.io All rights reserved
//
// Use of this source is governed by a BSD-style
// licence that can be found in the LICENCE file or at
// https://www.practically.io/copyright
//
// Console spinner inspired by janeczku/go-spinner.
// https://github.com/janeczku/go-spinner/blob/master/spinner.go
package console

import (
	"io"
	"sync"
	"time"
)

const (
	// 150ms per frame
	DEFAULT_FRAME_RATE = time.Millisecond * 150
)

var DefaultCharset = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

type Spinner struct {
	sync.Mutex
	Title     string
	Charset   []string
	FrameRate time.Duration
	runChan   chan struct{}
	stopOnce  sync.Once
	Output    io.Writer
	NoTty     bool
	frame     int
}

// create spinner object
func NewSpinner(title string) *Spinner {
	spinner := &Spinner{
		Title:     title,
		Charset:   DefaultCharset,
		FrameRate: DEFAULT_FRAME_RATE,
		runChan:   make(chan struct{}),
		frame:     0,
	}

	return spinner
}

// start spinner
func (sp *Spinner) Start() *Spinner {
	go sp.writer()
	return sp
}

func (spinner *Spinner) Success() {
	spinner.stop()

	ReplaceLine("✔ " + spinner.Title)
	WriteLine("")
}

func (spinner *Spinner) Error() {
	spinner.stop()

	ReplaceLine("✘ " + spinner.Title)
	WriteLine("")
}

// stop and clear the spinner
func (sp *Spinner) stop() {
	//prevent multiple calls
	sp.stopOnce.Do(func() {
		close(sp.runChan)
	})
}

// spinner animation
func (sp *Spinner) animate() {
	out := sp.Charset[sp.frame] + " " + sp.Title
	if sp.frame == len(sp.Charset)-1 {
		sp.frame = 0
	} else {
		sp.frame++
	}

	ReplaceLine(out)
	time.Sleep(sp.FrameRate)
}

// write out spinner animation until runChan is closed
func (sp *Spinner) writer() {
	sp.animate()
	for {
		select {
		case <-sp.runChan:
			return
		default:
			sp.animate()
		}
	}
}
