package v3

type Param map[string]interface{}

func NewParam() Param {
	return make(Param)
}

func (p Param) Add(key string, value interface{}) {
	p[key] = value
}

func (p Param) Delete(key string) (ok bool) {
	_, ok = p[key]
	delete(p, key)
	return
}

func (p Param) Get(key string) (value interface{}, ok bool) {
	value, ok = p[key]
	return
}

func (p Param) Range(fn func(key string, value interface{})) {
	for k, v := range p {
		fn(k, v)
	}
}

func (p Param) Check(fn func(key string, value interface{}) bool) (ok bool) {
	ok = true
	for k, v := range p {
		if fn(k, v) == false {
			ok = false
			break
		}
	}
	return
}
