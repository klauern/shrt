package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/klauern/shrt/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
