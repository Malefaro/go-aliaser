# Simple aliaser 

## About 

Simple app to add aliases on ubuntu. 

## How does it work

This app create two files: 

* `$HOME/.go-aliaser`
* `$HOME/.go-aliaser-map`


And add your `$HOME/.profile` file to include `$HOME/.go-aliaser`, which contains aliases

`$HOME/.go-aliaser-map` - JSON file with aliases

## Install

go get github.com/malefaro/go-aliaser

## Usage

#### Create

`go-aliaser create <name>="<command>"`

#### Remove 

`go-aliaser rm <name>`

