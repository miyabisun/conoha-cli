package endpoints

type Endpoint struct {
	Url string
}

type Response struct {
	Body   []byte
	Status int
}

type Request struct {
	Key   string
	Value string
}
