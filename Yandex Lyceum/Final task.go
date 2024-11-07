package main

import (
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	ger := ""
	var exp, d []string
	var c, b, yd int
	for i := 0; i < len(expression); i++ {
		if string(expression[i]) != "/" && string(expression[i]) != "*" && string(expression[i]) != "+" && string(expression[i]) != "-" && string(expression[i]) != "(" && string(expression[i]) != ")" {
			ger += string(expression[i])
		} else {
			if ger != "" {
				exp = append(exp, ger)
				ger = ""
			}
			exp = append(exp, string(expression[i]))
		}
	}
	exp = append(exp, ger)
	if len(exp) <= 1 {
		if len(exp) == 0 {
			return 0, fmt.Errorf("Empty string")
		} else {
			a, b := strconv.ParseFloat(string(exp[0]), 64)
			if b != nil {
				return 0, fmt.Errorf("Error in writing")
			} else {
				return a, nil
			}
		}
	}
	//СКОБКИ!!!!!!!!
	for i := 0; i < len(exp); i++ {
		if exp[i] == "(" {
			c = i
		} else if exp[i] == ")" {
			b = i
		}
	}
	if b != 0 {
		d = exp[c+1 : b]
	}
	for i := 0; i < len(d); i++ {
		if d[i] == "/" || d[i] == "*" {
			yd++
		}
	}
	tf, err := FindErrors(c, b, exp, d)
	if tf {
		return 0, err
	}
	a, f := 1.0, 1.0
	for len(d) > 1 {
		if len(d) == 0 {
			break
		}
		for i := 0; i < len(d); i++ {
			if d[i] == "/" {
				a, _ = strconv.ParseFloat(d[i-1], 64)
				f, _ = strconv.ParseFloat(d[i+1], 64)
				ger = strconv.FormatFloat(float64(a)/float64(f), 'f', 3, 64)
				d[i-1] = ger
				copy(d[i+1:], d[i+2:])
				copy(d[i:], d[i+1:])
				d = d[:len(d)-2]
				yd--
				i = 0
			} else if d[i] == "*" {
				a, _ = strconv.ParseFloat(d[i-1], 64)
				f, _ = strconv.ParseFloat(d[i+1], 64)
				ger = strconv.FormatFloat(float64(a)*float64(f), 'f', 3, 64)
				d[i-1] = ger
				copy(d[i+1:], d[i+2:])
				copy(d[i:], d[i+1:])
				d = d[:len(d)-2]
				yd--
				i = 0
			} else if d[i] == "-" && yd == 0 {
				a, _ = strconv.ParseFloat(d[i-1], 64)
				f, _ = strconv.ParseFloat(d[i+1], 64)
				ger = strconv.FormatFloat(float64(a)-float64(f), 'f', 3, 64)
				d[i-1] = ger
				copy(d[i+1:], d[i+2:])
				copy(d[i:], d[i+1:])
				d = d[:len(d)-2]
				i = 0
			} else if d[i] == "+" && yd == 0 {
				a, _ = strconv.ParseFloat(d[i-1], 64)
				f, _ = strconv.ParseFloat(d[i+1], 64)
				ger = strconv.FormatFloat(float64(a)+float64(f), 'f', 3, 64)
				d[i-1] = ger
				copy(d[i+1:], d[i+2:])
				copy(d[i:], d[i+1:])
				d = d[:len(d)-2]
				i = 0
			}
		}
	}
	if len(d) > 0 {
		copy(exp[c+2:], exp[b+1:])
		copy(exp[c:], exp[c+1:])
		exp = exp[:len(exp)-(b-c)]
	}
	if exp[len(exp)-1] == " " || exp[len(exp)-1] == "" {
		exp = exp[:len(exp)-1]
	}
	//ВЫРАЖЕНИЕ!!!!!!!!!!
	for i := 0; i < len(exp); i++ {
		if exp[i] == "*" || exp[i] == "/" {
			yd++
		}
	}
	for len(exp) > 1 {
		for i := 0; i < len(exp); i++ {
			if exp[i] == "/" {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				ger = strconv.FormatFloat(float64(a)/float64(f), 'f', 3, 64)
				exp[i-1] = ger
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				yd--
				i = 0
			} else if exp[i] == "*" {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				ger = strconv.FormatFloat(float64(a)*float64(f), 'f', 3, 64)
				exp[i-1] = ger
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				yd--
				i = 0
			} else if exp[i] == "-" && yd == 0 {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				ger = strconv.FormatFloat(float64(a)-float64(f), 'f', 3, 64)
				exp[i-1] = ger
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				i = 0
			} else if exp[i] == "+" && yd == 0 {
				a, _ = strconv.ParseFloat(exp[i-1], 64)
				f, _ = strconv.ParseFloat(exp[i+1], 64)
				ger = strconv.FormatFloat(float64(a)+float64(f), 'f', 3, 64)
				exp[i-1] = ger
				copy(exp[i+1:], exp[i+2:])
				copy(exp[i:], exp[i+1:])
				exp = exp[:len(exp)-2]
				i = 0
			}
		}
	}
	a, _ = strconv.ParseFloat(exp[0], 64)
	return a, nil
}
func main() {
	expression := "12/(5-5)"
	fmt.Println(Calc(expression))
}
func FindErrors(c, b int, exp, d []string) (bool, error) {
	for i := 0; i < len(exp); i++ {
		if i == len(exp)-1 {
			if exp[i] == "/" || exp[i] == "*" || exp[i] == "+" || exp[i] == "-" || exp[i] == "(" {
				return true, fmt.Errorf("Error in writing")
			} else {
				if (exp[i] == ")") && (exp[i-1] == "*" || exp[i-1] == "(" || exp[i-1] == "+" || exp[i-1] == "-" || exp[i-1] == "/") {
					return true, fmt.Errorf("Error in writing")
				} else {
					return false, fmt.Errorf("Good")
				}
			}
		}
		if i == 0 {
			if exp[i] == "(" {
				if exp[i+1] == "+" || exp[i+1] == "-" || exp[i+1] == "*" || exp[i+1] == "/" || exp[i+1] == ")" {
					return true, fmt.Errorf("Error in writing")
				} else {
					i++
				}
			}
		}
		if exp[i] == "/" || exp[i] == "*" || exp[i] == "+" || exp[i] == "-" {
			if exp[i+1] == "(" || exp[i-1] == ")" {
				continue
			}
			_, a := strconv.Atoi(exp[i+1])
			_, b := strconv.Atoi(exp[i-1])
			if exp[i+1] == "(" || exp[i] == "(" {
				if exp[i+1] == "(" {
					if b != nil {
						return true, fmt.Errorf("Error in writing")
					}
				} else if exp[i] == "(" {
					if a != nil || b == nil {
						return true, fmt.Errorf("Error in writing")
					}
				}
				continue
			}
			if a != nil || b != nil {
				return true, fmt.Errorf("Error in writing")
			}
			if exp[i] == "/" {
				if exp[i+1] == "0" {
					return true, fmt.Errorf("Divide by zero")
				}
			}
		}
	}
	return false, fmt.Errorf("Good")
}
