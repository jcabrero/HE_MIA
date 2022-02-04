package he

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"log"
)


func GenRLWEParameters() (p *ckks.Parameters) {
	// This function is meant to give a usable parameter set for HE.
	var err error
	var logN uint64 = 15
	var logModuli *ckks.LogModuli = new(ckks.LogModuli)
	logModuli.LogQi = []uint64{60, 40, 40, 40, 60};
	logModuli.LogPi = []uint64{60,}

	
	p, err = ckks.NewParametersFromLogModuli(logN, logModuli)
	if err != nil{
		log.Fatal("Couldn't create the parameters")
	}
	var scale float64 = 1 << 30
	p.SetLogSlots(logN - 1)
	p.SetScale(scale)
	return 
}

