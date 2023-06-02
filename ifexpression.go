package tips

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
