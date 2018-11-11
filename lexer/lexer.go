package lexer

type Lexer struct {
	input        string
	position     int  // points to current char in input
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char in examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
