package common

import (
	"fmt"
	//log "github.com/sirupsen/logrus"
)

func FormatSize(bytes uint64) string {
	var cNumber float32 = 1024
	kb := float32(bytes) / cNumber
	if kb >= cNumber {
		mb := kb / cNumber
		if mb >= cNumber {
			gb := mb / cNumber
			if gb >= cNumber {
				tb := gb / cNumber
				return fmt.Sprintf("%.2fT", tb)
			} else {
				return fmt.Sprintf("%.2fG", gb)
			}
		} else {
			return fmt.Sprintf("%.2fM", mb)
		}
	} else {
		return fmt.Sprintf("%.2fK", kb)
	}
}

// 定长管道
type Pipe struct {
	Length int
	Data   []float64
	head   int
	tail   int
}

func (P *Pipe) NewPipe(length int) {
	P.Length = length
	P.Data = make([]float64, 0, P.Length)
	P.head = 0
	P.tail = 0
}

func (P *Pipe) Push(n float64) {
	dataLength := len(P.Data)
	//log.Info("Length: ", dataLength, P.Length, P.Data)

	if dataLength < P.Length {
		P.Data = append(P.Data, n)
		if dataLength != 0 {
			P.head++
		}
	} else {
		P.Data[P.tail] = n
		P.tail = (P.tail + 1) % P.Length
		P.head = (P.head + 1) % P.Length
	}
}

func (P *Pipe) Show() []float64 {
	if P.head == P.tail {
		return P.Data
	} else {
		data := []float64{}
		for tempHead := P.head; tempHead >= 0; tempHead-- {
			data = append(data, P.Data[tempHead])
		}
		if P.head < P.tail {
			for tempTail := P.Length - 1; tempTail >= P.tail; tempTail-- {
				data = append(data, P.Data[tempTail])
			}
		}
		return data
	}
}

type Rate struct {
	data [2]float64
}

func (R *Rate) Push(n float64) {
	// [1]: new    [0]: old
	R.data[0] = R.data[1]
	R.data[1] = n
}

func (R *Rate) Get() (newData, oldData float64) {
	newData = R.data[1]
	oldData = R.data[0]
	return
}

func (R *Rate) GetRate(data float64, t float64) float64 {
	R.Push(data)
	newData, oldData := R.Get()
	rate := (newData - oldData) / t
	//log.Info("Ratedata: ", data, " newData ", newData, " oldData: ", oldData, " Rate ", rate)
	return rate
}
