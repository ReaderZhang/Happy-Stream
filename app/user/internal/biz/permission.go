package biz

type Permission struct {
	Id       int64
	Url      string
	Children []Permission
}
