type MyHashSet struct {
    set map[int]bool
}

func Constructor() MyHashSet {
    return MyHashSet{set: make(map[int]bool)}
}

func (this *MyHashSet) Add(key int) {
    this.set[key] = true
}

func (this *MyHashSet) Remove(key int) {
    delete(this.set, key)
}

func (this *MyHashSet) Contains(key int) bool {
    return this.set[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */