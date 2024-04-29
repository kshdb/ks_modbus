package main

import (
	"fmt"
	"time"

	modbus "github.com/kshdb/ks_modbus"
)

func main() {
	p := modbus.NewTCPClientProvider("192.168.1.21:501", modbus.WithEnableLogger())
	client := modbus.NewClient(p)
	err := client.Connect()
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	defer client.Close()

	//写入保持寄存器
	client.WriteMultipleRegistersBytes(2, 0, 2, []byte{00, 01, 00, 02})
	for {
		_results, err := client.ReadHoldingRegistersBytes(1, 0, 10)
		if err != nil {
			fmt.Println(err.Error())
		}

		//fmt.Printf("保存寄存器数据是 %#v\r\n", _results)

		grouped := groupBytes(_results, 2)
		for _, v := range grouped {
			_v := fmt.Sprintf("%x", v)
			fmt.Printf("保存寄存器数据分组结果是%s\n", _v)
		}

		time.Sleep(time.Millisecond * 500)
	}
}

// 数组分组
func groupBytes(input []byte, groupSize int) [][]byte {
	var grouped [][]byte
	for i := 0; i < len(input); i += groupSize {
		end := i + groupSize
		if end > len(input) {
			end = len(input)
		}
		grouped = append(grouped, input[i:end])
	}
	return grouped
}
