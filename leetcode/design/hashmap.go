package main

type pair struct {
	key int
	val int
}

// ListNode struct
type ListNode struct {
	pair *pair
	next *ListNode
}

// HashMap struct
type HashMap struct {
	size int
	hash []*ListNode
}

func NewHashMap() *HashMap {
	return &HashMap{
		size: 1000,
		hash: make([]*ListNode, 1000),
	}
}

func (hm *HashMap) put(key, value int) {
	idx := key % hm.size

	if hm.hash[idx] == nil {
		hm.hash[idx] = &ListNode{
			pair: &pair{
				key: key,
				val: value,
			},
		}
	} else {
		cur := hm.hash[idx]

		for cur != nil {
			// update , if key already present
			if cur.pair.key == key {
				cur.pair.val = value
			}

			if cur.next == nil {
				break
			}

			cur = cur.next
		}
		cur.next = &ListNode{
			pair: &pair{
				key: key,
				val: value,
			},
		}
	}
}

func (hm *HashMap) get(key int) int {
	idx := key % hm.size
	cur := hm.hash[idx]

	for cur != nil {
		if cur.pair.key == key {
			return cur.pair.val
		}
		cur = cur.next
	}
	return -1
}

func (hm *HashMap) remove(key int) {
	idx := key % hm.size
	cur, prev := hm.hash[idx], hm.hash[idx]

	if cur == nil {
		return
	}

	if cur.pair.key == key {
		hm.hash[idx] = cur.next
	} else {
		cur = cur.next

		for cur != nil {
			if cur.pair.key == key {
				prev.next = cur.next
				break
			} else {
				prev, cur = prev.next, cur.next
			}
		}
	}
}
