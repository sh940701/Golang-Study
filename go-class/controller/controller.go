package controller

type ctl struct {}

func NewCTL() *ctl {
	ctl := &ctl{}
	return ctl
}

func (c *ctl) CalcSum(a, b int) int {
	return a + b
}

func (c *ctl) CalcMul(a, b int) int {
	return a * b
}

func (c *ctl) CalcDiv(a, b float32) float32 {
	return b / a
}

func (c *ctl) CalcSub(a, b int) int {
	return b - a
}