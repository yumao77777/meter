package centiconv

import "fmt"

type inch float64
type meter float64

func (in inch) String() string { return fmt.Sprintf("%g英尺", in) }
func (me meter) String() string { return fmt.Sprintf("%g英尺", me) }
func IToM(in inch) meter { return meter(in * 0.3048) }
func MToI(me meter) inch { return inch(me * 3.2808) }
