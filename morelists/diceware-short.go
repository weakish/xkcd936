package morelists

import _ "embed"
import "strings"

var DicewareShort = strings.Split(strings.TrimSpace(diceware2k), "\n")

// This list is derived from the diceware8k list (see also diceware.go).
// diceware8k contains 2028 words with two or three characters.
//
//     cat diceware8k.txt | egrep '^.{2,3}$' | wc -l # 2028
//
// And diceware8k contains 51 single character "words": (seperated by comma):
//
//     ; cat diceware8k.txt | egrep '^.$' | tr '\n' ','
//     a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,
//     !,#,$,%,&,(,),*,+,-,0,1,2,3,4,5,6,7,8,9,:,;,=,?,@
//
// Removing the first 26 characters (lowercase letters) and the last five characters (symbols), we get 20 words:
//
//     ; cat diceware8k.txt | egrep '^[^:;=?@a-z]$' | tr '\n' ','
//     !,#,$,%,&,(,),*,+,-,0,1,2,3,4,5,6,7,8,9
//
// That is the 2048 words of this Diceware short list.

//go:embed diceware2k.txt
var diceware2k string
