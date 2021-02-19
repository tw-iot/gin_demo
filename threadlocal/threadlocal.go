package threadlocal

import . "github.com/jtolds/gls"

var (
	Mgr = NewContextManager()
	Rid = GenSym()
)
