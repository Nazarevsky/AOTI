using System;

class Compile {
    static void Main (string[] arg){
        Token t;
        if (arg.Length > 0) {
            Scanner scanner = new Scanner(arg[0]);
            while ((t = scanner.Scan()).kind != 0 ) {
                Console.WriteLine("Token: {0}, Lexeme: {1}", t.kind, t.val);
            }
        } 
        else {
            Console.WriteLine("No source file specified");
        }

        // Scanner scanner = new Scanner(arg[0]);
        // Parser parser = new Parser (scanner);
        // parser.Parse();

        // if (parser.errors.count == 0) {
        //     Console.WriteLine("No errors!");
        // } else {
        //     Console.WriteLine(parser.errors.count + " errors detected");
        // }
    }
}