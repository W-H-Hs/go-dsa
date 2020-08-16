package bloomfilter

import (
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/sha3"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

type BloomFilter struct {
	bitSize     int      // bitSize 为二进制向量的长度。
	bits        []uint64 // bits 为二进制向量，将uint64看做二进制向量的一个区间，从左到右为从低位到高位（逆序）。
	hashFuncNum int      // hashFuncNum 为哈希函数个数。
}

// NewBloomFilter 使用dataSize和allowableError初始化布隆过滤器。
// 初始化成功返回BloomFilter的指针以及error的空指针，初始化失败则返BloomFilter的空指针以及错误原因。
func NewBloomFilter(dataSize int, allowableError float64) (*BloomFilter, error) {
	if dataSize <= 0 || allowableError <= 0 || allowableError >= 1 {
		return nil, errors.New("invalid parameters")
	}

	bf := &BloomFilter{}

	// 根据数据规模dataSize和误判率allowableError求二进制向量长度的公式：
	// 			bitSize = -(dataSize * ln(allowableError)) / (ln2)^2
	bf.bitSize = int(-(float64(dataSize) * math.Log(allowableError)) / (math.Ln2 * math.Ln2))
	// 根据数据规模dataSize和误判率allowableError求哈希函数个数的公式：
	// 			hashFuncNum = (bitSize / dataSize) * ln2
	bf.hashFuncNum = int((float64(bf.bitSize) / float64(dataSize)) * math.Ln2)
	// 初始化bits切片。
	bf.bits = make([]uint64, (bf.bitSize+63)/64)

	return bf, nil
}

// getBits 获取二进制向量中的元素。
func (b *BloomFilter) getBits(index int) bool {
	if b.bits[index/64]^(1<<(index%64)) == 0 {
		return true
	}

	return false
}

// setBits 对二进制向量中的元素进行操作。
func (b *BloomFilter) setBits(index int) {
	// index/64计算出index的元素在二进制向量中的所在区间（即哪一个uint64元素），
	// index%64计算出index的元素在区间中的偏移量。
	b.bits[index/64] |= 1 << (index % 64)
}

// getIdx 计算element的hash值并转化为二进制数组中的索引。
func (b *BloomFilter) getIdx(element interface{}) ([]int, error) {
	// 类型检查。
	eleType := []byte(reflect.TypeOf(element).String())
	if string(eleType[0]) != "*" {
		return nil, errors.New("type of element must be pointer")
	}

	// 将element转化为bytes切片。
	size := int(unsafe.Sizeof(element))
	var sliceOfElement reflect.SliceHeader
	sliceOfElement.Len = size
	sliceOfElement.Cap = size
	sliceOfElement.Data = reflect.ValueOf(element).Pointer()
	bytesOfElement := *(*[]byte)(unsafe.Pointer(&sliceOfElement))

	// 计算对element的bytes切片使用keccak256算法进行哈希。
	var hashBytesOfElement []byte
	keccak := sha3.NewLegacyKeccak256()
	keccak.Write(bytesOfElement)
	hashBytesOfElement = keccak.Sum(hashBytesOfElement)

	// 获取element在bits中的所有坐标。
	var indexes []int
	firstIdx, _ := strconv.ParseInt(hex.EncodeToString(hashBytesOfElement[:4]), 16, 64)
	firstIdxInt := int(firstIdx)
	indexes = append(indexes, firstIdxInt%b.bitSize)
	for i := 1; i < b.hashFuncNum; i++ {
		idx := firstIdxInt + (firstIdxInt * i)
		indexes = append(indexes, idx%b.bitSize)
	}
	return indexes, nil
}

// Put 在BloomFilter中添加一个元素element，如果element为nil则返回错误。
func (b *BloomFilter) Put(element interface{}) error {
	if element == nil {
		return errors.New("element must not be nil")
	}

	indexes, err := b.getIdx(element)
	if err != nil {
		return err
	}

	for _, idx := range indexes {
		b.setBits(idx)
	}

	return nil
}

// IsContain 判断element是否在布隆过滤器中。
func (b *BloomFilter) IsContain(element interface{}) (bool, error) {
	if element == nil {
		return false, errors.New("element must not be nil")
	}

	indexes, err := b.getIdx(element)
	if err != nil {
		return false, err
	}

	// 检查element在bits中所有索引对应的元素，如果有一个不等于1，那么element一定不存在。
	for _, idx := range indexes {
		if !b.getBits(idx) {
			return false, nil
		}
	}

	return true, nil
}
