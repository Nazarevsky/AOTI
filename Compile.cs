using System;

class Compile {
    static void Main (string[] arg){
        Scanner scanner = new Scanner(arg[0]);
        
        Parser parser = new Parser (scanner);
        parser.Parse();

        if (parser.errors.count == 0) {
            Console.WriteLine("No errors!");
        } else {
            Console.WriteLine(parser.errors.count + " errors detected");
        }
    }
}