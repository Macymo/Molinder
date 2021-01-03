package server

// Gender wraps int and specifies whether someone is male of female
type Gender int

func (g Gender) String() string {
	switch g {
	case MALE:
		return "male"
	case FEMALE:
		return "female"
	default:
		return "unknown"
	}
}

const (
	MALE   Gender = iota // 0
	FEMALE               // 1
)

type Person struct {
	UID       string
	FirstName string
	LastName  string
	Username  string
	Birthday  int
	Gender    Gender
}

type PersonStorage interface {
	Get(uid string) *Person
	Delete(uid string)
	Create(p *Person)
}

type InMemoryPersonStore struct {
	store []*Person
}

func (i *InMemoryPersonStore) Get(uid string) *Person {
	return nil
}

func (i *InMemoryPersonStore) Delete(uid string) {

}

func (i *InMemoryPersonStore) Create(p *Person) {

}

type PersonServer struct {
}
