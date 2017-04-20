# aisola/log

A logging library because all other logging libraries are too complicated.

This comes from searching around for a not complicated logging library and
running into [this article](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).

This library is intended to only have two levels and be very a uncomplicated logger for
go programs.

## Installing

    go get -u -v github.com/aisola/log

## Example

    package main

    import "github.com/aisola/log"

    func main() {
    	 log.Info("The program is running")
         log.Infof("This is the %s message", "2nd")
         log.Debug("I'll only be written to the log if you've told the logger to.")
         log.Debugf("Hello, %s!", "world")
    }
