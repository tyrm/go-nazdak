package serger

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/juju/loggo"
	"github.com/tyrm/go-gfxos"
)

type Serger struct {
	Panels []*gfxos.Matrix
	PanelW int
	PanelH int
}

var logger *loggo.Logger

func init() {
	newLogger := loggo.GetLogger("serger")
	logger = &newLogger

	logger.Infof("Serger Module Initlized")
}

func New(ports []string, w int, h int, ) (*Serger, error) {
	// Create new Matrix
	newSerger := Serger{
		PanelW: w,
		PanelH: h,
	}
	newSerger.Panels = []*gfxos.Matrix{}

	for _, portName := range ports {
		// Set up options.
		options := serial.OpenOptions{
			PortName:        portName,
			BaudRate:        2000000,
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		}

		// Open the port.
		port, err := serial.Open(options)
		if err != nil {
			fmt.Print("Error: ", err)

			// Close Any Opened Ports
			newSerger.Close()
			return nil, err
		}

		panel, err := gfxos.Open(&port)
		if err != nil {
			fmt.Print("Error: ", err)

			// Close Any Opened Ports
			newSerger.Close()
			return nil, err
		}

		newSerger.Panels = append(newSerger.Panels, panel)
	}

	return &newSerger, nil
}

func (s *Serger) DrawPixel(x int, y int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		if isInBounds(x, y, s.PanelW, s.PanelH, offset) {
			offset_x := x - (s.PanelW * offset)

			err := panel.DrawPixel(offset_x, y, r, g, b)
			if err != nil {
				logger.Errorf("DrawPixel: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) SetRotation(r int) error {
	for _, panel := range s.Panels {
		err := panel.SetRotation(r)
		if err != nil {
			logger.Errorf("SetRotation: %s", err)
		}
	}

	return nil
}

func (s *Serger) InvertDisplay(i int) error {
	for _, panel := range s.Panels {
		err := panel.InvertDisplay(i)
		if err != nil {
			logger.Errorf("InvertDisplay: %s", err)
		}
	}

	return nil
}

func (s *Serger) DrawFastVLine(x int, y int, h int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		left := x
		right := x
		top := y
		bottom := y + h

		logger.Tracef("DrawFastVLine left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x := x - (s.PanelW * offset)

			err := panel.DrawFastVLine(offset_x, y, h, r, g, b)
			if err != nil {
				logger.Errorf("DrawFastVLine: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) DrawFastHLine(x int, y int, w int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		left := x
		right := x + w
		top := y
		bottom := y

		logger.Tracef("DrawFastHLine left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x := x - (s.PanelW * offset)

			err := panel.DrawFastHLine(offset_x, y, w, r, g, b)
			if err != nil {
				logger.Errorf("DrawFastHLine: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) FillRect(x int, y int, w int, h int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		left := x
		right := x + w
		top := y
		bottom := y + h

		logger.Tracef("DrawFastHLine left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {

			offset_x := x - (s.PanelW * offset)

			err := panel.FillRect(offset_x, y, w, h, r, g, b)
			if err != nil {
				logger.Errorf("DrawFastHLine: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) FillScreen(r int, g int, b int) error {
	for _, panel := range s.Panels {
		err := panel.FillScreen(r, g, b)
		if err != nil {
			logger.Errorf("FillScreen: %s", err)
		}
	}

	return nil
}

func (s *Serger) DrawLine(x0 int, y0 int, x1 int, y1 int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		var left int
		var right int
		var top int
		var bottom int

		if x0 < x1 {
			left = x0
			right = x1
		} else {
			right = x1
			left = x0
		}

		if y0 < y1 {
			top = y0
			bottom = y1
		} else {
			top = y1
			bottom = y0
		}

		logger.Tracef("DrawLine left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x0 := x0 - (s.PanelW * offset)
			offset_x1 := x1 - (s.PanelW * offset)

			err := panel.DrawLine(offset_x0, y0, offset_x1, y1, r, g, b)
			if err != nil {
				logger.Errorf("DrawLine: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) DrawRect(x int, y int, w int, h int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		left := x
		right := x + w
		top := y
		bottom := y + h

		logger.Tracef("DrawRect left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {

			offset_x := x - (s.PanelW * offset)

			err := panel.DrawRect(offset_x, y, w, h, r, g, b)
			if err != nil {
				logger.Errorf("DrawRect: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) DrawCircle(x int, y int, rad int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		left := x - rad
		right := x + rad
		top := y - rad
		bottom := y + rad

		logger.Tracef("DrawCircle left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x := x - (s.PanelW * offset)

			err := panel.DrawCircle(offset_x, y, rad, r, g, b)
			if err != nil {
				logger.Errorf("DrawCircle: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) FillCircle(x int, y int, rad int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		left := x - rad
		right := x + rad
		top := y - rad
		bottom := y + rad

		logger.Tracef("FillCircle left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x := x - (s.PanelW * offset)

			err := panel.FillCircle(offset_x, y, rad, r, g, b)
			if err != nil {
				logger.Errorf("FillCircle: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) DrawTriangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		var left int
		var right int
		var top int
		var bottom int

		if x0 < x1 && x0 < x2 {
			left = x0
		} else if x1 < x2 {
			left = x1
		} else {
			left = x2
		}

		if x0 > x1 && x0 > x2 {
			right = x0
		} else if x1 > x2 {
			right = x1
		} else {
			right = x2
		}

		if y0 < y1 && y0 < y2 {
			top = y0
		} else if y1 < y2 {
			top = y1
		} else {
			top = y2
		}

		if y0 > y1 && y0 > y2 {
			bottom = y0
		} else if y1 > y2 {
			bottom = y1
		} else {
			bottom = y2
		}

		logger.Tracef("DrawTriangle left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x0 := x0 - (s.PanelW * offset)
			offset_x1 := x1 - (s.PanelW * offset)
			offset_x2 := x2 - (s.PanelW * offset)

			logger.Tracef("FillTriangle offset 0[%d] 1[%d] 2[%d]", offset_x0, offset_x1, offset_x2)

			err := panel.DrawTriangle(offset_x0, y0, offset_x1, y1, offset_x2, y2, r, g, b)
			if err != nil {
				logger.Errorf("DrawTriangle: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) FillTriangle(x0 int, y0 int, x1 int, y1 int, x2 int, y2 int, r int, g int, b int) error {
	for offset, panel := range s.Panels {
		var left int
		var right int
		var top int
		var bottom int

		if x0 < x1 || x0 < x2 {
			left = x0
		} else if x1 < x2 {
			left = x1
		} else {
			left = x2
		}

		if x0 > x1 || x0 > x2 {
			right = x0
		} else if x1 > x2 {
			right = x1
		} else {
			right = x2
		}

		if y0 < y1 || y0 < y2 {
			top = y0
		} else if y1 < y2 {
			top = y1
		} else {
			top = y2
		}

		if y0 > y1 || y0 > y2 {
			bottom = y0
		} else if y1 > y2 {
			bottom = y1
		} else {
			bottom = y2
		}

		logger.Tracef("FillTriangle left: %d right: %d top: %d bottom: %d", left, right, top, bottom)

		if rectIsInBounds(left, right, top, bottom, s.PanelW, s.PanelH, offset) {
			offset_x0 := x0 - (s.PanelW * offset)
			offset_x1 := x1 - (s.PanelW * offset)
			offset_x2 := x2 - (s.PanelW * offset)

			logger.Tracef("FillTriangle offset 0[%d] 1[%d] 2[%d]", offset_x0, offset_x1, offset_x2)

			err := panel.FillTriangle(offset_x0, y0, offset_x1, y1, offset_x2, y2, r, g, b)
			if err != nil {
				logger.Errorf("FillTriangle: %s", err)
			}
		}
	}

	return nil
}

func (s *Serger) DrawChar(x int, y int, fr int, fg int, fb int, br int, bg int, bb int, size int, char int) error {
	panic("implement me")
}

func (s *Serger) SetCursor(x int, y int) error {
	for offset, panel := range s.Panels {
		if isInBounds(x, y, s.PanelW, s.PanelH, offset) {
			offset_x := x - (s.PanelW * offset)

			err := panel.SetCursor(offset_x, y)
			if err != nil {
				logger.Errorf("SetCursor: %s", err)
			}

		}
	}

	return nil
}

func (s *Serger) SetTextColor(r int, g int, b int) error {
	panic("implement me")
}

func (s *Serger) SetTextColorBG(fr int, fg int, fb int, br int, bg int, bb int) error {
	panic("implement me")
}

func (s *Serger) SetTextSize(sz int) error {
	panic("implement me")
}

func (s *Serger) SetTextWrap(w int) error {
	panic("implement me")
}

func (s *Serger) CP437(c int) error {
	panic("implement me")
}

func (s *Serger) Print(str string) error {
	panic("implement me")
}

func (s *Serger) PrintLn(str string) error {
	panic("implement me")
}

func (s *Serger) SetFont(f int) error {
	panic("implement me")
}

func (s *Serger) Close() {
	for _, panel := range s.Panels {
		panel.Close()
	}
}

func isInBounds(x int, y int, w int, h int, offset int) bool {
	leftBoundary := w * offset
	rightBoundary := (w * (offset + 1)) - 1
	topBoundary := 0
	bottomBoundary := h - 1

	if leftBoundary <= x && x <= rightBoundary && topBoundary <= y && y <= bottomBoundary {
		return true
	}

	return false
}

func rectIsInBounds(left int, right int, top int, bottom int, panelW int, panelH int, offset int) bool {
	return isInBounds(left, top, panelW, panelH, offset) ||
		isInBounds(right, top, panelW, panelH, offset) ||
		isInBounds(left, bottom, panelW, panelH, offset) ||
		isInBounds(right, bottom, panelW, panelH, offset)
}

