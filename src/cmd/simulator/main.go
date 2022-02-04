package main 

import(
	//"fmt"
	//"os"
	"pifs/qkd/internal/he"
	"flag" // For cmdline args
)

func simulator(params_file, public_key_file, secret_key_file, evaluation_key_file, rotation_key_file, input_file, ciphertext_file,  result_ciphertext_file string) (float64){
	//fmt.Println("This is my image", img)
	params := he.GenRLWEParameters()
	sk, pk, rlk, gks := he.GenKeys(params)
	
	img := he.GetImageFromFilePath(input_file)

	ct := he.EncryptImage(img, params, pk)
	he.CiphertextToFile(ciphertext_file, ct)
	
	weights, bias := he.OpenModel()
	encoded_weights := he.EncodeWeightsVector(weights, params)

	result := he.LR(ct, encoded_weights,bias, params, rlk, gks)


	pt := he.DecryptResult(result, sk, params)

	return pt
}

func main() {
	var params_file = flag.String("p", "data/he/parameters.params", "Encryption Parameters File")
	var public_key_file = flag.String("pk", "data/he/key.pk", "Public Key File")
	var secret_key_file = flag.String("sk", "data/he/key.sk", "Secret Key File")
	var evaluation_key_file = flag.String("rlk", "data/he/key.rlk", "Evaluation Key File")
	var rotation_key_file = flag.String("gks", "data/he/key.gks", "Rotation Key File")

	var input_file = flag.String("i", "data/img/frontal.png", "Input Image File")
	var ciphertext_file = flag.String("c", "data/he/img.enc", "Output Encrypted Image")

	var result_ciphertext_file = flag.String("oc", "data/he/result.enc", "Output Result Ciphertext")
	flag.Parse()

	simulator(*params_file, *public_key_file, *secret_key_file, *evaluation_key_file, *rotation_key_file, *input_file, *ciphertext_file,  *result_ciphertext_file)

}