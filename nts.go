package main

// comes from user config device 0 - /dev/ttyUSB0
var FileDescriptor string = "/dev/ttyUSB0"
var BaudRates = []int{2000000, 1000000, 500000, 460800, 115200}
var BaudRate = 1000000
var IsOpen bool = false

// nts -l
// for var in fpga:
//     read var
//     display var

// nts -set ntp ip "10.1.10.1"
// parse command
// write fpga
// show reponse

func main() {
	cli()
	//open_port("/dev/ttyUSB0")

	////detect_baudrate()
	//var addr int64 = 0xB0020000
	//var data int64 = 0x00000001
	//connect()
	//write_reg(addr, &data)
	//read_core_config()
	//var nFlag = flag.Int("n", 1234, "help message for flag n")
	//log.Println(&nFlag)
}

//func open_port(file_descriptor string) int {
//	// try to open the selected port
//
//	log.Println("Open Port: ", file_descriptor)
//
//	if IsOpen {
//		log.Println("port is already open")
//		return -1
//	}
//	// set default rate
//	mode := &serial.Mode{
//		BaudRate: 115200,
//	}
//
//	// try to open
//	port, err := serial.Open(file_descriptor, mode)
//
//	if err != nil {
//		log.Fatal("serial open err: ", err)
//	} else {
//		log.Println("opening port: ", file_descriptor, " was successful")
//		FileDescriptor = file_descriptor
//	}
//	port.Close()
//
//	return 0
//}

//func detect_baudrate() int {
//	log.Println("Detect Baud Rate")
//	for _, rate := range BaudRates {
//		log.Println("trying: ", rate)
//		write_data := make([]byte, 0, 32)
//		read_data := make([]byte, 32)
//
//		mode := &serial.Mode{
//			BaudRate: rate,
//		}
//
//		port, err := serial.Open(FileDescriptor, mode)
//		if err != nil {
//			log.Fatal("serial open err: ", err)
//		}
//
//		port.SetReadTimeout(time.Millisecond)
//
//		write_data = append(write_data, "$CC"...)
//		checksum := calculate_checksum(write_data)
//		write_data = append(write_data, '*')
//		write_data = append(write_data, checksum...)
//		write_data = append(write_data, '\r')
//		write_data = append(write_data, '\n')
//
//		log.Println("writing: ", string(write_data))
//		n, err := port.Write(write_data)
//
//		if err != nil {
//			log.Println("write error: ", err)
//		}
//
//		if n == 0 {
//			log.Println("response: none")
//		}
//
//		n, err = port.Read(read_data)
//		port.Close()
//		read_string := string(read_data)
//		log.Printf("received: %v", read_string)
//
//		if err != nil {
//			log.Println("read error: ", err)
//		}
//
//		if n == 0 {
//			log.Println("response: none")
//		}
//
//		// check response
//		if !strings.HasPrefix(read_string, "$") {
//			log.Println("No correct response received")
//			continue
//		} else {
//			log.Println("INFO: baudrate detected at: ", rate)
//			baud_rate = rate
//			return rate
//		}
//
//	}
//
//	return -1
//}
