grammar Calc;   // 声明

// 语法规则
body : expr EOF ;
expr : '(' expr ')'
     | expr md_operator expr
     | expr as_operator expr
     | num
     ;
md_operator : '*'|'/';
as_operator : '+'|'-';

num : INT;

// 词法规则
INT : '0'..'9' + ;
WS : [ \t\n\r]+ -> skip;
