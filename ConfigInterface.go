package main

type SetConfig interface {
	SetConfigByDefault() error
	SetConfigByProfile() error
	SetConfigByKey() error
}
