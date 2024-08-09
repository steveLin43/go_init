package main

type Value interface {
	String() string
	Set(string) error
}
