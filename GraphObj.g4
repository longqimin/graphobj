grammar GraphObj;

Name: ~[{} \t\r\n,]+;
WS: [ \t\r\n,]* -> skip;
Attr : '[' .* ']' ;

graph : WS? Name Attr? WS? '{' graph* '}';

graphs : graph+;