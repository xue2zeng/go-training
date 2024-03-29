// Copyright © 2018 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

//!+

// 2.2 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的 话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对 应英尺和米，重量单位可以对应磅和公斤等
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Feet float64
type Meter float64

type Pound float64
type Kilogram float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (f Feet) String() string       { return fmt.Sprintf("%g feet", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g meter", m) }
func (p Pound) String() string      { return fmt.Sprintf("%g pound", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%g kilogram", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }

func MToF(m Meter) Feet { return Feet(m / 0.3048) }

func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

func KToP(k Kilogram) Pound { return Pound(k / 0.453592) }
