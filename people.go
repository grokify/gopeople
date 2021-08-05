package contacts

import (
	"github.com/grokify/simplego/type/stringsutil"
)

type People []Person

func (p People) OrganizationEmailsNice() []string {
	emails := []string{}
	for _, person := range p {
		emails = append(emails, person.OrganizationEmailNice())
	}
	return stringsutil.SliceCondenseSpace(emails, true, true)
}

