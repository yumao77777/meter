package centiconv

import "fmt"

type Inch float64
type Meter float64

func (in Inch) String() string { return fmt.Sprintf("%g英尺", in) }
func (me Meter) String() string { return fmt.Sprintf("%g英尺", me) }
func IToM(in Inch) Meter { return Meter(in * 0.3048) }
func MToI(me Meter) Inch { return Inch(me * 3.2808) }
