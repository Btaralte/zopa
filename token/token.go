package token

const (
  PHALLOH = "PHALLOH"
  EOF = "EOF"

  //Identifiers + literals

  IDENT = "IDENT" //add,x,y,..
  INT = "INT" //7 
 
  //Operators
  ASSIGN = "="
  PLUS = "+"
  MINUS = "-"
  BANG = "!"
  ASTERISK = "*"
  SLASH = "/"
  LT = "<"
  GT = ">"

  //Delimiters
  COMMA = ","
  SEMICOLON = ";"

  LPAREN = "("
  RPAREN = ")"
  LBRACE = "{"
  RBRACE = "}"

  //Keywords
  KHAWNBAWL  = "KHAWNBAWL"
  ENGEMAW= "ENGEMAW"
  DIK = "DIK"
  DIKLO = "DIKLO"
  NI = "NI"
  LEHLO = "LEHLO"
  THUPE = "THUPE"

  EQ = "=="
  NOT_EQ = "!="

)

type TokenType string 
type Token struct{
  Type TokenType
  Literal string
}



var keywords = map[string]TokenType{
  "khawnbawl":KHAWNBAWL,
  "engemaw":ENGEMAW,
  "dik":DIK,
  "diklo":DIKLO,
  "ni":NI,
  "lehlo":LEHLO,
  "thupe":THUPE,
}

func LookupIdent(ident string)TokenType  {
  if tok,ok := keywords[ident];ok {
    return tok
  }
  return IDENT
}



