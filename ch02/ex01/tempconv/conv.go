package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + AbsoluteZeroC) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(FToC(f) + AbsoluteZeroC) }

func KToC(k Kelvin) Celsius    { return Celsius(k) - AbsoluteZeroC }
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
