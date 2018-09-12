package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "./videos/"
)

//控制
type controlChan chan string

//数据
type dataChan chan interface{}

type fn func(dc dataChan) error
