// slider implements a material slider component.
//
// See: https://material.io/components/web/catalog/input-controls/sliders/
package slider // import "github.com/dave/material/slider"

import (
	"agamigo.io/gojs"
	"github.com/dave/material/internal/base"
	"github.com/gopherjs/gopherjs/js"
)

// S is a material slider component.
type S struct {
	mdc *base.Component

	// The current value of the slider. Changing this will update the slider’s
	// value.
	Value float64 `js:"value"`

	// The minimum value the slider can have. Values set programmatically will
	// be clamped to this minimum value. Changing this property will update the
	// slider’s value if it is lower than the new minimum.
	Min float64 `js:"min"`

	// The maximum value a slider can have. Values set programmatically will be
	// clamped to this maximum value. Changing this property will update the
	// slider’s value if it is greater than the new maximum.
	Max float64 `js:"max"`

	// Specifies the increments at which a slider value can be set. Can be any
	// positive number, or 0 for no step. Changing this property will update the
	// slider’s value to be quantized along the new step increments.
	Step float64 `js:"step"`

	// Whether or not the slider is disabled.
	Disabled bool `js:"disabled"`
}

// Start initializes the component with an existing HTMLElement, rootElem. Start
// should only be used on a newly created component, or after calling Stop.
func (c *S) Start(rootElem *js.Object) error {
	return base.Start(c.Component(), rootElem)
}

// Stop removes the component's association with its HTMLElement and cleans up
// event listeners, etc.
func (c *S) Stop() error {
	return base.Stop(c.mdc)
}

// Component returns the component's underlying base.Component.
func (c *S) Component() *base.Component {
	if c.mdc == nil {
		c.mdc = &base.Component{}
		c.mdc.Type = base.ComponentType{
			MDCClassName:     "MDCSlider",
			MDCCamelCaseName: "slider",
		}
	}
	return c.mdc
}

// Layout recomputes the dimensions and re-lays out the component. This should
// be called if the dimensions of the slider itself or any of its parent
// elements change programmatically (it is called automatically on resize).
func (s *S) Layout() error {
	var err error
	gojs.CatchException(&err)
	s.mdc.Call("layout")
	return err
}

// TODO: Handle custom events
// - MDCSlider:input
// - MDCSlider:change
