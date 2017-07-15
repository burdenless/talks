package main

import "fmt"

var LoadedPlugins map[string]Plugin

func init() {
	LoadedPlugins = make(map[string]Plugin)
}

type Plugin interface {
	Run(string) error
	Register()
}

// START OMIT
type Hello struct {
	Name string
}

func (h Hello) Run(s string) error {
	fmt.Println("Hello,", s)
	return nil
}

func (h Hello) Register() {
	LoadedPlugins[h.Name] = h
}

func main() {
	h := Hello{Name: "hello"}
	Plugin.Register(h)
	fmt.Println("Loaded Plugins:", len(LoadedPlugins))
	plug := LoadedPlugins["hello"]
	Plugin.Run(plug, "GopherCon")
}

// END OMIT
