package textfield_test

import (
	"fmt"
	"log"

	"github.com/dave/material/internal/mdctest"
	"github.com/dave/material/textfield"
	"github.com/gopherjs/gopherjs/js"
)

func Example() {
	// Create a new instance of a material textfield component.
	c := &textfield.TF{}

	// Set up a DOM HTMLElement suitable for a textfield.
	js.Global.Get("document").Get("body").Set("innerHTML",
		mdctest.HTML(c.Component().Type.MDCClassName))
	rootElem := js.Global.Get("document").Get("body").Get("firstElementChild")

	// Start the component, which associates it with an HTMLElement.
	err := c.Start(rootElem)
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n",
			c.Component().Type, err)
	}

	printStatus(c)
	printState(c)
	c.Required = false
	c.HelperText = "Must be at least 8 characters."
	c.Value = "longerpassword"
	c.Disabled = true
	printState(c)

	err = c.Stop()
	if err != nil {
		log.Fatalf("Unable to stop component %s: %v\n",
			c.Component().Type, err)
	}

	// Output:
	// MDCTextField

	// Disabled: false, Valid: false, Required: true
	// Value: , HelperText: undefined

	// Disabled: true, Valid: true, Required: false
	// Value: longerpassword, HelperText: Must be at least 8 characters.
}

func printStatus(c *textfield.TF) {
	fmt.Printf("%s\n", c.Component().Type)
}

func printState(c *textfield.TF) {
	fmt.Println()
	fmt.Printf("Disabled: %v, Valid: %v, Required: %v\n",
		c.Disabled, c.Valid, c.Required)
	fmt.Printf("Value: %v, HelperText: %v\n", c.Value, c.HelperText)
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}
