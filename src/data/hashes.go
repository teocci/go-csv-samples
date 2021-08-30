// Package data
// Created by RTT.
// Author: teocci@yandex.com on 2021-Aug-30
package data

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"hash/fnv"
	"strconv"
)

// FNV32 hashes using fnv32 algorithm
func FNV32(s string) uint32 {
	return uint32Hasher(fnv.New32(), s)
}

// FNV32a hashes using fnv32a algorithm
func FNV32a(s string) uint32 {
	return uint32Hasher(fnv.New32a(), s)
}

// FNV64 hashes using fnv64 algorithm
func FNV64(s string) uint64 {
	return uint64Hasher(fnv.New64(), s)
}

// FNV64a hashes using fnv64a algorithm
func FNV64a(s string) uint64 {
	return uint64Hasher(fnv.New64a(), s)
}

func FNV64aS(s string) string {
	return strconv.FormatUint(uint64Hasher(fnv.New64a(), s), 10)
}

// MD5 hashes using md5 algorithm
func MD5(s string) string {
	return stringHasher(md5.New(), s)
}

// SHA1 hashes using sha1 algorithm
func SHA1(s string) string {
	return stringHasher(sha1.New(), s)
}

// SHA256 hashes using sha256 algorithm
func SHA256(s string) string {
	return stringHasher(sha256.New(), s)
}

// SHA512 hashes using sha512 algorithm
func SHA512(text string) string {
	return stringHasher(sha512.New(), text)
}

func stringHasher(alg hash.Hash, s string) string {
	alg.Write([]byte(s))
	return hex.EncodeToString(alg.Sum(nil))
}

func uint32Hasher(alg hash.Hash32, s string) uint32 {
	_, _ = alg.Write([]byte(s))
	return alg.Sum32()
}

func uint64Hasher(alg hash.Hash64, s string) uint64 {
	_, _ = alg.Write([]byte(s))
	return alg.Sum64()
}