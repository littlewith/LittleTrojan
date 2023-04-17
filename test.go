package main

/*
func main() {
	var programData bytes.Buffer
	var tmpBuffer []byte = make([]byte, 1024)
	programBuffer, _ := os.Open("nc.exe")
	for {
		length, _ := programBuffer.Read(tmpBuffer)
		if length == 0 {
			break
		}
		programData.Write(tmpBuffer)
	}
	newfile, _ := os.OpenFile("new.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	newfile.WriteString(programData.String())
}
*/
