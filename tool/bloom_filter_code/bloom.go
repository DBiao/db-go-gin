package bloom_filter_code

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/spaolacci/murmur3"
	"math"
)

type BloomFilter struct {
	m uint64 //数组集合大小
	k uint32 //hash函数个数
	b *bitset.BitSet
}

// New m代表位图长度
// k代表hash函数的个数
func New(m uint64, k uint32) *BloomFilter {
	return &BloomFilter{
		m, k, bitset.New(uint(m)),
	}
}

// NewWithExpected 计算最佳的配置
// https://hur.st/bloomfilter/
// n代表最多个不同元素
// p代表假阳率
// m代表位图长度
// k代表hash函数的个数
func NewWithExpected(n uint, p float64) *BloomFilter {
	return New(uint64(math.Ceil(-1*float64(n)*math.Log(p)/math.Pow(math.Log(2), 2))), uint32(math.Ceil(math.Log(2)*float64(n)/float64(n))))
}

func (f *BloomFilter) Add(data []byte) {
	for i := uint32(0); i < f.k; i++ {
		f.b.Set(f.locate(data, i))
	}
}
func (f *BloomFilter) Exist(data []byte) bool {
	for i := uint32(0); i < f.k; i++ {
		if !f.b.Test(f.locate(data, i)) { //一个不存在就绝对不存在
			return false
		}
	}
	return true
}

func (f *BloomFilter) locate(data []byte, seed uint32) uint {
	return getHash(data, seed) % uint(f.m)
}

func getHash(data []byte, seed uint32) uint {
	m := murmur3.New64WithSeed(seed)
	_, _ = m.Write(data)
	return uint(m.Sum64())
}
