package filter

import "testing"

func TestBloomFilter_Contains(t *testing.T) {

	b := NewBloomFilter(0)

	g := b.NewGenerator()

	g.Add([]byte("123456"))

	t.Log(b)

}
