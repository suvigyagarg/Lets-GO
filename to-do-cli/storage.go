package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	storageDir := filepath.Join(homeDir, ".todoapp")
	os.MkdirAll(storageDir, os.ModePerm)
	fullPath := filepath.Join(storageDir, filename)

	return &Storage[T]{
		FileName: fullPath,
	}
}

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}
