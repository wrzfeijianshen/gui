package mixins

import (
	"github.com/wrzfeijianshen/gui/gxui"
	"github.com/wrzfeijianshen/gui/gxui/math"
	"github.com/wrzfeijianshen/gui/gxui/mixins/outer"
	"github.com/wrzfeijianshen/gui/gxui/mixins/parts"
)

type WindowOuter interface {
	gxui.Window
	outer.Attachable
	outer.IsVisibler
	outer.LayoutChildren
	outer.PaintChilder
	outer.Painter
	outer.Parenter
	outer.Sized
}

type Window struct {
	parts.Attachable
	parts.BackgroundBorderPainter
	parts.Container
	parts.Paddable
	parts.PaintChildren

	driver             gxui.Driver
	outer              WindowOuter
	viewport           gxui.Viewport
	windowedSize       math.Size
	mouseController    *gxui.MouseController
	keyboardController *gxui.KeyboardController
	focusController    *gxui.FocusController
	layoutPending      bool
	drawPending        bool
	updatePending      bool
	onClose            gxui.Event // Raised by viewport
	onResize           gxui.Event // Raised by viewport
	onMouseMove        gxui.Event // Raised by viewport
	onMouseEnter       gxui.Event // Raised by viewport
	onMouseExit        gxui.Event // Raised by viewport
	onMouseDown        gxui.Event // Raised by viewport
	onMouseUp          gxui.Event // Raised by viewport
	onMouseScroll      gxui.Event // Raised by viewport
	onKeyDown          gxui.Event // Raised by viewport
	onKeyUp            gxui.Event // Raised by viewport
	onKeyRepeat        gxui.Event // Raised by viewport
	onKeyStroke        gxui.Event // Raised by viewport

	onClick       gxui.Event // Raised by MouseController
	onDoubleClick gxui.Event // Raised by MouseController

	viewportSubscriptions []gxui.EventSubscription
}
