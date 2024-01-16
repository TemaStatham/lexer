// Package mytoken предоставляет определения для типов токенов и структуры Token.
package mytoken

// Type определяет тип токена.
type Type int

const (
	// Blank представляет собой токен пробела.
	Blank Type = iota // 0
	// Comment представляет собой токен комментария.
	Comment // 1
	// Comma представляет собой токен запятой.
	Comma // 2
	// Semicolon представляет собой токен точки с запятой.
	Semicolon // 3

	// Assignment представляет собой токен оператора присваивания.
	Assignment // 4
	// Identifier представляет собой токен идентификатора.
	Identifier // 5
	// StringLiteral представляет собой токен строкового литерала.
	StringLiteral // 6

	// IntType представляет собой токен ключевого слова "int".
	IntType // 7
	// DoubleType представляет собой токен ключевого слова "double".
	DoubleType // 8
	// BoolType представляет собой токен ключевого слова "bool".
	BoolType // 9
	// StringType представляет собой токен ключевого слова "string".
	StringType // 10

	// IntegerNumber представляет собой токен целочисленного литерала.
	IntegerNumber // 11
	// RealNumber представляет собой токен литерала с плавающей точкой.
	RealNumber // 12
	// BinaryNumber представляет собой токен бинарного литерала.
	BinaryNumber // 13
	// OctalNumber представляет собой токен восьмеричного литерала.
	OctalNumber // 14
	// HexadecimalNumber представляет собой токен шестнадцатеричного литерала.
	HexadecimalNumber // 15

	// Addition представляет собой токен оператора сложения.
	Addition // 16
	// Subtraction представляет собой токен оператора вычитания.
	Subtraction // 17
	// Multiplication представляет собой токен оператора умножения.
	Multiplication // 18
	// Division представляет собой токен оператора деления.
	Division // 19

	// IfKeyword представляет собой токен ключевого слова "if".
	IfKeyword // 20
	// ElseKeyword представляет собой токен ключевого слова "else".
	ElseKeyword // 21
	// WhileKeyword представляет собой токен ключевого слова "while".
	WhileKeyword // 22
	// ForKeyword представляет собой токен ключевого слова "for".
	ForKeyword // 23
	// ReadKeyword представляет собой токен ключевого слова "read".
	ReadKeyword // 24
	// PrintKeyword представляет собой токен ключевого слова "print".
	PrintKeyword // 25

	// OpeningParenthesis представляет собой токен открывающей скобки.
	OpeningParenthesis // 26
	// ClosingParenthesis представляет собой токен закрывающей скобки.
	ClosingParenthesis // 27
	// OpeningBrace представляет собой токен открывающей фигурной скобки.
	OpeningBrace // 28
	// ClosingBrace представляет собой токен закрывающей фигурной скобки.
	ClosingBrace // 29

	// Equal представляет собой токен оператора равенства.
	Equal // 30
	// NotEqual представляет собой токен оператора неравенства.
	NotEqual // 31

	// Less представляет собой токен оператора "меньше".
	Less // 32
	// Greater представляет собой токен оператора "больше".
	Greater // 33
	// LessOrEqual представляет собой токен оператора "меньше или равно".
	LessOrEqual // 34
	// GreaterOrEqual представляет собой токен оператора "больше или равно".
	GreaterOrEqual // 35
	
	// Err представляет собой токен для обозначения ошибок.
	Err // 36

	// OneLineComment представляет собой однострочный комменатрий.
	OneLineComment // 37

	// MultilineComment представляет собой многострочный комментарий.
	MultilineComment // 38
)

// ReservedWords представляет собой резервированные слова
var ReservedWords map[string]Type
// ReservedType представляет собой резервированные типы
var ReservedType map[Type]string

// Token представляет собой единичный токен с дополнительной информацией.
type Token struct {
	Type         Type   // Тип токена
	Lexeme       string // Лексема (текст токена)
	LineNumber   uint   // Номер строки в исходном коде
	ColumnNumber uint   // Номер столбца в исходном коде
}

// InitReservedWords инициализирует начальными значениями ReservedWords
func InitReservedWords() {
	ReservedWords = make(map[string]Type)

	ReservedWords["int"] = IntType
	ReservedWords["double"] = DoubleType
	ReservedWords["bool"] = BoolType
	ReservedWords["string"] = StringType

	ReservedWords["if"] = IfKeyword
	ReservedWords["else"] = ElseKeyword
	ReservedWords["while"] = WhileKeyword
	ReservedWords["for"] = ForKeyword
	ReservedWords["read"] = ReadKeyword
	ReservedWords["print"] = PrintKeyword
}

// InitReservedType инициализирует начальными значениями ReservedType
func InitReservedType() {
	ReservedType = make(map[Type]string)
	ReservedType[Blank] = "blank"
	ReservedType[Comment] = "comment"
	ReservedType[Comma] = "comma"
	ReservedType[Semicolon] = "semicolon"

	ReservedType[Assignment] = "assignment"
	ReservedType[Identifier] = "identifier"
	ReservedType[StringLiteral] = "stringLiteral"

	ReservedType[IntType] = "int"
	ReservedType[DoubleType] = "double"
	ReservedType[BoolType] = "bool"
	ReservedType[StringType] = "string"

	ReservedType[IntegerNumber] = "integerNumber"
	ReservedType[RealNumber] = "realNumber"
	ReservedType[BinaryNumber] = "binaryNumber"
	ReservedType[OctalNumber] = "octalNumber"
	ReservedType[HexadecimalNumber] = "hexadecimalNumber"

	ReservedType[Addition] = "addition"
	ReservedType[Subtraction] = "subtraction"
	ReservedType[Multiplication] = "multiplication"
	ReservedType[Division] = "division"

	ReservedType[IfKeyword] = "if"
	ReservedType[ElseKeyword] = "else"
	ReservedType[WhileKeyword] = "while"
	ReservedType[ForKeyword] = "for"
	ReservedType[ReadKeyword] = "read"
	ReservedType[PrintKeyword] = "print"

	ReservedType[OpeningParenthesis] = "openingParenthesis"
	ReservedType[ClosingParenthesis] = "closingParenthesis"
	ReservedType[OpeningBrace] = "openingBrace"
	ReservedType[ClosingBrace] = "closingBrace"

	ReservedType[Equal] = "equal"
	ReservedType[NotEqual] = "notEqual"

	ReservedType[Less] = "less"
	ReservedType[Greater] = "greater"
	ReservedType[LessOrEqual] = "lessOrEqual"
	ReservedType[GreaterOrEqual] = "greaterOrEqual"

	ReservedType[Err] = "err"

	ReservedType[OneLineComment] = "oneLineComment"
	ReservedType[MultilineComment] = "multilineComment"
}
