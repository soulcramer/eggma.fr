package frontpage

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/eggma.fr/components"
)

// Get ...
func Get(ctx *aero.Context) string {

	return ctx.HTML(components.FrontPage())
}