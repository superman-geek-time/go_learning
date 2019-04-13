package maps

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 100
	NumOfWriter = 10
)

type Map interface {
	Set(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncmap(b *testing.B) {
	b.Run("map with RWLock", func(b *testing.B) {
		hm := CreateRWLockMap()
		benchmarkMap(b, hm)
	})

	b.Run("sync.map", func(b *testing.B) {
		hm := CreateSyncMapBenchmarkAdapter()
		benchmarkMap(b, hm)
	})

	b.Run("concurrent map", func(b *testing.B) {
		superman := CreateConcurrentMapBenchmarkAdapter(199)
		benchmarkMap(b, superman)
	})
}
