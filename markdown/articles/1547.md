# 使用Flex Bison 和LLVM编写自己的编译器

**使用Flex Bison 和 LLVM编写你自己的编译器**  

原文出处：[http://gnuu.org/2009/09/18/writing-your-own-toy-compiler](http://gnuu.org/2009/09/18/writing-your-own-toy-compiler/)

## 1、介绍

我总是对编译器和语言非常感兴趣，但是兴趣并不会让你走的更远。大量的编译器的设计概念可以搞的任何一个程序员迷失在这些概念之中。不用说，我也曾今尝试过，但是并没有取得太大的成功，我以前的尝试都停留在语义分析阶段。本文的灵感主要来源于我最近一次的尝试，并且在这一次中我取得一点成就。

幸运的是，最近的几年，我参加了一些项目，这些项目给了我在建立编译器上很多有用的经验和观点。另外一件事是，我非常幸运得到[LLVM](http://llvm.org/)的帮助。对于这个工具，我不知道改怎么去形容它，但是他给我的这个编译器的确带来非常大的帮助。  



### 1.1、你为什么要阅读本文

你也许想看看我正在做的事情，但是更有可能的是，你也是和我一样对编译器和语言非常感兴趣，并且也可能遇到了一些在探索的过程中遇到了一些难题，你可能正打算解决这些难题，但是却没有发现好的资源。本文的目标就是提供这些资源，并以一种手把手的方式教你从头到尾的去创建一个具有基本功能的语言编译器。

在本文，我不会去解释一些编译器基本理论，所以你要在开始本文前去了解什么是[BNF语法](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_Form)，什么是[抽象语法树数据结构 AST data structure](https://en.wikipedia.org/wiki/Abstract_syntax_tree)，什么是基础[编译器流水线 complier pipline](https://en.wikipedia.org/wiki/Compiler)。就是说，我会把本文描述的尽量简单。本文的目的就是以一种简单易懂的方式来介绍相关编译器资源的方式来帮助那些从来没有编译器经验的人。

### 1.2、达到的成果

如果你根据文章内容一步步来，你将会得到一个能定义函数，调用函数，定义变量，给变量赋值执行基本数学操作的语言。这门语言支持两种基本类型，double和integer类型。还有一些功能还未实现，因此，你可以通过自己去实现这些功能得到你满意的功能并且能为你理解编写一个编译器提供不少的帮助。

### 1.3 问题解答

#### 1.3.1 我需要了解什么样的语言

我们使用的工具是基于C/C++的。LLVM是基于C++的，我们的这个语言也基于C++，因为C++具有很多面向对象的优点和可以被重用的STL。此外对于C，Lex和Bison都具有那些初看起来令人迷惑的语法，但是我将尽可能的去解释他。我们需要处理的语法非常小，最多就100行，因此它是比较容易理解的。

#### 1.3.2 很复杂吗？

是或否，这里面有很多的东西你需要了解，甚至多的让人感觉到害怕，但是老实说，其实这些都非常简单，我们同样会使用很多工具分解这些层次的复杂性，并使得这些内容可管理。

#### 1.3.3 完成它需要多长时间

我们将要完成的编译器花了我三天的时间。但是如果你按“follow me”的方式来完成这个编译器的话，你将会花费更少的时间。如果要全部理解这里面的内容可能会花去稍微长一点的时间，但是你至少应该在一个下午就将整个编译器运行起来。

好，如果你已经准备好，我们开始吧。

## 2、准备开始

### 2.1 构成编译器的最基本的要素

一个编译器是由一组有三个到四个组件(还有一些子组件)构成，数据以管道的方式从一个组件输入并流向下一个组件。在我们这个编译器中，可能会用到一些稍微不同的工具。下面这个图展示了我们构造一个编译器的步骤，和每个步骤中将使用的工具。

![Compiler Pipeline](http://gnuu.org/wp-content/uploads/2009/09/pipeline.png)

从上图你可以看到在Linking这一步是灰掉的。我们的语言将不支持编译器的连接(很多的语言都不支持编译器的连接)。在文法分析阶段，我们将使用开源工具Lex，即如今的[Flex](http://flex.sourceforge.net/)，文法分析一般都伴随者语法分析，我们使用的语法分析工具将会是Yacc，或者说是[Bison](http://www.gnu.org/software/bison/)，最后一旦语义分析完成，我们将遍历我们的抽象语法树，并生成我们的”bytecode 字节码”，或”机器码 matchine code”。做这一步，我们将使用[LLVM](http://llvm.org/)，它能生成快速字节码，我们将使用LLVM的JIT(Just In Tinme)来在我们的机器上编译执行它

总结一下，步骤如下：

1. **文法分析用*Flex***:将数据分隔成一个个的标记token (标示符identifiers，关键字keywords，数字numbers, 中括号brackets, 大括号braces, 等等etc.)
2. **语法分析用*Bison***: 在分析标记的时候生成抽象语法树. Bison 将会做掉几乎所有的这些工作, 我们定义好我们的抽象语法树就OK了.
3. **组装用*LLVM***: 这里我们将遍历我们的抽象语法树，并未每一个节点生成字节/机器码。 这听起来似乎很疯狂，但是这几乎就是*最简单的* 一步了.

在我们开始下一步之前，你应该准备安装好Flex,Bison和LLVM。因为我们马上就要使用到它们。

### 2.2 定义我们的语法

我们语法是我们语言中最核心的部分，我们的语法使用类似标准C的语法，因为这样的语法非常熟悉，而且简单。我们语法的一个典型的例子如下：

```c
int do_math(int a) {
  int x = a * 5 + 3
}

do_math(10)

```

看起来很简单。它和C非常相似，但是它没有使用分号做语句的分隔，同时你也会注意到我们的语法中没有return语句。这就是你可以自己实现的部分。

现在我们还没有任何机制来验证结果。我们将通过检查我们编译之后LLVM打印出的字节码验证我们程序的正确性。

## 3、 第一步，使用Flex进行文法分析

这是最简单的一步，给定语法之后，我们需要将我们的输入转换一系列内部标记token。如前所述，我们的语法具有非常基础的标记token:标示符identifier ，数字常量(整型和浮点型)，数学运算符号，括号，中括号，我们的文法定义文件称为token.l，它具有一些固定的语法。定义如下：

```
%{
#include 
#include "node.h"
#include "parser.hpp"
#define SAVE_TOKEN yylval.string = new std::string(yytext, yyleng)
#define TOKEN(t) (yylval.token = t)
extern "C" int yywrap() { }
%}

%%

[ \t\n]                 ;
[a-zA-Z_][a-zA-Z0-9_]*  SAVE_TOKEN; return TIDENTIFIER;
[0-9]+\.[0-9]*          SAVE_TOKEN; return TDOUBLE;
[0-9]+                  SAVE_TOKEN; return TINTEGER;
"="                     return TOKEN(TEQUAL);
"=="                    return TOKEN(TCEQ);
"!="                    return TOKEN(TCNE);
"<"                     return TOKEN(TCLT);
"<="                    return TOKEN(TCLE);
">"                     return TOKEN(TCGT);
">="                    return TOKEN(TCGE);
"("                     return TOKEN(TLPAREN);
")"                     return TOKEN(TRPAREN);
"{"                     return TOKEN(TLBRACE);
"}"                     return TOKEN(TRBRACE);
"."                     return TOKEN(TDOT);
","                     return TOKEN(TCOMMA);
"+"                     return TOKEN(TPLUS);
"-"                     return TOKEN(TMINUS);
"*"                     return TOKEN(TMUL);
"/"                     return TOKEN(TDIV);
.                       printf("Unknown token!\n"); yyterminate();

%%

```

在第一节(译者注：即%{%}中定义的部分)声明了一些特定的C代码。由于Bison不会去访问我门的yytext变量，我们使用宏”SAVE\_TOKEN”来保证标示符的文本和文本长度是安全的(而不是靠标记本身来保证)。第一个token告诉我们要忽略掉那些空白字符。你会注意到我们有些一些等价性比较的标记和其他。还有一些标记还没有实现，你可以非常自由的将这些标记加到你自己的编译器中去。

现在我们在这里做的是定义这些标记和他们的符号名。符号(比如TIDENTFIER)将成为我们语法中的终结符。我们只是返回它，我们从未定义它，他们是在什么地方定义的？当然是在bison语法文件中。我们包含的parser.hpp头文件将会被bison所生成，并且里面的所有符号都将被生成，并被我们在这里使用。

我们对这个token.l运行flex命令，并生成tokens.cpp文件，这个程序将会和我们的语法分析器一起编译并提供yylex()函数来识别这些标记。我们将在稍后运行这个命令，因为现在我们需要从bison那里生成头文件。

## 4、第2步 使用Bison进行语法分析

这是我们工作中最富有挑战性的一部分。生成一个正确的无二义的语法并不是一项简单的工作，要经过很多实践努力。庆幸的是，我们例子中的语法是简单而完整的。在我们实现我们的语法之前，我们需要详细的讲解一下我们的设计。

### 4.1、设计AST(抽象语法树)

语法分析的最终结果是抽象语法树AST，正如我们将看到的，Bison生成抽象语法树的最优工具；我们唯一需要做的事情就是将我们的代码插入到语法中去。

文本形式字符串，例如”int x”代表了我们语言的文本形式，和这个类似，抽象语法树AST则代表了我们语言在内存中的表现形式一样(在语言在组装成而进程码之前)。正因如此，我们要在把这些插入在语法分析中的数据结构首先设计好。这个过程是非常直接的，因为我们为语法中的每个语义单元创建了一个结构。方法声明、方法调用，变量声明，引用，这些都构成了抽象语法树的节点。我们语言的抽象语法树的节点如下图：  

![Our Toy Language AST](http://gnuu.org/wp-content/uploads/2009/09/ClassDiagram.png)  

上图的C++代码如下：  

node.h文件

```c
#include <iostream>
#include <vector>
#include <llvm/Value.h>

class CodeGenContext;
class NStatement;
class NExpression;
class NVariableDeclaration;

typedef std::vector<NStatement*> StatementList;
typedef std::vector<NExpression*> ExpressionList;
typedef std::vector<NVariableDeclaration*> VariableList;

class Node {
public:
    virtual ~Node() {}
    virtual llvm::Value* codeGen(CodeGenContext& context) { }
};

class NExpression : public Node {
};

class NStatement : public Node {
};

class NInteger : public NExpression {
public:
    long long value;
    NInteger(long long value) : value(value) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NDouble : public NExpression {
public:
    double value;
    NDouble(double value) : value(value) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NIdentifier : public NExpression {
public:
    std::string name;
    NIdentifier(const std::string& name) : name(name) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NMethodCall : public NExpression {
public:
    const NIdentifier& id;
    ExpressionList arguments;
    NMethodCall(const NIdentifier& id, ExpressionList& arguments) :
        id(id), arguments(arguments) { }
    NMethodCall(const NIdentifier& id) : id(id) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NBinaryOperator : public NExpression {
public:
    int op;
    NExpression& lhs;
    NExpression& rhs;
    NBinaryOperator(NExpression& lhs, int op, NExpression& rhs) :
        lhs(lhs), rhs(rhs), op(op) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NAssignment : public NExpression {
public:
    NIdentifier& lhs;
    NExpression& rhs;
    NAssignment(NIdentifier& lhs, NExpression& rhs) :
        lhs(lhs), rhs(rhs) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NBlock : public NExpression {
public:
    StatementList statements;
    NBlock() { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NExpressionStatement : public NStatement {
public:
    NExpression& expression;
    NExpressionStatement(NExpression& expression) :
        expression(expression) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NVariableDeclaration : public NStatement {
public:
    const NIdentifier& type;
    NIdentifier& id;
    NExpression *assignmentExpr;
    NVariableDeclaration(const NIdentifier& type, NIdentifier& id) :
        type(type), id(id) { }
    NVariableDeclaration(const NIdentifier& type, NIdentifier& id, NExpression *assignmentExpr) :
        type(type), id(id), assignmentExpr(assignmentExpr) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};

class NFunctionDeclaration : public NStatement {
public:
    const NIdentifier& type;
    const NIdentifier& id;
    VariableList arguments;
    NBlock& block;
    NFunctionDeclaration(const NIdentifier& type, const NIdentifier& id,
            const VariableList& arguments, NBlock& block) :
        type(type), id(id), arguments(arguments), block(block) { }
    virtual llvm::Value* codeGen(CodeGenContext& context);
};


```

非常的清晰明了，我们省略了getter和setter方法，这里只列出了共有成员；这些类也不需要影藏私有数据，并省略了codeGen方法。在我们导出AST成LLVM的字节码时，就需要使用到这个方法。

### 4.2、Bison介绍

bison的语法定义文件同样是由这些标记构成的最复杂的部分。这并不是说技术上有多复杂，但是我也会花一些时间来讨论一下Bison的语法细节，好，现在让我们立刻来熟悉一下Bison的语法。我们将使用基于类似于BNF的语法，使用定义的好终结符和非终结符来组成我们有效的每一个语句和表达式(这些语句和表达式就代表我们之前定义的AST节点)。例如：

```
if_stmt : IF '(' condition ')' block { /* do stuff when this rule is encountered */ }
        | IF '(' condition ')'       { ... }
        ;

```

在上面例子中，我们定义了一个if语句(如果我们支持if语句话)，它和BNF不同之处在于，每个语法后面都跟了一系列动作(在'{‘和’}’之间的内容)。这个动作将在此条语法被识别(译者注：归约)的时候被执行。这个过程将会递归地按从叶子符号到根节点符号的次序执行，在这个过程，每一个非终结符最终会被合并为一棵大的语法树。你将会看到的’$$’符号代表着当前树的跟节点(译者注：’$$’代表本条语法规则中冒号左边的部分的语义内容)。此外’$1’代表了本条规则叶子中的第一个符号(译者注：’$1’代表了本条语法规则冒号右边的内容，$1代表冒号右边的第一个符号的语义值)。在上面的例子中，当’condition’有效时我们将会把$3 赋值给$$。这个例子可以解释如何将我们AST和语法规则关联起来。我们将在每一条规则中通常赋值一个节点到$$，最后这些规则会合并成一个大的抽象语法树。Bison的部分是我们语言最复杂的部分，你需要花一点时间去理解它。此外到此为止，你还没有看到完整的代码。下面就是完整的Bison部分的代码：  

parser.y

```
%{
    #include "node.h"
    NBlock *programBlock; /* the top level root node of our final AST */

    extern int yylex();
    void yyerror(const char *s) { printf("ERROR: %s\n", s); }
%}

/* Represents the many different ways we can access our data */
%union {
    Node *node;
    NBlock *block;
    NExpression *expr;
    NStatement *stmt;
    NIdentifier *ident;
    NVariableDeclaration *var_decl;
    std::vector *varvec;
    std::vector *exprvec;
    std::string *string;
    int token;
}

/* Define our terminal symbols (tokens). This should
   match our tokens.l lex file. We also define the node type
   they represent.
 */
%token  TIDENTIFIER TINTEGER TDOUBLE
%token  TCEQ TCNE TCLT TCLE TCGT TCGE TEQUAL
%token  TLPAREN TRPAREN TLBRACE TRBRACE TCOMMA TDOT
%token  TPLUS TMINUS TMUL TDIV

/* Define the type of node our nonterminal symbols represent.
   The types refer to the %union declaration above. Ex: when
   we call an ident (defined by union type ident) we are really
   calling an (NIdentifier*). It makes the compiler happy.
 */
%type  ident
%type  numeric expr
%type  func_decl_args
%type  call_args
%type  program stmts block
%type  stmt var_decl func_decl
%type  comparison

/* Operator precedence for mathematical operators */
%left TPLUS TMINUS
%left TMUL TDIV

%start program

%%

program : stmts { programBlock = $1; }
        ;

stmts : stmt { $$ = new NBlock(); $$->statements.push_back($1); }
      | stmts stmt { $1->statements.push_back($2); }
      ;

stmt : var_decl | func_decl
     | expr { $$ = new NExpressionStatement(*$1); }
     ;

block : TLBRACE stmts TRBRACE { $$ = $2; }
      | TLBRACE TRBRACE { $$ = new NBlock(); }
      ;

var_decl : ident ident { $$ = new NVariableDeclaration(*$1, *$2); }
         | ident ident TEQUAL expr { $$ = new NVariableDeclaration(*$1, *$2, $4); }
         ;

func_decl : ident ident TLPAREN func_decl_args TRPAREN block
            { $$ = new NFunctionDeclaration(*$1, *$2, *$4, *$6); delete $4; }
          ;

func_decl_args : /*blank*/  { $$ = new VariableList(); }
          | var_decl { $$ = new VariableList(); $$->push_back($1); }
          | func_decl_args TCOMMA var_decl { $1->push_back($3); }
          ;

ident : TIDENTIFIER { $$ = new NIdentifier(*$1); delete $1; }
      ;

numeric : TINTEGER { $$ = new NInteger(atol($1->c_str())); delete $1; }
        | TDOUBLE { $$ = new NDouble(atof($1->c_str())); delete $1; }
        ;

expr : ident TEQUAL expr { $$ = new NAssignment(*$1, *$3); }
     | ident TLPAREN call_args TRPAREN { $$ = new NMethodCall(*$1, *$3); delete $3; }
     | ident { $$ = $1; }
     | numeric
     | expr comparison expr { $$ = new NBinaryOperator(*$1, $2, *$3); }
     | TLPAREN expr TRPAREN { $$ = $2; }
     ;

call_args : /*blank*/  { $$ = new ExpressionList(); }
          | expr { $$ = new ExpressionList(); $$->push_back($1); }
          | call_args TCOMMA expr  { $1->push_back($3); }
          ;

comparison : TCEQ | TCNE | TCLT | TCLE | TCGT | TCGE
           | TPLUS | TMINUS | TMUL | TDIV
           ;

%%

```
## 5、生成Flex和Bison代码

现在我们有了Flex的token.l文件和Bison的parser.y文件。我们需要将这两个文件传递给工具，并由工具来生成c++代码文件。注意Bison同时会为Flex生成parser.hpp头文件；这样做是通过-d开关实现的，这个开关是的我们的标记声明和源文件分开，这样就是的我们可以让这些token标记被其他的程序使用。下面的命令创建parser.cpp，parser.hpp和tokens.cpp源文件。

```
$ bison -d -o parser.cpp parser.y
$ lex -o tokens.cpp tokens.l

```

如果上述工作都没有出错的话，我们现在位置已经完成了我们编译器工作总量的2/3。如果你现在想测试一下我们的代码，你可以编译一个简单的main.cpp程序：

```c
#include <iostream>
#include "node.h"
extern NBlock* programBlock;
extern int yyparse();

int main(int argc, char **argv)
{
    yyparse();
    std::cout << programBlock << endl;
    return 0;
}

```

你可以编译这些文件：  

$ g++ -o parser parser.cpp tokens.cpp main.cpp  

现在你需要安装LLVM了，因为llvm::Value被node.h引用了。如果你不想这么做，只是想测试一下Flex和Bison部分，你可以注释掉node.h中codeGen()方法。

如果上述工作都完成了，你现在将有一个语法分析器，这个语法分析器将从标准输入读入，并打出在内存中代表抽象语法树跟节点的内存非零地址。

## 6、组装AST和LLVM

编译器的下一步很自然地是应该将AST转换成机器码。这意味着将每一个语义节点转换成等价的机器指令。LLVM将帮助我们把这步变得非常简单，因为LLVM将真实的指令抽象成类似AST的指令。这意味着我们真正要做的事就是将AST转换成抽象指令。  

你可以想象这个过程是从抽象语法树的根节点开始遍历每一个树上节点并产生字节码的过程。现在就是使用我们在Node中定义的codeGen方法的时候了。例如，当我们遍历NBlock代码的时候(语义上NBlock代表一组我们语言的语句的集合)，我们将调用列表中每条语句的codeGen方法。上面步骤代码类似如下的形式：

```c
Value* NBlock::codeGen(CodeGenContext& context)
{
    StatementList::const_iterator it;
    Value *last = NULL;
    for (it = statements.begin(); it != statements.end(); it++) {
        std::cout << "Generating code for " << typeid(**it).name() << endl;
        last = (**it).codeGen(context);
    }
    std::cout << "Creating block" << endl;
    return last;
}

```

我们将实现抽象语法树上所有节点的codeGen方法，然后在向下遍历树的时候调用它，并隐式的遍历我们整个抽象语法树。在这个过程中，我们在CodeGenContext类来告诉我们生成字节码的位置。

### 6.1、关于LLVM要注意的一些信息

LLVM最大的一个确定就是，你很难找到LLVM的相关文档。在线手册、教程、或其他的文档都没有及时的得到相关维护，这些文档大部分都是过期的文档。除非你去深入研究，否则你很难找到关于C++ API的信息。如果你自己安装LLVM，docs  

是最新的文档。

我发现最好学习LLVM的方法就是通过LLVM的例子去学习。在LLVM的压缩包的’example’目录下有很多快速生成字节码的例子。在[LLVM demo site](http://llvm.org/demo/)上可以将C做输入，然后生成C++ API的例子。以这些例子提供的方法，我找到了类似于int x = 5 ;的指令的生成方法。我使用这些工具实现大部分的节点。

关于LLVM的问题描述到此为止，我将在下面罗列出codegen.h和codegen.cpp的源代码

codegen.h的内容。

```c
#include <stack>
#include <llvm/Module.h>
#include <llvm/Function.h>
#include <llvm/PassManager.h>
#include <llvm/CallingConv.h>
#include <llvm/Bitcode/ReaderWriter.h>
#include <llvm/Analysis/Verifier.h>
#include <llvm/Assembly/PrintModulePass.h>
#include <llvm/Support/IRBuilder.h>
#include <llvm/ModuleProvider.h>
#include <llvm/ExecutionEngine/GenericValue.h>
#include <llvm/ExecutionEngine/JIT.h>
#include <llvm/Support/raw_ostream.h>

using namespace llvm;

class NBlock;

class CodeGenBlock {
public:
    BasicBlock *block;
    std::map<std::string , Value*> locals;
};

class CodeGenContext {
    std::stack<codegenblock  *> blocks;
    Function *mainFunction;

public:
    Module *module;
    CodeGenContext() { module = new Module("main"); }

    void generateCode(NBlock& root);
    GenericValue runCode();
    std::map<std::string , Value*>& locals() { return blocks.top()->locals; }
    BasicBlock *currentBlock() { return blocks.top()->block; }
    void pushBlock(BasicBlock *block) { blocks.push(new CodeGenBlock()); blocks.top()->block = block; }
    void popBlock() { CodeGenBlock *top = blocks.top(); blocks.pop(); delete top; }
};

```

codegen.cpp的内容。

```c
#include "node.h"
#include "codegen.h"
#include "parser.hpp"

using namespace std;

/* Compile the AST into a module */
void CodeGenContext::generateCode(NBlock& root)
{
    std::cout << "Generating code...\n";

    /* Create the top level interpreter function to call as entry */
    vector<const type*> argTypes;
    FunctionType *ftype = FunctionType::get(Type::VoidTy, argTypes, false);
    mainFunction = Function::Create(ftype, GlobalValue::InternalLinkage, "main", module);
    BasicBlock *bblock = BasicBlock::Create("entry", mainFunction, 0);

    /* Push a new variable/block context */
    pushBlock(bblock);
    root.codeGen(*this); /* emit bytecode for the toplevel block */
    ReturnInst::Create(bblock);
    popBlock();

    /* Print the bytecode in a human-readable format
       to see if our program compiled properly
     */
    std::cout << "Code is generated.\n";
    PassManager pm;
    pm.add(createPrintModulePass(&outs()));
    pm.run(*module);
}

/* Executes the AST by running the main function */
GenericValue CodeGenContext::runCode() {
    std::cout << "Running code...\n";
    ExistingModuleProvider *mp = new ExistingModuleProvider(module);
    ExecutionEngine *ee = ExecutionEngine::create(mp, false);
    vector<genericvalue> noargs;
    GenericValue v = ee->runFunction(mainFunction, noargs);
    std::cout << "Code was run.\n";
    return v;
}

/* Returns an LLVM type based on the identifier */
static const Type *typeOf(const NIdentifier& type)
{
    if (type.name.compare("int") == 0) {
        return Type::Int64Ty;
    }
    else if (type.name.compare("double") == 0) {
        return Type::FP128Ty;
    }
    return Type::VoidTy;
}

/* -- Code Generation -- */

Value* NInteger::codeGen(CodeGenContext& context)
{
    std::cout << "Creating integer: " << value << endl;
    return ConstantInt::get(Type::Int64Ty, value, true);
}

Value* NDouble::codeGen(CodeGenContext& context)
{
    std::cout << "Creating double: " << value << endl;
    return ConstantFP::get(Type::FP128Ty, value);
}

Value* NIdentifier::codeGen(CodeGenContext& context)
{
    std::cout << "Creating identifier reference: " << name << endl;
    if (context.locals().find(name) == context.locals().end()) {
        std::cerr << "undeclared variable " << name << endl;
        return NULL;
    }
    return new LoadInst(context.locals()[name], "", false, context.currentBlock());
}

Value* NMethodCall::codeGen(CodeGenContext& context)
{
    Function *function = context.module->getFunction(id.name.c_str());
    if (function == NULL) {
        std::cerr << "no such function " << id.name << endl;
    }
    std::vector<value *> args;
    ExpressionList::const_iterator it;
    for (it = arguments.begin(); it != arguments.end(); it++) {
        args.push_back((**it).codeGen(context));
    }
    CallInst *call = CallInst::Create(function, args.begin(), args.end(), "", context.currentBlock());
    std::cout << "Creating method call: " << id.name << endl;
    return call;
}

Value* NBinaryOperator::codeGen(CodeGenContext& context)
{
    std::cout << "Creating binary operation " << op << endl;
    Instruction::BinaryOps instr;
    switch (op) {
        case TPLUS:     instr = Instruction::Add; goto math;
        case TMINUS:    instr = Instruction::Sub; goto math;
        case TMUL:      instr = Instruction::Mul; goto math;
        case TDIV:      instr = Instruction::SDiv; goto math;

        /* TODO comparison */
    }

    return NULL;
math:
    return BinaryOperator::Create(instr, lhs.codeGen(context),
        rhs.codeGen(context), "", context.currentBlock());
}

Value* NAssignment::codeGen(CodeGenContext& context)
{
    std::cout << "Creating assignment for " << lhs.name << endl;
    if (context.locals().find(lhs.name) == context.locals().end()) {
        std::cerr << "undeclared variable " << lhs.name << endl;
        return NULL;
    }
    return new StoreInst(rhs.codeGen(context), context.locals()[lhs.name], false, context.currentBlock());
}

Value* NBlock::codeGen(CodeGenContext& context)
{
    StatementList::const_iterator it;
    Value *last = NULL;
    for (it = statements.begin(); it != statements.end(); it++) {
        std::cout << "Generating code for " << typeid(**it).name() << endl;
        last = (**it).codeGen(context);
    }
    std::cout << "Creating block" << endl;
    return last;
}

Value* NExpressionStatement::codeGen(CodeGenContext& context)
{
    std::cout << "Generating code for " << typeid(expression).name() << endl;
    return expression.codeGen(context);
}

Value* NVariableDeclaration::codeGen(CodeGenContext& context)
{
    std::cout << "Creating variable declaration " << type.name << " " << id.name << endl;
    AllocaInst *alloc = new AllocaInst(typeOf(type), id.name.c_str(), context.currentBlock());
    context.locals()[id.name] = alloc;
    if (assignmentExpr != NULL) {
        NAssignment assn(id, *assignmentExpr);
        assn.codeGen(context);
    }
    return alloc;
}

Value* NFunctionDeclaration::codeGen(CodeGenContext& context)
{
    vector<const type*> argTypes;
    VariableList::const_iterator it;
    for (it = arguments.begin(); it != arguments.end(); it++) {
        argTypes.push_back(typeOf((**it).type));
    }
    FunctionType *ftype = FunctionType::get(typeOf(type), argTypes, false);
    Function *function = Function::Create(ftype, GlobalValue::InternalLinkage, id.name.c_str(), context.module);
    BasicBlock *bblock = BasicBlock::Create("entry", function, 0);

    context.pushBlock(bblock);

    for (it = arguments.begin(); it != arguments.end(); it++) {
        (**it).codeGen(context);
    }

    block.codeGen(context);
    ReturnInst::Create(bblock);

    context.popBlock();
    std::cout << "Creating function: " << id.name << endl;
    return function;
}

```

上述罗列很多的代码，然而这部份代码的含义需要你自己去探索。我在这里只会提及一下你需要注意的内容：

- 我们在CodeGenContext类中使用一个语句块的栈来保存最后进入的block(因为语句都被增加到blocks中)
- 我们同样用个堆栈来保存每组语句块中的[符号表](https://en.wikipedia.org/wiki/Symbol_table)
- 我们设计的语言只会知道他当前范围内的内容.要支持“全局”上下的做法，你必须向上搜索整个堆栈中每一个语句块，知道你最后发现你匹配的符号(现在我们只是简单地搜索堆栈中最顶层的符号表)。
- 在我们进入一个语句块之前，我们需要将语句块压栈，离开语句块时将语句块出栈

剩下的内容都LLVM相关了，在这个主题上我并不是专家。但是迄今为止，我们已经有了编译和运行我们语言的所有代码。

## 7、编译和运行我们的语言

### 7.1、编译我们的语言

我们已经有了代码，现在我们怎么运行它？LLVM有着非常复杂的联接link，幸运的是，如果你是自己安装的LLVM，那么你就应该有一个llvm-config二进制程序，这个程序返回你需要的所有编译和联接选项。  

我们也要同时更新我们的main.cpp的内容使之可以编译和运行我们的代码，这次我们main.cpp的内容如下：

```c
#include <iostream>
#include "codegen.h"
#include "node.h"

using namespace std;

extern int yyparse();
extern NBlock* programBlock;

int main(int argc, char **argv)
{
    yyparse();
    std::cout << programBlock << endl;

    CodeGenContext context;
    context.generateCode(*programBlock);
    context.runCode();

    return 0;
}

```

现在我们需要这样来编译这些代码  

$ g++ -o parser `llvm-config --libs core jit native --cxxflags --ldflags` \*.cpp  

你也可以编写一个Makefile来进行编译

```
all: parser

clean:
	rm parser.cpp parser.hpp parser tokens.cpp

parser.cpp: parser.y
	bison -d -o $@ $^

parser.hpp: parser.cpp

tokens.cpp: tokens.l parser.hpp
	lex -o $@ $^

parser: parser.cpp codegen.cpp main.cpp tokens.cpp
	g++ -o $@ llvm-config --libs core jit native --cxxflags --ldflags *.cpp

```
### 7.2、运行我们的语言

假设上述所有工作都圆满完成，那么现在你将有一个名为parser的二进制程序。运行它，还记得我们那个典型例子吗？让我们看看我们的编译器工作的如何。

```
$ echo 'int do_math(int a) { int x = a * 5 + 3 } do_math(10)' | ./parser
0x100a00f10
Generating code...
Generating code for 20NFunctionDeclaration
Creating variable declaration int a
Generating code for 20NVariableDeclaration
Creating variable declaration int x
Creating assignment for x
Creating binary operation 276
Creating binary operation 274
Creating integer: 3
Creating integer: 5
Creating identifier reference: a
Creating block
Creating function: do_math
Generating code for 20NExpressionStatement
Generating code for 11NMethodCall
Creating integer: 10
Creating method call: do_math
Creating block
Code is generated.
; ModuleID = 'main'

define internal void @main() {
entry:
	%0 = call i64 @do_math(i64 10)		;  [#uses=0]
	ret void
}

define internal i64 @do_math(i64) {
entry:
	%a = alloca i64		;  [#uses=1]
	%x = alloca i64		;  [#uses=1]
	%1 = add i64 5, 3	;  [#uses=1]
	%2 = load i64* %a	;  [#uses=1]
	%3 = mul i64 %2, %1	;  [#uses=1]
	store i64 %3, i64* %x
	ret void
}
Running code...
Code was run.

```
## 8、结论

是不是非常的酷？我同意如果你只是从这篇文章中拷贝粘贴的话，你可能会对运行得到的结果感觉有点失望，但是这点结果可能也会激发你更大的兴趣。此外，这就是本文的意义，这不是本篇指导文章的结束，这只是一个开始。因为有了这篇文章的介绍，你可以在文法分析，语法分析和装配语言的时候附加上一些疯狂的特性，然后创造出一个你自己命名的语言。你现在已经可以编译语句块了，那么你现在应该已经有如何继续下去的基本想法。  

本文完整的代码在Github[这里](https://github.com/lsegal/my_toy_compiler)。我一直都在避免提到这个代码，因为这个代码不是本文的重点，而仅仅是带过这部分代码。

我意识到这是一篇非常长的文章，并且这篇文章中难免会有出错的地方，如果你找到了任何问题，在你觉得有空的时候，欢迎你给我发电子邮件，我将会调整我的文章。你如果向想我们共享一些信息，你也可以在你觉得有空的时候写信给我们。

**（转载本站文章请注明作者和出处 [酷 壳 – CoolShell](https://coolshell.cn/) ，请勿用于任何商业用途）**
