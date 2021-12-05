package vars

import (
	"reflect"
	"strconv"
)

// TradeType 支付类型
type TradeType string

func (p TradeType) Valid() (ok bool) {
	ok = true
	switch p {
	case JSAPI:
	case H5:
	case APP:
	case Native:
	default:
		ok = false
	}
	return
}

// QueryType 订单查询类型
type QueryType string

func (q QueryType) Valid() bool {
	return q == QueryOutTradeNo || q == QueryTransactionId
}

type Kvs map[string]interface{}

func NewKvs() Kvs {
	return make(Kvs)
}

func (p Kvs) Add(key string, value interface{}) {
	p[key] = value
}

func (p Kvs) Delete(key string) (ok bool) {
	_, ok = p[key]
	delete(p, key)
	return
}

func (p Kvs) Get(key string) (value interface{}, ok bool) {
	value, ok = p[key]
	return
}

func (p Kvs) Range(fn func(key string, value interface{}) (breakOut bool)) {
	for k, v := range p {
		if fn(k, v) {
			break
		}
	}
}

func (p Kvs) GetString(key string) (string, error) {
	value, ok := p.Get(key)
	if !ok {
		return "", ErrNotExistKey
	}
	return reflect.ValueOf(value).String(), nil
}

func (p Kvs) GetInt64(key string) (n int64, err error) {
	value, ok := p.Get(key)
	if !ok {
		return 0, ErrNotExistKey
	}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Bool:
		if v.Bool() == true {
			n = 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n = v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n = int64(v.Uint())
	case reflect.String:
		n, err = strconv.ParseInt(v.String(), 10, 64)
	case reflect.Float32, reflect.Float64:
		n = int64(v.Float())
	default:
		err = ErrConvert
	}
	return
}
