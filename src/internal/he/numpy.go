package he

import (
	"log"
	"os"
	"github.com/sbinet/npyio/npy"
)

func OpenModel()([]float32, float32){

	var bias float32 = -0.05991028
	f, err := os.Open("data/model/w.npy")
	if err != nil{
		log.Fatal(err)
	}
	var weights []float32
	err = npy.Read(f, &weights)
	if err != nil {
		log.Fatal(err)
	}
	
	return  weights, bias
}

func OpenModelFromFile(model_file string)([]float32, float32){

	var bias float32 = -0.05991028
	f, err := os.Open(model_file)
	if err != nil{
		log.Fatal(err)
	}
	var weights []float32
	err = npy.Read(f, &weights)
	if err != nil {
		log.Fatal(err)
	}
	
	return  weights, bias
}