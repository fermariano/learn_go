package main

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (k Kelvin) celsius() Celsius {
	return Celsius(k - 273.15)
}

func (k Kelvin) fahrenheit() Fahrenheit {
	return Fahrenheit(k.celsius().fahrenheit())
}

func (f Fahrenheit) celsius() Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

func (f Fahrenheit) kelvin() Kelvin {
	return Kelvin(f.celsius().kelvin())
}

func (c Celsius) fahrenheit() Fahrenheit {
	return Fahrenheit((c * 9.0 / 5.0) + 32)
}

func (c Celsius) kelvin() Kelvin {
	return Kelvin(c + 273.15)
}
