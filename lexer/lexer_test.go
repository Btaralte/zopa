package lexer

import (
	"testing"

	"github.com/Btaralte/zopa/token"
)

func TestNextToken(t *testing.T){
  input := `engemaw five = 5;
  engemaw ten = 10;
  engemaw add = khawnbawl(x, y) {
    x + y;
  };
  engemaw result = add(five, ten);
  !-/*5;
  5 < 10 > 5;
  ni(5<10){
    thupe dik;
  }
  lehlo{
    thupe diklo;
  }
  5==5;
  5!=10;
`  

tests := []struct {
    expectedType token.TokenType
    expectedLiteral string
  }{
    {token.ENGEMAW, "engemaw"},
    {token.IDENT, "five"},
    {token.ASSIGN, "="},
    {token.INT, "5"},
    {token.SEMICOLON, ";"},
    {token.ENGEMAW, "engemaw"},
    {token.IDENT, "ten"},
    {token.ASSIGN, "="},
    {token.INT, "10"},
    {token.SEMICOLON, ";"},
    {token.ENGEMAW, "engemaw"},
    {token.IDENT, "add"},
    {token.ASSIGN, "="},
    {token.KHAWNBAWL, "khawnbawl"},
    {token.LPAREN, "("},
    {token.IDENT, "x"},
    {token.COMMA, ","},
    {token.IDENT, "y"},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.IDENT, "x"},
    {token.PLUS, "+"},
    {token.IDENT, "y"},
    {token.SEMICOLON, ";"},
    {token.RBRACE, "}"},
    {token.SEMICOLON, ";"},
    {token.ENGEMAW, "engemaw"},
    {token.IDENT, "result"},
    {token.ASSIGN, "="},
    {token.IDENT, "add"},
    {token.LPAREN, "("},
    {token.IDENT, "five"},
    {token.COMMA, ","},
    {token.IDENT, "ten"},
    {token.RPAREN, ")"},
    {token.SEMICOLON, ";"},
    {token.BANG,"!"},
    {token.MINUS,"-"},
    {token.SLASH,"/"},
    {token.ASTERISK,"*"},
    {token.INT,"5"},
    {token.SEMICOLON,";"},
    {token.INT,"5"},
    {token.LT,"<"},
    {token.INT,"10"},
    {token.GT,">"},
    {token.INT,"5"},
    {token.SEMICOLON,";"},
    {token.NI,"ni"},
    {token.LPAREN,"("},
    {token.INT,"5"},
    {token.LT,"<"},
    {token.INT,"10"},
    {token.RPAREN,")"},
    {token.LBRACE,"{"},
    {token.THUPE,"thupe"},
    {token.DIK,"dik"},
    {token.SEMICOLON,";"},
    {token.RBRACE,"}"},
    {token.LEHLO,"lehlo"},
    {token.LBRACE,"{"},
    {token.THUPE,"thupe"},
    {token.DIKLO,"diklo"},
    {token.SEMICOLON,";"},
    {token.RBRACE,"}"},
    {token.INT,"5"},
    {token.EQ,"=="},
    {token.INT,"5"},
    {token.SEMICOLON,";"},
    {token.INT,"5"},
    {token.NOT_EQ,"!="},
    {token.INT,"10"},
    {token.SEMICOLON,";"},
    {token.EOF, ""},


  }
  l := New(input)
  for i, tt := range tests {
    tok := l.NextToken()
    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
      i, tt.expectedType, tok.Type)
    }
    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
      i, tt.expectedLiteral, tok.Literal)
    }
}

}




