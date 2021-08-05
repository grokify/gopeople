package contacts

import (
	"strings"

	"github.com/huandu/xstrings"
)

type Person struct {
	Name         Name
	Organization Organization
	Email        string
	LinkedIn     string
	Location     Location
}

func (p Person) Finalize() {
	p.Name.TrimSpace()
	p.Organization.Finalize()
}

func (p Person) OrganizationEmailNice() string {
	p.Name.TrimSpace()
	emailMore := []string{}
	displayName := []string{}
	if len(p.Name.Last) > 0 {
		displayName = append(displayName, p.Name.Last)
	}
	if len(p.Name.First) > 0 {
		displayName = append(displayName, p.Name.First)
	}
	if len(displayName) > 0 {
		emailMore = append(emailMore, "\""+
			xstrings.Translate(strings.Join(displayName, ", "), "\"", "")+
			"\"")
	}
	if len(p.Organization.Email) > 0 {
		emailMore = append(emailMore, "<"+p.Organization.Email+">")
	}
	return strings.Join(emailMore, " ")
}

type Name struct {
	First   string
	Middle  string
	Last    string
	Pronoun string
}

func (n Name) TrimSpace() {
	n.First = strings.TrimSpace(n.First)
	n.Middle = strings.TrimSpace(n.Middle)
	n.Last = strings.TrimSpace(n.Last)
	n.Pronoun = strings.TrimSpace(n.Pronoun)
}

type Organization struct {
	Name       string
	Title      string
	Area       string
	Level      string
	Department string
	Email      string
}

func (o Organization) Finalize() {
	o.Email = strings.ToLower(strings.TrimSpace(o.Email))
}

type Location struct {
	Unlocode5 string
	City      string
	CityASCII string
	State     string
	Country   string
}
