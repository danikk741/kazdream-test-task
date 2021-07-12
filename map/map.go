package custom_map

import (
	string "kazdream-test-task/string"
	strings "kazdream-test-task/strings"
)

type elem struct {
	k string.String
	v int
}

type Elem struct {
	K string.String
	V int
}

type bucket struct {
	elems []elem
}

type MapGo struct {
	length     int64
	buckets    []bucket
	bucketsNum int
}

func NewMap(bucketsNum int) *MapGo {
	buckets := make([]bucket, bucketsNum)
	return &MapGo{
		bucketsNum: bucketsNum,
		buckets:    buckets,
		length:     0,
	}
}

func (m *MapGo) Add(k string.String) {
	h := hash(k.Bytes)
	ok := false
	for j, el := range m.buckets[h].elems {
		if strings.Equal(el.k, k) {
			m.buckets[h].elems[j].v++
			ok = true
			break
		}
	}
	if !ok {
		m.buckets[h].elems = append(m.buckets[h].elems, elem{
			k: k,
			v: 1,
		})
	}
}

func (m *MapGo) ToArray() []Elem {
	var arr []Elem
	for _, bucket := range m.buckets {
		for _, elem := range bucket.elems {
			arr = append(arr, Elem{
				K: elem.k,
				V: elem.v,
			})
		}
	}
	return arr
}

func (m *MapGo) Delete(k string.String) bool {
	var res bool
	h := hash(k.Bytes)
	for j, el := range m.buckets[h].elems {
		if strings.Equal(el.k, k) {
			res = true
			m.buckets[h].elems[j] = m.buckets[h].elems[len(m.buckets[h].elems)-1]
			m.buckets[h].elems = m.buckets[h].elems[:len(m.buckets[h].elems)-1]
			break
		}
	}
	return res
}

func (m *MapGo) Lookup(k string.String) (int, bool) {
	var ok bool
	var val int
	h := hash(k.Bytes)
	for _, el := range m.buckets[h].elems {
		if strings.Equal(el.k, k) {
			ok = true
			val = el.v
			break
		}
	}
	return val, ok
}
