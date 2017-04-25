# aisola/log

A logging library because all other logging libraries are too complicated.

[![GoDoc Widget](https://godoc.org/github.com/aisola/log?status.svg)](https://godoc.org/github.com/aisola/log)
[![Code Climate](https://codeclimate.com/github/aisola/log/badges/gpa.svg)](https://codeclimate.com/github/aisola/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/aisola/log)](https://goreportcard.com/report/github.com/aisola/log)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

This comes from searching around for a not complicated logging library and
running into [this article](https://dave.cheney.net/2015/11/05/lets-talk-about-logging).

This library is intended to only have two levels and be very a uncomplicated logger for
go programs.

### Warning

Like any other random code you find on the internet, this package should not be relied
upon in important, production systems without thorough testing to ensure that it meets
your needs. I wrote this for handling my simple needs and I never expect that it will
really be useful for anybody else. It has no tests, and has barely been used.

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
