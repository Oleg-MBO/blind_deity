package gamelogic

import (
	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

type GroundInterface interface {
	GetLimits() (h, w int)
	IsInhabitExistOn(h, w int) bool
	IsCreatureOn(h, w int) cr.InhabitInterface
	SetCreatureOn(h, w int, inh cr.InhabitInterface)
}
