package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/uuid"
)

const (
	maxShards = 8
)

func uuidToInt(uuid string) int {
	// Hash the UUID
	hash := md5.Sum([]byte(uuid))
	hashStr := hex.EncodeToString(hash[:])
	// fmt.Println(hashStr)
	// Convert hexadecimal hash to an integer
	seed, _ := strconv.ParseInt(hashStr, 16, 64)

	// Initialize the random number generator with the hashed UUID
	rand.Seed(seed)

	// Generate a random number between 1 and 8
	randomNumber := rand.Intn(maxShards) + 1

	return randomNumber
}

func uuidToInt2(uuid string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(uuid))
	hash := uint32(hasher.Sum32())

	rand.Seed(int64(hash % uint32(maxShards)))

	randomNumber := rand.Intn(maxShards) + 1

	return randomNumber
}

func main() {
	// // Example UUID
	// uuid := "550e8400-e29b-41d4-a716-446655440000"

	// // Call the function with the same UUID multiple times
	// for i := 0; i < 5; i++ {
	// 	result := uuidToShard(uuid)
	// 	fmt.Println("Mapped number:", result)
	// }

	result := testing.Benchmark(BenchmarkUUIDToInt)
	fmt.Printf("Time taken: %s\n", result.T)
	fmt.Printf("Memory used: %d bytes\n", result.MemBytes)
	fmt.Printf("Iterations: %d\n", result.N)

	fmt.Printf("\n \n \n")

	result = testing.Benchmark(BenchmarkUUIDToInt2)
	fmt.Printf("Time taken: %s\n", result.T)
	fmt.Printf("Memory used: %d bytes\n", result.MemBytes)
	fmt.Printf("Iterations: %d\n", result.N)
}

func BenchmarkUUIDToInt(b *testing.B) {
	//for i := 0; i < b.N; i++ {
	newUUID := uuid.New()
	uuidToInt(newUUID.String())
	//}
}

func BenchmarkUUIDToInt2(b *testing.B) {

	//for i := 0; i < b.N; i++ {
	newUUID := uuid.New()
	uuidToInt2(newUUID.String())
	//}
}
