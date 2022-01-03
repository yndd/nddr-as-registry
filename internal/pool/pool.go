package pool

import (
	"k8s.io/apimachinery/pkg/labels"
)

type Pool interface {
	InsertByIndex(uint32, string, map[string]string) uint32
	DeleteByIndex(uint32, string, map[string]string)
	InsertByName(string, string, map[string]string) uint32
	DeleteByName(string, string, map[string]string)
	GetAllocated() (uint32, []*uint32)
}

type node struct {
	key      uint32
	name     string
	register map[string]*labels.Set
}

type pool struct {
	size               uint32
	allocationStrategy string
	nodes              []*node
}

func New(start, end uint32, allocStrategy string) Pool {
	s := end - start
	p := &pool{
		size:               s,
		allocationStrategy: allocStrategy,
		nodes:              make([]*node, s),
	}

	for i := 0; i < len(p.nodes); i++ {
		p.nodes[i] = &node{
			key:      start + uint32(i),
			name:     "",
			register: make(map[string]*labels.Set),
		}
	}
	return p
}

func (h *pool) InsertByIndex(idx uint32, n string, l map[string]string) uint32 {
	hidx := idx
	return h.insert(hidx, "dummy", n, l)
}

func (h *pool) DeleteByIndex(idx uint32, n string, l map[string]string) {
	hidx := idx
	h.delete(0, hidx, "dummy", n, l)
}

func (h *pool) InsertByName(k, n string, l map[string]string) uint32 {
	hidx := uint32(0)
	return h.insert(hidx, k, n, l)
}

func (h *pool) DeleteByName(k, n string, l map[string]string) {
	hidx := uint32(0)
	h.delete(0, hidx, k, n, l)
}

func (h *pool) GetAllocated() (uint32, []*uint32) {
	used := make([]*uint32, 0)
	allocated := uint32(0)
	for _, n := range h.nodes {
		if n.name != "" {
			allocated++
			used = append(used, &n.key)
		}
	}
	return allocated, used
}

func (h *pool) insert(hidx uint32, k, n string, l map[string]string) uint32 {
	mergedlabel := labels.Merge(labels.Set(l), nil)
	// if entry is empty or the key is already used, insert the key and return the hash index
	if h.nodes[hidx].name == "" || h.nodes[hidx].name == k {
		h.nodes[hidx].name = k
		h.nodes[hidx].register[n] = &mergedlabel
		return h.nodes[hidx].key
	}
	hidx++
	if hidx >= h.size {
		hidx = 0
	}
	return h.insert(hidx, k, n, l)
}

// k is the hashkey
// n is the name of the register or allocation
// l is the label
// ofidx is the overflow idx, used to ensure if we delete a resource that does not exist we stop
// hidx is the hash idx and we use overflow mapping by incrementing the hidx if a hash collision occurs
func (h *pool) delete(ofidx, hidx uint32, k, n string, l map[string]string) {
	// if entry is empty, insert the key and return the hash index
	if h.nodes[hidx].name == k {
		delete(h.nodes[hidx].register, n)
		// if the index has no longer has registers/allocations, so we can delete the key
		if len(h.nodes[hidx].register) == 0 {
			h.nodes[hidx].name = ""
			h.nodes[hidx].register = make(map[string]*labels.Set)
		}
		return
	}
	hidx++
	ofidx++
	if hidx >= h.size {
		hidx = 0
	}
	if ofidx == h.size {
		// the entry was not found, so we can stop
		return
	}
	h.delete(ofidx, hidx, k, n, l)
}
