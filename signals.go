package signals

import (
	"os"
	"os/signal"
	"syscall"
)

func USR1(fn func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGUSR1)
	go func() {
		for {
			select {
			case <-ch:
				fn()
			}
		}
	}()
}

func USR2(fn func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGUSR2)
	go func() {
		for {
			select {
			case <-ch:
				fn()
			}
		}
	}()
}

func INT(fn func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	go func() {
		for {
			select {
			case <-ch:
				fn()
			}
		}
	}()
}

func TERM(fn func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		for {
			select {
			case <-ch:
				fn()
			}
		}
	}()
}

func HUP(fn func()) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP)
	go func() {
		for {
			select {
			case <-ch:
				fn()
			}
		}
	}()
}
