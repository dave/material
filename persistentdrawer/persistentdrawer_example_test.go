package persistentdrawer_test

import (
	"fmt"
	"log"

	"agamigo.io/material/component/componenthtml"
	"agamigo.io/material/mdctest"
	"agamigo.io/material/persistentdrawer"
)

func Example() {
	// Create a new instance of a material persistentdrawer component.
	c, err := persistentdrawer.New()
	if err != nil {
		log.Fatalf("Unable to create component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)

	// Set up a DOM HTMLElement suitable for an persistentdrawer.
	mdctest.Dom.SetHTML("<html><body>" + componenthtml.HTML(c.CType()) +
		"</body></html>")

	// Start the component, which associates it with an HTMLElement.
	err = c.Start()
	if err != nil {
		log.Fatalf("Unable to start component %s: %v\n", c, err.Error())
	}
	fmt.Printf("%s\n", c)
	fmt.Printf("MDC Open: %v\n", c.GetObject().Get("open"))

	// Open the drawer
	c.Open = true
	fmt.Printf("MDC Open: %v\n", c.GetObject().Get("open"))

	// Output:
	// {"component":"MDCPersistentDrawer","status":"stopped"}
	// {"component":"MDCPersistentDrawer","status":"running"}
	// MDC Open: false
	// MDC Open: true
}

func init() {
	// We emulate a DOM here since tests run in NodeJS.
	// Not needed when running in a browser.
	err := mdctest.Init()
	if err != nil {
		log.Fatalf("Unable to setup test environment: %v", err)
	}
}