package utils

type (
	Celsius    float64
	Kelvin     float64
	Fahrenheit float64
)

func Celsius2Fahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func Celsius2Kelvin(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func Kelvin2Celsius(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func Fahrenheit2Celsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
