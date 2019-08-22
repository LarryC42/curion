package fixed

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Fixed is a four decimal place fixed data type
type Fixed struct {
	value int64
}

// FromInt creates a Fixed from an int
func FromInt(v int) Fixed {
	return Fixed{value: int64(v) * 10000}
}

// FromInt64 creates a Fixed from an int64
func FromInt64(v int64) Fixed {
	return Fixed{value: v * 10000}
}

// FromFloat64 creats a Fixed from a float64
func FromFloat64(v float64) Fixed {
	return Fixed{value: int64(v * 10000.0)}
}

// FromString creates a Fixed from a string
func FromString(value string) (Fixed, error) {
	pieces := strings.Split(value, ".")
	if len(pieces) == 1 {
		pieces = append(pieces, "0000")
	}
	if len([]rune(pieces[1])) < 4 {
		pieces[1] = string([]rune(pieces[1] + "0000")[0:4])
	}
	whole, err := strconv.ParseInt(pieces[0], 10, 64)
	if err != nil {
		return Fixed{}, err
	}
	frac, err := strconv.ParseInt(pieces[1], 10, 64)
	if err != nil {
		return Fixed{}, err
	}
	return Fixed{value: whole*10000 + frac}, nil
}

// FromFixed clones a Fixed
func FromFixed(v Fixed) Fixed {
	return Fixed{value: v.value}
}

func toThousands(n int64, sep byte) string {
	in := strconv.FormatInt(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = sep
		}
	}
}

// ToString outputs fixed as string with decimals, optional separator left padded to width
func (f *Fixed) ToString(decimals int, sep byte, width int) string {
	frac := f.value % 10000
	whole := f.value / 10000
	if decimals > 4 {
		decimals = 4
	}
	if decimals < 0 {
		decimals = 0
	}

	for places := 4 - decimals; places > 0; places-- {
		frac /= 10
	}
	fr := ""
	if decimals > 0 {
		format := fmt.Sprintf("%s%dd", "%", decimals)
		fr = fmt.Sprintf(format, frac)
	}
	wh := ""
	if sep != 0 {
		wh = toThousands(whole, sep)
	} else {
		wh = strconv.FormatInt(whole, 10)
	}
	out := wh + fr

	if width > 0 {
		out = strings.Repeat(" ", width) + out
		r := []rune(out)
		out = string(r[len(r)-width : len(r)])
	}
	return wh + fr
}

// Compare compares two Fixed
func (f *Fixed) Compare(b Fixed) int {
	if f.value < b.value {
		return -1
	}
	if f.value > b.value {
		return 1
	}
	return 0
}

// Add adds two Fixed
func (f *Fixed) Add(b Fixed) Fixed {
	return Fixed{value: f.value + b.value}
}

// Sub subtracts two Fixed
func (f *Fixed) Sub(b Fixed) Fixed {
	return Fixed{value: f.value - b.value}
}

// Mul multiplies two Fixed
func (f *Fixed) Mul(b Fixed) Fixed {
	return Fixed{value: f.value * b.value / 10000}
}

// Div divides two Fixed
func (f *Fixed) Div(b Fixed) Fixed {
	return Fixed{value: f.value * 10000 / b.value}
}

// Round rounds a Fixed to decimal places
func (f *Fixed) Round(decimals int) error {
	if decimals < 0 || decimals > 4 {
		return errors.New("Decimals must be between 0 and 4")
	}
	var round int64
	for i := decimals; i < 4; i++ {
		if i == 3 {
			round = f.value % 10
		}
		f.value /= 10
	}
	if round >= 5 {
		f.value++
	}
	for i := decimals; i < 4; i++ {
		f.value *= 10
	}
	return nil
}

// Fix truncates a Fixed to decimal places
func (f *Fixed) Fix(decimals int) error {
	if decimals < 0 || decimals > 4 {
		return errors.New("Decimals must be between 0 and 4")
	}
	for i := decimals; i < 4; i++ {
		f.value /= 10
	}
	for i := decimals; i < 4; i++ {
		f.value *= 10
	}
	return nil
}

// ToWords outputs Fixed as English words
func (f *Fixed) ToWords() string {
	v := f.value / 10000
	groups := []string{"", "Thousand", "Million", "Billion", "Trillion", "Quadrillion"}
	output := ""
	for a := 5; a >= 0; a-- {
		var p int64 = 1
		for j := 0; j < a; j++ {
			p *= 1000
		}
		a2 := v / p
		a1 := a2 * p
		if a2 > 0 {
			output += threeDigitGroup(a2)
			if a > 0 {
				output += " " + groups[a]
			}
			v -= a1
			if v > 0 {
				output += " "
			}
		}
	}
	return output
}

// ToDollars returns Fixed as a dollar amount
func (f *Fixed) ToDollars() string {
	return fmt.Sprint(f.ToWords(), "Dollars and", f.cents())
}

func (f *Fixed) cents() string {
	v := (f.value % 10000) / 100
	if v == 0 {
		return "no/100***"
	}
	return fmt.Sprintf("%d/100***", v)
}

func threeDigitGroup(v int64) string {
	ones := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
		"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	output := ""
	if v >= 100 {
		output = ones[v/100] + " Hundred "
	}
	v %= 100
	if v < 20 {
		output += ones[v]
	} else {
		output += fmt.Sprint(tens[v/10], ones[v%10])
	}
	return output
}

// 		string output;
// 		if (v >= 100)
// 		{
// 			output = ones[cast(uint)(v / 100)] ~ " Hundred ";
// 		}
// 		v %= 100;
// 		if (v < 20)
// 		{
// 			output ~= ones[cast(uint) v];
// 		}
// 		else
// 		{
// 			output ~= text(tens[cast(uint)(v / 10)], " ", ones[cast(uint)(v % 10)]);
// 		}
// 		return output;
// 	}
// }
