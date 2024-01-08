package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(in []byte) (int, error) {
	for {
		in = append(in, 'A')
	}
	return 0, nil
}


func main() {
	reader.Validate(MyReader{})
}
