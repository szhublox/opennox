package sdl

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"

	"nox/v1/client/seat"
	"nox/v1/common/env"
	"nox/v1/common/log"
	"nox/v1/common/types"
)

var (
	Log       = log.New("sdl")
	debugGpad = os.Getenv("NOX_DEBUG_GPAD") == "true"
)

var _ seat.Seat = &Window{}

// New creates a new SDL window which implements a Seat interface.
func New(title string, sz types.Size) (*Window, error) {
	// TODO: if we ever decide to use multiple windows, this will need to be moved elsewhere; same for sdl.Quit
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER | sdl.INIT_GAMECONTROLLER); err != nil {
		return nil, fmt.Errorf("SDL Initialization failed: %w", err)
	}
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	win, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(sz.W), int32(sz.H), sdl.WINDOW_RESIZABLE)
	if err != nil {
		sdl.Quit()
		return nil, fmt.Errorf("SDL Window creation failed: %w", err)
	}
	ren, err := sdl.CreateRenderer(win, 0, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		win.Destroy()
		sdl.Quit()
		return nil, fmt.Errorf("SDL cannot create renderer: %w", err)
	}
	h := &Window{win: win, ren: ren, prevSz: sz}
	h.SetScreenMode(seat.Windowed)
	return h, nil
}

type Window struct {
	win      *sdl.Window
	ren      *sdl.Renderer
	prevPos  types.Point
	prevSz   types.Size
	mode     seat.ScreenMode
	rel      bool
	mpos     types.Point
	onResize []func(sz types.Size)
	onInput  []func(ev seat.InputEvent)
}

func (win *Window) Close() error {
	if win.win == nil {
		return nil
	}
	if win.ren != nil {
		_ = win.ren.Destroy()
		win.ren = nil
	}
	err := win.win.Destroy()
	win.win = nil
	win.onResize = nil
	win.onInput = nil
	sdl.Quit()
	return err
}

func (win *Window) OnInput(fnc func(ev seat.InputEvent)) {
	win.onInput = append(win.onInput, fnc)
}

func (win *Window) SetTextInput(enable bool) {
	if enable {
		sdl.StartTextInput()
	} else {
		sdl.StopTextInput()
	}
}

func (win *Window) ScreenSize() types.Size {
	w, h := win.win.GetSize()
	return types.Size{
		W: int(w), H: int(h),
	}
}

func (win *Window) screenPos() types.Point {
	x, y := win.win.GetPosition()
	return types.Point{
		X: int(x), Y: int(y),
	}
}

func (win *Window) displayRect() sdl.Rect {
	disp, err := win.win.GetDisplayIndex()
	if err != nil {
		Log.Println("can't get display index: ", err)
		return sdl.Rect{}
	}
	rect, err := sdl.GetDisplayBounds(disp)
	if err != nil {
		Log.Println("can't get display bounds: ", err)
		return sdl.Rect{}
	}
	return rect
}

func (win *Window) ScreenMaxSize() types.Size {
	rect := win.displayRect()
	return types.Size{
		W: int(rect.W), H: int(rect.H),
	}
}

func (win *Window) setSize(sz types.Size) {
	Log.Printf("window size: %dx%d", sz.W, sz.H)
	win.win.SetSize(int32(sz.W), int32(sz.H))
}

func (win *Window) center() {
	win.win.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
}

func (win *Window) ResizeScreen(sz types.Size) {
	if win.mode != seat.Windowed {
		return
	}
	win.setSize(sz)
	win.prevSz = sz
}

func (win *Window) setRelative(rel bool) {
	if win.rel == rel {
		return
	}
	win.rel = rel
	win.win.SetGrab(rel)
	sdl.SetRelativeMouseMode(rel)
}

func (win *Window) SetScreenMode(mode seat.ScreenMode) {
	if win.mode == mode {
		return
	}
	if win.mode == seat.Windowed {
		// preserve size and pos, so we can restore them later
		win.prevSz = win.ScreenSize()
		win.prevPos = win.screenPos()
	}
	switch mode {
	case seat.Windowed:
		win.win.SetFullscreen(0)
		win.win.SetResizable(true)
		win.win.SetBordered(true)
		win.setSize(win.prevSz)
		if win.prevPos != (types.Point{}) {
			win.win.SetPosition(int32(win.prevPos.X), int32(win.prevPos.Y))
		} else {
			win.center()
		}
		if env.IsDevMode() {
			sdl.ShowCursor(sdl.ENABLE)
		} else {
			sdl.ShowCursor(sdl.DISABLE)
		}
		win.setRelative(false)
	case seat.Fullscreen:
		win.win.SetResizable(false)
		win.win.SetBordered(false)
		win.setSize(win.ScreenMaxSize())
		win.win.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
		sdl.ShowCursor(sdl.DISABLE)
		win.setRelative(true)
	case seat.Borderless:
		win.win.SetFullscreen(0)
		win.win.SetResizable(false)
		win.win.SetBordered(true)
		win.setSize(win.ScreenMaxSize())
		win.center()
		sdl.ShowCursor(sdl.DISABLE)
		win.setRelative(false)
	}
	win.mode = mode
}

// SetGamma sets screen gamma parameter.
func (win *Window) SetGamma(v float32) {
	var ramp [256]uint16
	sdl.CalculateGammaRamp(v, &ramp)
	if err := win.win.SetGammaRamp(&ramp, &ramp, &ramp); err != nil {
		Log.Printf("cannot set gamma: %v", err)
	}
}

func (win *Window) OnScreenResize(fnc func(sz types.Size)) {
	win.onResize = append(win.onResize, fnc)
}

func (win *Window) NewSurface(sz types.Size) seat.Surface {
	tex, err := win.ren.CreateTexture(sdl.PIXELFORMAT_RGB555, sdl.TEXTUREACCESS_STREAMING, int32(sz.W), int32(sz.H))
	if err != nil {
		err = fmt.Errorf("cannot create surface: %w", err)
		Log.Print(err)
		panic(err)
	}
	return &Surface{win: win, sz: sz, tex: tex}
}

func (win *Window) InputTick() {
	for {
		switch ev := sdl.PollEvent().(type) {
		case nil:
			// no more events
			return
		case *sdl.TextEditingEvent:
			win.processTextEditingEvent(ev)
		case *sdl.TextInputEvent:
			win.processTextInputEvent(ev)
		case *sdl.KeyboardEvent:
			win.processKeyboardEvent(ev)
		case *sdl.MouseButtonEvent:
			win.processMouseButtonEvent(ev)
		case *sdl.MouseMotionEvent:
			win.processMotionEvent(ev)
		case *sdl.MouseWheelEvent:
			win.processWheelEvent(ev)
		case *sdl.ControllerAxisEvent:
			if debugGpad {
				Log.Printf("SDL event: SDL_CONTROLLERAXISMOTION (%x): joy=%d, axis=%d, val=%d\n",
					ev.GetType(), ev.Which, ev.Axis, ev.Value)
			}
			win.processGamepadAxisEvent(ev)
		case *sdl.ControllerButtonEvent:
			if debugGpad {
				Log.Printf("SDL event: SDL_CONTROLLERBUTTON (%x): joy=%d, btn=%d, state=%d\n",
					ev.GetType(), ev.Which, ev.Button, ev.State)
			}
			win.processGamepadButtonEvent(ev)
		case *sdl.ControllerDeviceEvent:
			switch ev.GetType() {
			case sdl.CONTROLLERDEVICEADDED:
				if debugGpad {
					Log.Printf("SDL event: SDL_CONTROLLERDEVICEADDED (%x): joy=%d\n", ev.GetType(), ev.Which)
				}
				win.processGamepadDeviceEvent(ev)
			case sdl.CONTROLLERDEVICEREMOVED:
				if debugGpad {
					Log.Printf("SDL event: SDL_CONTROLLERDEVICEREMOVED (%x): joy=%d\n", ev.GetType(), ev.Which)
				}
				win.processGamepadDeviceEvent(ev)
			case sdl.CONTROLLERDEVICEREMAPPED:
				if debugGpad {
					Log.Printf("SDL event: SDL_CONTROLLERDEVICEREMAPPED (%x)\n", ev.GetType())
				}
			}
		case *sdl.WindowEvent:
			win.processWindowEvent(ev)
		case *sdl.QuitEvent:
			win.processQuitEvent(ev)
		}
		// TODO: touch events for WASM
	}
}

func (win *Window) inputEvent(ev seat.InputEvent) {
	for _, fnc := range win.onInput {
		fnc(ev)
	}
}

func (win *Window) processQuitEvent(ev *sdl.QuitEvent) {
	win.inputEvent(seat.WindowClosed)
}

func (win *Window) processWindowEvent(ev *sdl.WindowEvent) {
	switch ev.GetType() {
	case sdl.WINDOWEVENT_FOCUS_LOST:
		win.inputEvent(seat.WindowUnfocused)
	case sdl.WINDOWEVENT_FOCUS_GAINED:
		win.inputEvent(seat.WindowFocused)
	}
}

func (win *Window) processTextEditingEvent(ev *sdl.TextEditingEvent) {
	win.inputEvent(&seat.TextEditEvent{
		Text: ev.GetText(),
	})
}

func (win *Window) processTextInputEvent(ev *sdl.TextInputEvent) {
	win.inputEvent(&seat.TextInputEvent{
		Text: ev.GetText(),
	})
}

func (win *Window) processKeyboardEvent(ev *sdl.KeyboardEvent) {
	key := scanCodeToKeyNum[ev.Keysym.Scancode]
	win.inputEvent(&seat.KeyboardEvent{
		Key:     key,
		Pressed: ev.State == sdl.PRESSED,
	})
}

func (win *Window) processMouseButtonEvent(ev *sdl.MouseButtonEvent) {
	pressed := ev.State == sdl.PRESSED
	// TODO: handle focus, or move to other place
	//if pressed {
	//	h.iface.WindowEvent(WindowFocus)
	//}

	var button seat.MouseButton
	switch ev.Button {
	case sdl.BUTTON_LEFT:
		button = seat.MouseButtonLeft
	case sdl.BUTTON_RIGHT:
		button = seat.MouseButtonRight
	case sdl.BUTTON_MIDDLE:
		button = seat.MouseButtonMiddle
	default:
		return
	}
	win.inputEvent(&seat.MouseButtonEvent{
		Button:  button,
		Pressed: pressed,
	})
}

func (win *Window) processMotionEvent(ev *sdl.MouseMotionEvent) {
	win.inputEvent(&seat.MouseMoveEvent{
		Relative: win.rel,
		Pos:      types.Point{X: int(ev.X), Y: int(ev.Y)},
		Rel:      types.Point{X: int(ev.XRel), Y: int(ev.YRel)},
	})
}

func (win *Window) processWheelEvent(ev *sdl.MouseWheelEvent) {
	win.inputEvent(&seat.MouseWheelEvent{
		Wheel: int(ev.Y),
	})
}

func (win *Window) processGamepadButtonEvent(ev *sdl.ControllerButtonEvent) {
	// TODO: handle gamepads (again)
}

func (win *Window) processGamepadAxisEvent(ev *sdl.ControllerAxisEvent) {
	// TODO: handle gamepads (again)
}

func (win *Window) processGamepadDeviceEvent(ev *sdl.ControllerDeviceEvent) {
	// TODO: handle gamepads (again)
}
