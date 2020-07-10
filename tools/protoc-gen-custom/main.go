package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/imura-tatsuya-ab/protoc-gen-option-example/custom"
)

func parseReq(r io.Reader) (*plugin.CodeGeneratorRequest, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var req plugin.CodeGeneratorRequest
	if err = proto.Unmarshal(buf, &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func processReq(req *plugin.CodeGeneratorRequest) *plugin.CodeGeneratorResponse {
	files := make(map[string]*descriptor.FileDescriptorProto)
	for _, f := range req.ProtoFile {
		files[f.GetName()] = f
	}

	var buf bytes.Buffer
	for _, fname := range req.FileToGenerate {
		f := files[fname]
		for _, name := range targetNames(f) {
			io.WriteString(&buf, name)
			io.WriteString(&buf, "\n")
		}
	}
	return &plugin.CodeGeneratorResponse{
		File: []*plugin.CodeGeneratorResponse_File{
			{
				Name:    proto.String("messages.txt"),
				Content: proto.String(buf.String()),
			},
		},
	}
}

func targetNames(f *descriptor.FileDescriptorProto) []string {
	var list []string
	for _, m := range f.MessageType {
		if !isTarget(m) {
			continue
		}
		list = append(list, m.GetName())
	}
	return list
}

func isTarget(m *descriptor.DescriptorProto) bool {
	var opts = m.GetOptions()
	if opts == nil {
		return false
	}
	ext, err := proto.GetExtension(opts, custom.E_Custom)
	if err == proto.ErrMissingExtension {
		return false
	}
	if err != nil {
		panic("unexpected error")
	}

	mopts := ext.(*custom.MessageCustomOptions)
	return mopts.GetTarget()
}

func emitResp(resp *plugin.CodeGeneratorResponse) error {
	buf, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(buf)
	return err
}

func run() error {
	req, err := parseReq(os.Stdin)
	if err != nil {
		return err
	}

	resp := processReq(req)

	return emitResp(resp)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
