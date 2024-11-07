package service

type Expression struct {
	Filter   string
	Column   string
	Value    string
	Operator string

	Children []Expression
}

func (e *Expression) NewSubExpression() *Expression {
	sub := &Expression{}
	e.Children = append(e.Children, *sub)
	return sub
}
