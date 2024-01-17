package lexer

import (
	"bufio"
	"fmt"
	mytoken "lexer/pkg/token"
	"os"
	"unicode"
)

// Parse выполняет лексический анализ содержимого файла и выводит литералы найденных токенов.
func Parse(file *os.File) []mytoken.Token {
	var result []mytoken.Token
	mytoken.InitReservedWords()
	scanner := bufio.NewScanner(file)

	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()

		for i := 0; i < len(line); {
			token := mytoken.Token{}
			token.LineNumber = uint(lineNumber)
			token.ColumnNumber = uint(i + 1)

			currentChar := line[i]
			nextChar := byte(0)
			if i < len(line)-1 {
				nextChar = line[i+1]
			}
			
			switch currentChar {
			case ' ':
				token.Type = mytoken.Blank
			case ',':
				token.Type = mytoken.Comma
				token.Lexeme = ","
			case ';':
				token.Type = mytoken.Semicolon
				token.Lexeme = ";"
			case '\'':
				token.Type, token.Lexeme = parseStringLiteral(line, &i)
			case '+':
				token.Type = mytoken.Addition
				token.Lexeme = "+"
			case '-':
				token.Type = mytoken.Subtraction
				token.Lexeme = "-"
			case '*':
				token.Type = mytoken.Multiplication
				token.Lexeme = "*"
			case '/':
				if nextChar == '/' {
					token.Type, token.Lexeme = parseOneLineComment(line, &i)
					break
				}
				if nextChar == '*' {
					token.Type, token.Lexeme = parseMultilineComment(scanner, line, &i)
					break
				}
				token.Type = mytoken.Division
				token.Lexeme = "/"
			case '(':
				token.Type = mytoken.OpeningParenthesis
				token.Lexeme = "("
			case ')':
				token.Type = mytoken.ClosingParenthesis
				token.Lexeme = ")"
			case '{':
				token.Type = mytoken.OpeningBrace
				token.Lexeme = "{"
			case '}':
				token.Type = mytoken.ClosingBrace
				token.Lexeme = "}"
			case '=':
				if nextChar == '=' {
					token.Type = mytoken.Equal
					token.Lexeme = "=="
					i++
				} else {
					token.Type = mytoken.Assignment
					token.Lexeme = "="
				}
			case '!':
				if nextChar == '=' {
					token.Type = mytoken.NotEqual
					token.Lexeme = "!="
					i++
				} else {
					token.Type = mytoken.Err
				}
			case '<':
				if nextChar == '=' {
					token.Type = mytoken.LessOrEqual
					token.Lexeme = "<="
					i++
				} else {
					token.Type = mytoken.Less
					token.Lexeme = "<"
				}
			case '>':
				if nextChar == '=' {
					token.Type = mytoken.GreaterOrEqual
					token.Lexeme = ">="
					i++
				} else {
					token.Type = mytoken.Greater
					token.Lexeme = ">"
				}
			default:
				if isIdentifierStart(rune(currentChar)) {
					token.Type, token.Lexeme = parseIdentifier(line, &i)
					break
				}
				if isNumber(rune(currentChar)) {
					token.Type, token.Lexeme = parseNumber(line, &i)
					break
				}
				token.Type = mytoken.Err
			}
			result = append(result, token)
			i++
		}
	}

	return result
}

func isIdentifierStart(char rune) bool {
	return char == '_' || unicode.IsLetter(char) 
}

func isIdentifierSymbol(char rune) bool {
	return char == '_' || unicode.IsLetter(char) || unicode.IsDigit(char)
}

func parseIdentifier(line string, i *int) (mytoken.Type, string) {
	lexema := ""
	
	ch := rune(line[*i])

	for isIdentifierSymbol(ch) {
		lexema += string(ch)
		(*i)++
		if (*i) == len(line) {
			break
		}
		ch = rune(line[*i])
	}

	if (*i) != len(line) {
		(*i)--
	}

	if e, ok := mytoken.ReservedWords[lexema]; ok {
		return e, lexema
	}

	return mytoken.Identifier, lexema
}

func isNumber(char rune) bool {
	return unicode.IsDigit(char)
}

func parseNumber(line string, i *int) (mytoken.Type, string) {
	number := readNumber(line, i)
	return getTokenType(number), number
}

func readNumber(line string, i *int) string {
	lexema := ""

	ch := rune(line[*i])

	for isNumber(rune(line[*i])) || string(line[*i]) == "." {
		lexema += string(ch)
		(*i)++
		if (*i) == len(line) {
			break
		}
		ch = rune(line[*i])
	}

	if (*i) != len(line) {
		(*i)--
	}

	return lexema
}

func getTokenType(number string) mytoken.Type {
	if len(number) == 1 {
		return mytoken.IntegerNumber
	}

	i := 0
	
	if string(number[i]) == "0" {
		i++
		switch string(number[i]) {
		case "b":
			return parseBinary(number)
		case "x":
			return parseHex(number)
		case ".":
			return parseReal(number[i+1:])
		default:
			return parseOctal(number)
		}
	}

	return parseInteger(number)
}

func isHexDigit(ch rune) bool {
	return unicode.IsDigit(ch) || (65 <= ch && ch <= 70)
}

func parseHex(s string) mytoken.Type {
	// Реализация для шестнадцатеричного числа
	if len(s) == 0 {
		return mytoken.Err
	}

	for _, r := range s {
		if !isHexDigit(r) {
			return mytoken.Err
		}
	}

	return mytoken.HexadecimalNumber
}

func isBinaryDigit(ch rune) bool {
	return ch == 48 || ch == 49
}

func parseBinary(s string) mytoken.Type {
	// Реализация для бинарного числа
	if len(s) == 0 {
		return mytoken.Err
	}

	for _, r := range s {
		if !isBinaryDigit(r) {
			return mytoken.Err
		}
	}

	return mytoken.BinaryNumber
}

func parseReal(s string) mytoken.Type {
	// Реализация для числа с плавающей точкой
	//fmt.Print(s)
	for _, r := range s {
		if !isNumber(r)  {
			return mytoken.Err
		}
	}

	return mytoken.RealNumber
}

func isOctalDigit(ch rune) bool {
	return 48 >= ch || ch <= 55
}

func parseOctal(s string) mytoken.Type {
	// Реализация для восьмеричного числа

	for i := 0; i < len(s); i++ {
		if !isOctalDigit(rune(s[i])) {
			if s[i] == 46 {
				return parseReal(s[i+1:])
			}
			return mytoken.Err
		}
	}

	return mytoken.OctalNumber
}

func parseInteger(s string) mytoken.Type {
	// Реализация для целого числа

	for i := 0; i < len(s); i++ {
		if !isNumber(rune(s[i])) {
			if s[i] == 46 {
				return parseReal(s[i+1:])
			}
			return mytoken.Err
		}
	}

	return mytoken.IntegerNumber
}

func parseStringLiteral(line string, i *int) (mytoken.Type, string) {
	lexema := ""
	*i++

	for line[*i] != '\'' {
		lexema += string(line[*i])
		(*i)++
		if *i >= len(line) {
			return mytoken.Err, lexema
		}
	}

	return mytoken.StringLiteral, lexema
}

func parseOneLineComment(line string, i *int) (mytoken.Type, string) {
	lexema := ""
	
	for *i < len(line) {
		lexema += string(line[*i])
		(*i)++
	}

	return mytoken.OneLineComment, lexema
}

func parseMultilineComment(scanner *bufio.Scanner, line string, i *int) (mytoken.Type, string) {
	 lexeme := ""

    for *i < len(line) {
        lexeme += string(line[*i])
        if *i+1 < len(line) && line[*i] == '*' && line[*i+1] == '/' {
            *i += 2
            return mytoken.MultilineComment, lexeme
        }
        (*i)++
    }

	*i = 0

	for scanner.Scan() {
        line := scanner.Text()
        lexeme +=  "\n" + line 

        for *i < len(line) {
            if *i+1 < len(line) && line[*i] == '*' && line[*i+1] == '/' {
                *i += 2
                return mytoken.MultilineComment, lexeme
            }
            (*i)++
        }

        *i = 0
    }

    return mytoken.Err, lexeme
}

// Print выводит результат работы лексера
func Print(tokens []mytoken.Token) {
	mytoken.InitReservedType()

	for _, t := range tokens {
		fmt.Printf("Лексема: %s,\n \tТип: %v,\n \tНомер строки %d,\n \tНомер позиции %d\n", t.Lexeme, mytoken.ReservedType[t.Type], t.LineNumber, t.ColumnNumber)
	}
}

