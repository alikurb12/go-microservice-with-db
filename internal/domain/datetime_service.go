package domain

type DateTimeService struct {
	clock Clock
}

func (dt DateTimeService) CurrentUnixTimestamp() int64 {
	return dt.clock.NowUnix()
}

func NewDateTimeService(clock Clock) *DateTimeService {
	return &DateTimeService{clock: clock}
}
