package basic

// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
const offset64 uint64 = 14695981039346656037
// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
const prime64 uint64= 1099511628211

// FNV-1a
func Hash(b []byte) uint64 {
	res := offset64
	for _,b := range b{
		res ^= uint64(b)
		res *= prime64
	}
	return res
}
