package graphlator

import "io"

type StringByteWriter interface {
	io.Writer
	io.ByteWriter
	io.StringWriter
}