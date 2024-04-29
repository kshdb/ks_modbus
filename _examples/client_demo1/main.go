package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goburrow/modbus"
)

func main() {
	handler := modbus.NewTCPClientHandler("192.168.1.21:501")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 1
	handler.Logger = log.New(os.Stdout, "log: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	if err != nil {
		fmt.Println("connect failed, ", err)
		return
	}
	defer handler.Close()
	client := modbus.NewClient(handler)
	// client.WriteSingleRegister(10, uint16(1001))
	//写指令
	client.WriteMultipleRegisters(1, 1, []byte{46, 63})
	for {
		//01：COIL STATUS（线圈状态）：用于读取和控制远程设备的开关状态，通常用于控制继电器等开关设备。
		//client.ReadCoils(1, 10)//读操作
		//client.WriteSingleCoil(1, 0) //写操作
		//client.ReadCoils(1, 10) //读操作

		//02：INPUT STATUS（输入状态）：用于读取远程设备的输入状态，通常用于读取传感器等输入设备的状态。
		//client.ReadInputRegisters(1, 10)

		//03：HOLDING REGISTER（保持寄存器）：用于存储和读取远程设备的数据，通常用于存储控制参数、设备状态等信息。
		//results, err := client.ReadHoldingRegisters(0, 10)
		results, _ := client.ReadHoldingRegisters(0, 10)
		grouped := groupBytes(results, 2)
		for _, v := range grouped {
			_v := fmt.Sprintf("%x", v)
			fmt.Printf("保存寄存器数据分组结果是%s\n", _v)
		}

		//04：INPUT REGISTER（输入寄存器）：用于存储远程设备的输入数据，通常用于存储传感器等输入设备的数据。
		//client.ReadInputRegisters(1, 10)

		// results, err = client.ReadDiscreteInputs(15, 2)
		// results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
		// results, err = client.WriteMultipleCoils(5, 10, []byte{4, 3})

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
