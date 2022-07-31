// License Mechanism provides licenses for concurrent users of the software
// packages. If there are no more copies to be given than the purchased
// licenses, it does not provide any new ones. Only the Licenses that are
// discarded by a user can be used by a new user.

package main

import "fmt"

func main() {
	// number of purchased licenses by company.
	purchasedLicence := new(int)

	// let's say the company purchased 5 licenses.
	*purchasedLicence = 5

	// check the number of licenses in use by concurrent users.
	concurrentUsers(purchasedLicence, &License{})
}

// create an interface that will be used by License and LicenseProxy structs.
type LicenseKey interface {
	Authorization()
}

type License struct{}

type LicenseProxy struct {
	license    License
	LicenseReq *LicenseReq
}

// adding license requirements, for now, just the number of users is enough.
type LicenseReq struct {
	NoLicense int
}

// create a method for the License struct that acts as a real license object,
// which will be used by proxyLicense struct not directly.
func (l *License) Authorization() {
	fmt.Println("Licence is being used.")
}

// create a method for LicenseProxy struct to check more detailed information
// for providing controlled access to the real license struct.
func (l *LicenseProxy) Authorization() {
	if l.LicenseReq.NoLicense < 1 {
		fmt.Println("No more licenses are available.")
	} else {
		l.LicenseReq.NoLicense -= 1
		fmt.Printf("%v licenses remained.\n", l.LicenseReq.NoLicense)
	}
}

// create license proxy.
func NewLicenseProxy(LicenseReq *LicenseReq) *LicenseProxy {
	return &LicenseProxy{License{}, LicenseReq}
}

// check the number of concurrent users of the licenses.
func concurrentUsers(purchasedLicence *int, l *License) {
	// create a proxy with number of licenses.
	license := NewLicenseProxy(&LicenseReq{NoLicense: *purchasedLicence})

	// to show there are no more licenses will be given to new users
	// after all licenses are in use.
	for i := -3; i < *purchasedLicence; i++ {
		license.Authorization()
	}
}
