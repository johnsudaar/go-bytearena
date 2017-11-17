package agent

type EmptyAgent struct{}

func (*EmptyAgent) Raw([]byte, int) {}
func (*EmptyAgent) Worker()         {}
