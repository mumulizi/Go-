package blance

import (
	"errors"
	"math/rand"
)

type RandomBlance struct {

}

func (p *RandomBlance) Doblance(insts []*Instance) (inst *Instance,err error) {
	if len(insts) ==0 {
		err = errors.New("No instance found......")
		return
	}
	index := rand.Intn(len(insts))
	inst = insts[index]
	return
}
