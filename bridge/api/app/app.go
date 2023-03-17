package app

import (
	"tractor.dev/apptron/bridge/api/menu"
	"tractor.dev/apptron/bridge/resource"
)

var Module *module

type module struct{}

func init() {
	Module = &module{}
}

type Options struct {
	Identifier string
	Agent      bool // app should not terminate when last window closes
	Accessory  bool // app should not be task switchable
}

func (m *module) Menu() *menu.Menu {
	return Menu()
}

func (m *module) SetMenu(handle resource.Handle) error {
	mm, err := menu.Get(handle)
	if err != nil {
		return err
	}
	SetMenu(mm)
	return nil
}

func (m *module) NewIndicator(icon []byte, items []menu.Item) error {
	NewIndicator(icon, items)
	return nil
}

func (m *module) Run(options Options) error {
	if options.Identifier == "" {
		options.Identifier = "com.progrium.Apptron"
	}
	return Run(options)
}
