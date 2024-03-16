using System;

class Compile {
    static void Main (string[] arg){
        Dictionary<int, string> tokenNames = new Dictionary<int, string>() {
            {1, "Operation"},
            {2, "Comparison"},
            {3, "VarName"},
            {4, "IntNum"},
            {5, "FloatNum"},
            {6, "Semicol"},
            {7, "Equal"},
            {8, "IfStatement"},
            {9, "ParOpen"},
            {10, "ParClose"},
            {11, "CurlyOpen"},
            {12, "CurlyClose"},
            {13, "ElseStatement"},
            {14, "FloatType"},
            {15, "IntType"},
            {16, "UnknownSymbol"},
        };
        Token t;
        if (arg.Length > 0) {
            Scanner scanner = new Scanner(arg[0]);
            while ((t = scanner.Scan()).kind != 0 ) {
                //Console.WriteLine("TokenId: {0}, TokenName: {1}, Lexeme: {2}", t.kind, tokenNames[t.kind], t.val);
                Console.WriteLine(tokenNames[t.kind]);
            }
        } 
        else {
            Console.WriteLine("No source file specified");
        }
    }
}