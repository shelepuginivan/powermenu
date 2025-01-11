// Package icons provides default icons for powermenu.
//
// Icons used are part of [Carbon] icon pack.
//
// [Carbon]: https://github.com/carbon-design-system/carbon
package icons

import _ "embed"

//go:embed svg/hibernate.svg
var Hibernate []byte

//go:embed svg/lock.svg
var Lock []byte

//go:embed svg/logout.svg
var Logout []byte

//go:embed svg/poweroff.svg
var PowerOff []byte

//go:embed svg/reboot.svg
var Reboot []byte

//go:embed svg/suspend.svg
var Suspend []byte
