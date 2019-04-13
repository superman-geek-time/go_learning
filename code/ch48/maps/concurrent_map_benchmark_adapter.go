package maps

import "github.com/easierway/concurrent_map"

type ConcurrentMapBenchmarkAdapter struct {
	cm *concurrent_map.ConcurrentMap
}

func (m *ConcurrentMapBenchmarkAdapter) Set(key interface{}, value interface{}) {
	m.cm.Set(concurrent_map.StrKey(key.(string)), value)
}

func (m *ConcurrentMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.cm.Get(concurrent_map.StrKey(key.(string)))
}

func (m *ConcurrentMapBenchmarkAdapter) Del(key interface{}) {
	m.cm.Del(concurrent_map.StrKey(key.(string)))
}

func CreateConcurrentMapBenchmarkAdapter(numOfPartitions int) *ConcurrentMapBenchmarkAdapter {
	conMap := concurrent_map.CreateConcurrentMap(numOfPartitions)
	return &ConcurrentMapBenchmarkAdapter{conMap}
}
