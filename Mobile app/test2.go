package main

import (
	"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var (
	images   *glutil.Images
	buf      gl.Buffer
	farge = 0
)

func main() {
	app.Main(func(a app.App) {
		var glctx gl.Context
		var sz size.Event
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ = e.DrawContext.(gl.Context)
					onStart(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					onStop(glctx)
					glctx = nil
				}
			case size.Event:
				sz = e

			case paint.Event:
				if glctx == nil || e.External {
					continue
				}
				onPaint(glctx, sz)
				a.Publish()
				a.Send(paint.Event{})

			case touch.Event:
				if (e.X != 0) && (e.Y != 0) {
					farge++
				}
			}
		}
	})
}

func onStart(glctx gl.Context) {
	var err error

	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}
	buf = glctx.CreateBuffer()
	glctx.BindBuffer(gl.ARRAY_BUFFER, buf)
	images = glutil.NewImages(glctx)
}

func onStop(glctx gl.Context) {
	glctx.DeleteBuffer(buf)
	images.Release()
}

func onPaint(glctx gl.Context, sz size.Event) {
	if farge == 0 {
		glctx.ClearColor(1, 0, 0, 1)
	} else if farge == 1 {
		glctx.ClearColor(0, 1, 0, 1)
	} else if farge == 2 {
		glctx.ClearColor(0, 0, 1, 1)
	} else {
		farge = 0
	}
	glctx.Clear(gl.COLOR_BUFFER_BIT)
}
