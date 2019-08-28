package metas

type SqlMate interface {
	GetMetaStr() (string, error)
}
