package main

import "bytes"

func middleware(bodyBytes []byte) ([]byte, error) {

	body_slices := SplitByBlankLine(bodyBytes)

	if len(body_slices) < 3 {
		return bodyBytes, nil
	}

	var buffer bytes.Buffer
	buffer.Write(body_slices[len(body_slices)-3])
	buffer.Write([]byte("\n\n"))
	buffer.Write([]byte("data: [DONE]"))
	buffer.Write([]byte("\n\n"))
	newBodyBytes := buffer.Bytes()
	println(string(newBodyBytes))

	return newBodyBytes, nil
}
