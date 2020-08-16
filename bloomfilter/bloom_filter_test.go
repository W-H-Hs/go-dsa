package bloomfilter

import (
	"fmt"
	"testing"
)

var bf *BloomFilter

func init() {
	bf, _ = NewBloomFilter(1_0000_0000, 0.01)
}

func TestBloomFilter(t *testing.T) {
	a := 1
	b := 2
	c := 3
	d := 4
	err := bf.Put(&a)
	if err != nil {
		panic(err)
	}
	bf.Put(&b)
	bf.Put(&c)
	bf.Put(&d)

	isContain, err := bf.IsContain(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(isContain)
	isContain, err = bf.IsContain(&b)
	fmt.Println(isContain)
	isContain, err = bf.IsContain(&c)
	fmt.Println(isContain)
	isContain, err = bf.IsContain(&d)
	fmt.Println(isContain)
	e := 5
	f := 6
	g := 7
	isContain, err = bf.IsContain(&e)
	fmt.Println(isContain)
	isContain, err = bf.IsContain(&f)
	fmt.Println(isContain)
	isContain, err = bf.IsContain(&g)
	fmt.Println(isContain)
}
