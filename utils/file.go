package utils

import (
	"bytes"
	"mime/multipart"
	"net/textproto"
	"os"
)

func StringToMultipartFileHeader(filePath string) (*multipart.FileHeader, error) {
	// Open the file from the file path
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Get file info for size and name
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Create a buffer to simulate a multipart file
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Create a form file field
	part, err := writer.CreateFormFile("file", fileInfo.Name())
	if err != nil {
		return nil, err
	}

	// Copy file content to the form file
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteTo(part)
	if err != nil {
		return nil, err
	}

	// Close the writer to finalize the multipart message
	writer.Close()

	// Create a FileHeader manually
	fileHeader := multipart.FileHeader{
		Filename: fileInfo.Name(),
		Size:     fileInfo.Size(),
		Header:   textproto.MIMEHeader{"Content-Type": []string{"application/octet-stream"}},
	}

	return &fileHeader, nil
}