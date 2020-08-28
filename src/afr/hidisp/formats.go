package hidisp

type format struct {
	Resolution string
	Fps        int
}

var Formats = map[string]format{
	"0":  {"1080", 60},
	"1":  {"1080", 50},
	"2":  {"1080", 30},
	"3":  {"1080", 25},
	"4":  {"1080", 24},
	"5":  {"1080i", 60},
	"6":  {"1080i", 50},
	"7":  {"720", 60},
	"8":  {"720", 50},
	"9":  {"576", 50},
	"10": {"480", 60},
	"64": {"3840X2160", 24},
	"65": {"3840X2160", 25},
	"66": {"3840X2160", 30},
	"67": {"3840X2160", 50},
	"68": {"3840X2160", 60},
	"69": {"4096X2160", 24},
	"70": {"4096X2160", 25},
	"71": {"4096X2160", 30},
	"72": {"4096X2160", 50},
	"73": {"4096X2160", 60},
	"74": {"3840X2160", 23},
	"75": {"3840X2160", 29},
	"76": {"720", 59},
	"77": {"1080", 59},
	"78": {"1080", 29},
	"79": {"1080", 23},
	"80": {"1080i", 59},
	//81: {"2560X1080", 60},
	//82: {"2560X1440", 60},
	//83: {"3440X1440", 60},
	//84: {"3840X1080", 60},
	//85: {"3840X1200", 60},
}
