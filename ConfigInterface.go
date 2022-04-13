package goAws

type SetConfig interface {
	SetConfigByDefault() error
	SetConfigByProfile() error
	SetConfigByKey() error
}
