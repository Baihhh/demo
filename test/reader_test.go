package test

import "testing"

func TestReaderLimit(t *testing.T) {
	ReaderLimit()
}

func TestReaderMulti(t *testing.T) {
	ReaderMulti()
}

func TestReaderTee(t *testing.T) {
	ReaderTee()
}

func TestReaderRune(t *testing.T) {
	ReaderRune()
}

func TestReaderSection(t *testing.T) {
	ReaderSection()
}

func TestReaderReadLine(t *testing.T) {
	ReaderReadLine()
}
