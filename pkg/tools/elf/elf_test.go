package elf

import "testing"

var (
	goGrpcWriterHeader = "google.golang.org/grpc/internal/transport.(*loopyWriter).writeHeader"
	goGrpcLoopyWriter  = "google.golang.org/grpc/internal/transport.loopyWriter"
)

func TestElf(t *testing.T) {
	// filePath := "/home/xince/cxc/beyla/examples/example-http-service/example-http-service"
	filePath := "/home/xince/cxc/GO-grpc-demo/server/server"
	file, err := NewFile(filePath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	reader, err := file.NewDwarfReader(
		goGrpcWriterHeader, goGrpcLoopyWriter,
	)
	if err != nil {
		t.Fatal(err)
	}
	if reader.language != ReaderLanguageGolang {
		t.Fatalf("Expected Language: %d, But Got Language: %d", ReaderLanguageGolang, reader.language)
	}
	t.Logf("reader.language: %d reader.Producer: %s", reader.language, reader.producer)

	readFunction := reader.GetFunction(goGrpcWriterHeader)
	if readFunction == nil {
		t.Fatal("Expected none emptry function, get nil")
	}
	t.Log("readFunction", readFunction)
	t.Logf("readFunction.Name: %s\t readFunction.Args: %v", readFunction.name, readFunction.args)
	for k, x := range readFunction.args {
		t.Logf("Name: %s\tType: %v\tIsRet: %v\tLocation.Type: %d\tLocation.Offset: %d", k, x.tp, x.IsRet, x.Location.Type, x.Location.Offset)
	}

	readStructure := reader.GetStructure(goGrpcLoopyWriter)
	if readStructure == nil {
		t.Fatal("Expected none emptry function, get nil")
	}
	t.Logf("readStructure: %s\t readStructure.Fields: %v", readStructure.name, readStructure.fields)
	for _, field := range readStructure.fields {
		t.Logf("Name: %s\tOffset: %d", field.Name, field.Offset)
	}
}
