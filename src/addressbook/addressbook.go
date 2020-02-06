package addressbookpb

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io/ioutil"
	"log"
	"time"
)

func RunAddressBook() {
	fmt.Println("wazzzup")

	ab := newAddressBook([]*Person{
		newPerson("Tommy Vercetti", 1, "tommy.vercetti@gmail.com", []*Person_PhoneNumber{
			newPhoneNumber("123-456-7890", Person_HOME),
			newPhoneNumber("098-456-4321", Person_MOBILE),
		}, newTimestamp()),
		newPerson("Carl Johnson", 2, "carl.johnson@gmail.com", []*Person_PhoneNumber{
			newPhoneNumber("123-456-7890", Person_HOME),
			newPhoneNumber("098-456-4321", Person_MOBILE),
		}, newTimestamp()),
		newPerson("Niko Bellic", 3, "niko.bellic@gmail.com", []*Person_PhoneNumber{
			newPhoneNumber("123-456-7890", Person_HOME),
			newPhoneNumber("098-456-4321", Person_MOBILE),
		}, newTimestamp()),
	})

	fmt.Println("address book:", ab)

	if err := writeAddressBook(ab, "address-book.bin"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("addressbook successfully written to file")

	ab2, err := readAddressBook("address-book.bin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ab2)

	fmt.Println(getAsJson(ab2))
}

func newAddressBook(people []*Person) *AddressBook {
	return &AddressBook{
		People: people,
	}
}

func newPerson(name string, id int32, email string, phones []*Person_PhoneNumber, lastUpdated *timestamp.Timestamp) *Person {
	return &Person{
		Name:        name,
		Id:          id,
		Email:       email,
		Phones:      phones,
		LastUpdated: lastUpdated,
	}
}

func newPhoneNumber(number string, phoneType Person_PhoneType) *Person_PhoneNumber {
	return &Person_PhoneNumber{
		Number: number,
		Type:   phoneType,
	}
}

func newTimestamp() *timestamp.Timestamp {
	now := time.Now()

	return &timestamp.Timestamp{
		Seconds: int64(now.Second()),
		Nanos:   int32(now.Nanosecond()),
	}
}

func writeAddressBook(ab *AddressBook, filename string) error {
	out, err := proto.Marshal(ab)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		return err
	}

	return nil
}

func readAddressBook(filename string) (*AddressBook, error) {
	out, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var ab AddressBook
	if err := proto.Unmarshal(out, &ab); err != nil {
		return nil, err
	}

	return &ab, nil
}

func getAsJson(ab *AddressBook) (string, error) {
	marshaler := jsonpb.Marshaler{}

	out, err := marshaler.MarshalToString(ab)
	if err != nil {
		return "", nil
	}

	return out, nil
}