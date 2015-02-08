package intromain

import ( "fmt" ; log "github.com/alecthomas/log4go" )

// Person ...
type Person struct {
	origName string
	origAge int
}

func NewPerson(name string, age int) *Person {
	pers := Person{origName: name, origAge: age}
	log.Debug("NewPerson()")
	
	return &pers //&Person{origName: "I.M. Computer", origAge: 35}
}

func (p Person) GetName() string { return p.origName }
func (p Person) GetAge() int { return p.origAge }

func (p *Person) SetName(name string) () { p.origName = name }
func (p *Person) SetAge(age int) () { p.origAge = age }

func (p Person) ToString() string {
    var res string = fmt.Sprintf("Person{name: %s; age: %d}", 
    	p.origName, p.origAge)
    return res
}
