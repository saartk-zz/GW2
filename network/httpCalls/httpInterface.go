package httpCalls

type HttpCaller interface{
	Get(path string) ([]byte,error)
}
