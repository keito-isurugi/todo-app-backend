package entity

type PutObjectInput struct {
	Key         string
	ContentType string
	FileContent []byte
}

func NewPutObjectInput(key, contentType string, fileContent []byte) *PutObjectInput {
	return &PutObjectInput{
		Key:         key,
		ContentType: contentType,
		FileContent: fileContent,
	}
}
