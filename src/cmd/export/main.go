package main 

import(
	"C"
	"github.com/ldsec/lattigo/v2/ckks"
	"pifs/qkd/internal/he"
	"unsafe"
)


var sk *ckks.SecretKey = &ckks.SecretKey{}
var pk *ckks.PublicKey = &ckks.PublicKey{}
var rlk *ckks.EvaluationKey = &ckks.EvaluationKey{}
var gks *ckks.RotationKeys = &ckks.RotationKeys{}
var executed_once = false
func model_simulator(img, weights []float32, bias float32, reuse_key bool) (float64){

	params := he.GenRLWEParameters()

	if !reuse_key || !executed_once{
		executed_once = true
		sk, pk, rlk, gks = he.GenKeys(params)

	}
	
	
	

	ct := he.EncryptVector(img, params, pk)
	
	encoded_weights := he.EncodeWeightsVector(weights, params)

	result := he.LR(ct, encoded_weights, bias, params, rlk, gks)


	pt := he.DecryptResult(result, sk, params)

	return pt
}

func CDoublePointerToSlice(length C.int, ptr *C.double)([]float64){
	l  := int(length)
	ret1 := (*[1 << 30]C.double)(unsafe.Pointer(ptr))[:l:l]
	ret2 := make([]float64, l) 
	for i := 0; i < l ; i++ {
		ret2[i] = float64(ret1[i])
	}
	return ret2
}

func CFloatPointerToSlice(length C.int, ptr *C.float)([]float32){
	l  := int(length)
	ret1 := (*[1 << 30]C.float)(unsafe.Pointer(ptr))[:l:l]
	ret2 := make([]float32, l) 
	for i := 0; i < l ; i++ {
		ret2[i] = float32(ret1[i])
	}
	return ret2
}

func cBool(val C.int) bool {
	if val > 0{
		return true
	} else {
		return false
	}
}

//export test
func test(img_shape C.int, img_ptr *C.float, 
		  weights_shape C.int, weights_ptr *C.float, 
		  c_bias C.float,
		  c_reuse_key C.int) (C.double){
	img := CFloatPointerToSlice(img_shape, img_ptr)
	weights := CFloatPointerToSlice(weights_shape, weights_ptr)
	bias := float32(c_bias)
	reuse_key := cBool(c_reuse_key)
	res := model_simulator(img, weights, bias, reuse_key)
	return C.double(res)
}

func main() {}