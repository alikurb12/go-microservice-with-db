package domain

type Clock interface {
	NowUnix() int64
}