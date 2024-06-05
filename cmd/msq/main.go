package main

func main() {
	dataLog, err := NewLog("topic1")
	if err != nil {
		panic(err)
	}
	defer dataLog.Close()

	dataLog.Append([]byte("Hello, World!"))
	dataLog.Append([]byte("Hello, World!"))
}
