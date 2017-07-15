package main

import (
	"fmt"
	"strings"
	"time"
)

var LoadedPlugins map[string]*Plugin
var commands map[string]func(string)

type RunCmd struct {
	Cmd  string
	Args []string
}

func Runner() {
	for _, p := range LoadedPlugins {
		go p.Run()
	}
}

// START OMIT
type Plugin struct {
	Name     string
	RunChan  chan RunCmd
	ErrChan  chan error
	DoneChan chan int
	Commands map[string]func(string)
}

func (p *Plugin) Run() {
	for {
		select {
		case runcmd := <-p.RunChan:
			cmd := p.Commands[runcmd.Cmd]
			foo := strings.Join(runcmd.Args, ",")
			cmd(foo)
		case <-p.DoneChan:
			println("done")
			return
		}
	}
}

// END OMIT

func New(name string) *Plugin {
	p := Plugin{Name: name}
	p.RunChan = make(chan RunCmd)
	p.ErrChan = make(chan error)
	p.DoneChan = make(chan int)
	p.Commands = commands
	LoadedPlugins[name] = &p
	return &p
}

// START FOMIT
// package hello-plugin
func hello(s string) {
	fmt.Println("Hello,", s)
}

// main application code
func main() {
	commands = map[string]func(string){"hello": hello}
	LoadedPlugins = make(map[string]*Plugin)
	plug := New("hello")
	fmt.Println("New Plugin:", plug)
	Runner()

	plug.RunChan <- RunCmd{"hello", []string{"GopherCon"}}
	plug.DoneChan <- 1
	time.Sleep(time.Millisecond * 50)
}

// END FOMIT
