package models

type Cache interface {
	Start() error
	Stop() error
	Set(key, value string)
	Get(key string) (string, bool)
	Update(key, value string) (string, bool)
	Delete(key string)
}
