package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func hex(str string) string {
	values, _ := strconv.ParseInt(str, 16, 64)
	return strconv.Itoa(int(values))
}

func bin(str string) string {
	values, _ := strconv.ParseInt(str, 2, 64)
	return strconv.Itoa(int(values))
}

func up(str string) string {
	return strings.ToUpper(str)
}

func low(str string) string {
	return strings.ToLower(str)
}

func capitalize(str string) string {
	return strings.Title(str)
}

func fixPuncts(line string) string {
	punctuationPattern := regexp.MustCompile(`\s*([.,;:?!]+)\s*`)
	line = punctuationPattern.ReplaceAllString(line, "$1 ") + "\n"
	return line
}

func aposthropeFix(line string) string {
	aposthropeFixer := regexp.MustCompile(`\s*'\s*`)
	// tek tırnak sorunu
	line = aposthropeFixer.ReplaceAllString(line, "'")
	count := 0
	preIndex := 0
	for i := 0; i < len(line); i++ {
		if string(line[i]) == "'" {
			count++
			if count < 2 {
				preIndex = i
			} else if count == 2 {
				line = line[:preIndex] + " " + line[preIndex:i+1] + " " + line[i+1:] // ardışık iki tırnak arası boşluk
				count = 0
				preIndex = 0
			}
		}
	}
	return strings.TrimSpace(line) // Son karakter boşluğunu kaldır
}

func readtext(inputName string) (string, error) {
	dosya, err := os.Open(inputName)
	if err != nil {
		return "", err
	}
	defer dosya.Close()

	tarayici := bufio.NewScanner(dosya)
	var metin strings.Builder
	for tarayici.Scan() {
		metin.WriteString(tarayici.Text())
	}
	return metin.String(), tarayici.Err()
}

func writetext(dosyaAdi, metin string) error {
	dosya, err := os.Create(dosyaAdi)
	if err != nil {
		fmt.Println("hata. Dosya yazilirken hata oluştu")
		return err
	}
	defer dosya.Close()
	_, err = dosya.WriteString(metin)
	return err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("HATA. ARGÜMAN SAYİSİ HATALI")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	text, err := readtext(input)
	if err != nil {
		fmt.Println("HATA. DOSYA OKUNURKEN PROBLEM OLUSTU")
		return
	}
	words := strings.Fields(text)
	res := UseFunc(words)

	err = writetext(output, res)
	if err != nil {
		fmt.Println("HATA. Yazilirken problem oluştu")
		return
	}
}

func UseFunc(sepContent []string) string {
	var res string
	for i := len(sepContent) - 1; i >= 0; i-- {
		if sepContent[i] == "(up)" {
			sepContent[i-1] = up(sepContent[i-1])
			i--
		}
		if strings.Contains(sepContent[i], ")") && sepContent[i-1] == "(up," {
			nStr := strings.Trim(sepContent[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sepContent[i-n+a-1] = up(sepContent[i-n+a-1])
			}
			i--
			continue
		}
		if sepContent[i] == "(low)" {
			sepContent[i-1] = low(sepContent[i-1])
			i--
		}
		if strings.Contains(sepContent[i], ")") && sepContent[i-1] == "(low," {
			nStr := strings.Trim(sepContent[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sepContent[i-n+a-1] = low(sepContent[i-n+a-1])
			}
			i--
			continue
		}
		if sepContent[i] == "(cap)" {
			sepContent[i-1] = capitalize(sepContent[i-1])
			i--
		}
		if strings.Contains(sepContent[i], ")") && sepContent[i-1] == "(cap," {
			nStr := strings.Trim(sepContent[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sepContent[i-n+a-1] = capitalize(sepContent[i-n+a-1])
			}
			i--
			continue
		}
		if sepContent[i] == "(hex)" {
			sepContent[i-1] = hex(sepContent[i-1])
			i--
		}
		if strings.Contains(sepContent[i], ")") && sepContent[i-1] == "(hex," {
			nStr := strings.Trim(sepContent[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sepContent[i-n+a-1] = hex(sepContent[i-n+a-1])
			}
			i--
			continue
		}
		if sepContent[i] == "(bin)" {
			sepContent[i-1] = bin(sepContent[i-1])
			i--
		}
		if strings.Contains(sepContent[i], ")") && sepContent[i-1] == "(bin," {
			nStr := strings.Trim(sepContent[i], ")")
			n, _ := strconv.Atoi(nStr)
			for a := 0; a < n; a++ {
				sepContent[i-n+a-1] = bin(sepContent[i-n+a-1])
			}
			i--
			continue
		}
		res = sepContent[i] + " " + res
	}
	res = fixPuncts(res)
	res = aposthropeFix(res) // noktalama düzeltmelerini entegre et

	kelime := strings.Fields(res)

	// a && an control

	for i := 0; i < len(kelime)-1; i++ {
		if kelime[i] == "a" && (kelime[i+1][0] == 'a' || kelime[i+1][0] == 'e' || kelime[i+1][0] == 'i' || kelime[i+1][0] == 'o' || kelime[i+1][0] == 'u' || kelime[i+1][0] == 'A' || kelime[i+1][0] == 'E' || kelime[i+1][0] == 'I' || kelime[i+1][0] == 'O' || kelime[i+1][0] == 'U') {
			kelime[i] = "an"
		} else if kelime[i] == "A" && (kelime[i+1][0] == 'a' || kelime[i+1][0] == 'e' || kelime[i+1][0] == 'i' || kelime[i+1][0] == 'o' || kelime[i+1][0] == 'u' || kelime[i+1][0] == 'A' || kelime[i+1][0] == 'E' || kelime[i+1][0] == 'I' || kelime[i+1][0] == 'O' || kelime[i+1][0] == 'U') {
			kelime[i] = "AN"
		} else {
			continue
		}
	}
	return strings.Join(kelime, " ")
}
