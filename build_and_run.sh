clear
rm Parser.cs.old
rm Scanner.cs.old
rm Compile.exe

coco SyntaxAnalyzer.cs.atg 
csc Compile.cs Scanner.cs Parser.cs
mono Compile.exe input.txt