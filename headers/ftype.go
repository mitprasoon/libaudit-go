package headers

//Ftype is a map of file type and their correspondig values
var Ftype = map[string]int{
	"block":     24576,
	"character": 8192,
	"dir":       16384,
	"fifo":      4096,
	"file":      32768,
	"link":      40960,
	"socket":    49152,
}
