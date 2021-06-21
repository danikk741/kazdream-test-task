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
	hash  int
	elems []elem
}

type MapGo struct {
	length     int64
	buckets    []bucket
	bucketsNum int
}

func NewMap(bucketsNum int) *MapGo {
	buckets := make([]bucket, bucketsNum)
	for i := range buckets {
		buckets[i].hash = i
	}
	return &MapGo{
		bucketsNum: bucketsNum,
		buckets:    buckets,
		length:     0,
	}
}

func (m *MapGo) Add(k string.String) {
	for i, bucket := range m.buckets {
		h := hash(k.Bytes)
		if bucket.hash == h {
			ok := false
			for j, el := range m.buckets[i].elems {
				if strings.Equal(el.k, k) {
					m.buckets[i].elems[j].v++
					ok = true
					break
				}
			}
			if !ok {
				m.buckets[i].elems = append(m.buckets[i].elems, elem{
					k: k,
					v: 1,
				})
			}
		}
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
	for i, bucket := range m.buckets {
		h := hash(k.Bytes)
		if bucket.hash == h {
			for j, el := range m.buckets[i].elems {
				if strings.Equal(el.k, k) {
					res = true
					m.buckets[i].elems[j] = m.buckets[i].elems[len(m.buckets[i].elems)-1]
					m.buckets[i].elems = m.buckets[i].elems[:len(m.buckets[i].elems)-1]
					break
				}
			}
		}
	}
	return res
}

func (m *MapGo) Lookup(k string.String) (int, bool) {
	var ok bool
	var val int
	for i, bucket := range m.buckets {
		h := hash(k.Bytes)
		if bucket.hash == h {
			for _, el := range m.buckets[i].elems {
				if strings.Equal(el.k, k) {
					ok = true
					val = el.v
					break
				}
			}
		}
	}
	return val, ok
}
