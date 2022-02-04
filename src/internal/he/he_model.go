
package he

import (
	"github.com/ldsec/lattigo/v2/ckks"
	"log"
)

func EncodeWeightsVector(v []float32, params *ckks.Parameters)(*ckks.Plaintext){
	values := make([]complex128, len(v))
	for i := 0; i < len(v); i++ {
		values[i] = complex128(complex(v[i], 0))
	}
	var encoder ckks.Encoder = ckks.NewEncoder(params)
	plaintext := encoder.EncodeNTTAtLvlNew(params.MaxLevel(), values, params.LogSlots())
	return plaintext
}

func LR(ct *ckks.Ciphertext, 
	weights *ckks.Plaintext,
	bias float32,
	params *ckks.Parameters, 
	rlk *ckks.EvaluationKey,
	gks *ckks.RotationKeys) (*ckks.Ciphertext){

	var evaluator ckks.Evaluator = ckks.NewEvaluator(params)
	
	r1 := evaluator.MulRelinNew(ct, weights , rlk)

	if err := evaluator.Rescale(r1, params.Scale(), r1); err != nil {
		log.Fatal(err)
	}

	var t *ckks.Ciphertext
	for i := 0; i < int(params.LogN()); i++{
		t = evaluator.RotateNew(r1, uint64(1 << i), gks)
		evaluator.Add(r1, t, r1)
	}
	//cbias := complex128(complex(bias, 0))
	res := evaluator.AddConstNew(r1, bias)
    return res
}