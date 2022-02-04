package he

import(
	"fmt"
	"log"
	"os"
	"encoding/binary"
	"bytes"
	"math"
	"io/ioutil"
	"encoding/json"
	//"github.com/ldsec/lattigo/v2/ckks"

)

type Ops struct{
	Ops []Op `json:"ops"`
}

type Op struct{
	Name string `json:"name"`
	Opcode int `json:"opcode"`
	NumArgs int `json:"num_args"`
	Args []string `json:"args"`
	Out string `json:"out"`
	Properties OpProps `json:"characteristics"`
}

type OpProps struct{
	Int bool `json:"int"`
	Float bool `json:"float"`
	Sym bool `json:"sym"`
	Free bool `json:"free"`
	Vector bool `json:"vector"`
}

var Operations Ops
var op_name_index map[string]int = make(map[string]int)
var op_code_index map[int]int = make(map[int]int)

func LoadOps(){
	// Open our jsonFile
	filename := os.Getenv("PATH_TO_CONFIG")
	fmt.Println(filename)
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully loaded")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array


	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &Operations)
	for i := 0; i < len(Operations.Ops); i++{
		op := Operations.Ops[i]
		op_name_index[op.Name] = i
		op_code_index[op.Opcode] = i
		fmt.Printf("%+v\n", op)
		
	}

}
func get_op_name(op byte) string{
	return Operations.Ops[op_code_index[int(op)]].Name
}

func get_out_type(op byte) string{
	return Operations.Ops[op_code_index[int(op)]].Out
}

func get_op_args(op byte) []string{
	t := append(Operations.Ops[op_code_index[int(op)]].Args, Operations.Ops[op_code_index[int(op)]].Out)
	return t
}
func get_op_args_no_out(op byte) []string{
	t := Operations.Ops[op_code_index[int(op)]].Args
	return t
}

func get_opcode(s string) int{
	return Operations.Ops[op_name_index[s]].Opcode
}

func is_int_op(op byte) bool{
	return Operations.Ops[op_code_index[int(op)]].Properties.Int
}


func is_float_op(op byte) bool{
	return Operations.Ops[op_code_index[int(op)]].Properties.Float
}


func is_float_or_int_op(op byte) bool{
	return is_int_op(op) || is_float_op(op)
}

func is_sym_op(op byte) bool{
	return Operations.Ops[op_code_index[int(op)]].Properties.Sym
}

func is_free_op(op byte) bool{
	return Operations.Ops[op_code_index[int(op)]].Properties.Free
}

func is_vector_op(op byte) bool{
	return Operations.Ops[op_code_index[int(op)]].Properties.Vector
}

func op_num_args(op byte) int{
	// TODO: Elaborate a bit more (switch)
	return Operations.Ops[op_code_index[int(op)]].NumArgs
}

func Use(vals ...interface{}) {
    for _, val := range vals {
        _ = val
    }
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func writeNextBytes(file *os.File, bytes []byte) {

	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

}

func write_uint32(file *os.File, num uint32){
	var bin_buf *bytes.Buffer = new(bytes.Buffer)
	err := binary.Write(bin_buf, binary.BigEndian, num)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	writeNextBytes(file, bin_buf.Bytes())
}

/*Helper functions to read from memory*/


func uint32_from_bytes(data [] byte) (d uint32){
	d = binary.BigEndian.Uint32(data)
	return 
}

func uint64_from_bytes(data [] byte) (d uint64){
	d = binary.BigEndian.Uint64(data)
	return 
}

func float32_from_bytes(data [] byte) (d float32){
	mem_int := binary.BigEndian.Uint32(data)
	d = math.Float32frombits(mem_int)
	return 
}

func float64_from_bytes(data [] byte) (d float64){
	mem_int := binary.BigEndian.Uint64(data)
	d = math.Float64frombits(mem_int)
	return 
}


// READ FROM MEMORY INFORMATION

func read_byte(f *os.File) (byte){
	data := readNextBytes(f, 1)
	return data[0]
}

func read_int32(f *os.File) (int32){
	return int32(read_uint32(f))
}
func read_uint32(f *os.File) (uint32){
	data := readNextBytes(f, 4)
	mem_int := uint32_from_bytes(data)
	return mem_int
}

func read_uint64(f *os.File) (uint64){
	data := readNextBytes(f, 8)
	mem_int := uint64_from_bytes(data)
	return mem_int
}

func read_float32(f *os.File) (float32){
	data := readNextBytes(f, 4)
	mem_float := float32_from_bytes(data)
	return mem_float
}

func read_float64(f *os.File) (float64){
	data := readNextBytes(f, 8)
	mem_float := float64_from_bytes(data)
	return mem_float
}

func read_float32_vector(f *os.File) ([]complex128){
	dim := read_uint32(f)
	
	var arr []complex128 = make([]complex128, dim)
	// var arr []float32 =  make([]float32, dim)
	for j := uint32(0); j < dim; j++ {
		float_num := read_float32(f)
		arr[j] =  complex128(complex(float_num, 0))
	}
	return arr
}
func binary_write_wrap(bin_buf *bytes.Buffer, d interface{}){
	err := binary.Write(bin_buf, binary.BigEndian, d)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
}

/* Filesystem interaction */

func CreateDir(path string){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	} else {
		fmt.Println("Directory ", path, " already exists")
	}
}
func RemoveFile(path string){
	e := os.Remove(path)
	if e != nil {
		log.Fatal(e)
	}
}

func list_dir(path string) []os.FileInfo{
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
