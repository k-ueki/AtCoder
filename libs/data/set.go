package data

type Set struct {
	m map[string]struct{}
}

func NewSet(ss ...string) *Set {
	set := &Set{}
	set.m = make(map[string]struct{})
	if len(ss) > 0 {
		for _, i := range ss {
			set.m[i] = struct{}{}
		}
	}
	return set
}

func (self *Set) Add(ss ...string) {
	if len(ss) > 0 {
		for _, i := range ss {
			self.m[i] = struct{}{}
		}
	}
}

func (self *Set) Delete(s string) {
	delete(self.m, s)
}

func (self *Set) Len() int {
	return len(self.m)
}

func (self *Set) Include(s string) bool {
	_, b := self.m[s]
	return b
}

func (self *Set) Union(other *Set) {
	self.Add(other.Elements()...)
}

func (self *Set) Intersection(other *Set) *Set {
	r := NewSet()
	for k, _ := range other.m {
		if self.Include(k) {
			r.Add(k)
		}
	}
	return r
}

func (self *Set) Elements() []string {
	r := make([]string, self.Len(), self.Len())
	var i int = 0
	for k, _ := range self.m {
		r[i] = k
		i++
	}
	return r
}
