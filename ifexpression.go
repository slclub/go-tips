package tips

/**
 * The first method
 * 比较浪费性能的，需要内存申请的 三目运算
 */
type ifExpression struct {
	cond  bool
	value any
}

func Three(cond bool) *ifExpression {
	return &ifExpression{
		cond: cond,
	}
}

func (self *ifExpression) If(v any) *ifExpression {
	if self.cond {
		self.value = v
	}
	return self
}

func (self *ifExpression) Else(v any) *ifExpression {
	if !self.cond {
		self.value = v
	}
	return self
}

func (self *ifExpression) Value() any {
	return self.value
}

func (self *ifExpression) Int() int {
	return Int(self.Value())
}

func (self *ifExpression) String() string {
	return String(self.Value())
}

/**
 * The second method
 */
func IfThree(cond bool) func(v1, v2 any) any {
	return func(v1, v2 any) any {
		if cond {
			return v1
		}
		return v2
	}
}
