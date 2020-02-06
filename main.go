package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	addressbookpb "github.com/protobuffers-golang/src/addressbook"
	complexpb "github.com/protobuffers-golang/src/complex"
	enumpb "github.com/protobuffers-golang/src/enum_example"
	"github.com/protobuffers-golang/src/simple"
	"io/ioutil"
	"log"
)

func doSimple() *simplepb.SimpleMessage{
	sm := simplepb.SimpleMessage{
		Id:                   12345,
		IsSimple:             true,
		Name:                 "my simple message",
		SampleList:           []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())

	return &sm
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("cannot serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatal("cannot write to file", err)
		return err
	}

	fmt.Println("data has been written")

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal("cannot read file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatal("cannot unmarshal data into the pb", err)
		return err
	}

	return nil
}

func readAndWriteDemo(sm proto.Message) {
	_ = writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	_ = readFromFile("simple.bin", sm2)
	fmt.Println("read the content:", sm2)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatal("cannot convert to JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatal("cannot unmarshal JSON into the pb struct", err)
	}
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)

	fmt.Println("successfully created proto struct:", sm2)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:                   42,
		DayOfTheWeek:         enumpb.DayOfTheWeek_THURSDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY

	fmt.Println(em)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy:             &complexpb.DummyMessage{
			Id:                   1,
			Name:                 "First message",
		},
		MultipleDummy:        []*complexpb.DummyMessage{
			{
				Id:                   2,
				Name:                 "Second message",
			},
			{
				Id:                   3,
				Name:                 "Third message",
			},
		},
	}

	fmt.Println(cm)
}



func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()

	addressbookpb.RunAddressBook()
}
