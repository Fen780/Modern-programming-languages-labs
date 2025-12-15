import re

class Translator:
    rules = [
        (r"{", ""),
        (r"}", ""),
        (r";", ""),
        (r"public static void main\(String\[\] args\)", r"def main():"),
        
        (r"public static (boolean|void) (\w+)\s*\((.*?)\)", r"def \2(\3):"),
        (r"(int|String\[\])\s+", ""),
        
        (r"(int|boolean)\s+(\w+)\s*=\s*(.*)", r"\2 = \3"),
        
        (r"while\s*\((.*?)\)", r"while \1:"),
        
        (r"if\s*\((.*?)\)", r"if \1:"),
        (r"else", r"else:"),
        
        (r"System\.out\.print\((.*?)\)", r"print(\1)"),
        
        (r"(\w+)\s*\+\+", r"\1 += 1"),        
        (r"\&\&", r"and"),
    ]

    def translate_line(self, line):
        stripped_line = line.strip()
        
        for pattern, replacement in self.rules:
            stripped_line = re.sub(pattern, replacement, stripped_line)

        if stripped_line.startswith("print("):
            match = re.search(r"print\((.*?)\)", stripped_line)
            if match:
                content = match.group(1)
                new_content = content.replace('+', ', ')
                stripped_line = f"print({new_content})"

        return stripped_line.strip()

    def translate(self, java_code):
        py_lines = []
        tabs = 0
        
        for l in java_code.splitlines():
            line = l.strip()
            
            if not line:
                continue
            
            tabs -= line.count("}")
            tabs = max(0, tabs)
            
            translated = self.translate_line(line)
            
            if translated:
                py_lines.append("    " * tabs + translated)
            
            tabs += line.count("{")
        
        py_lines.append("\nif __name__ == '__main__':")
        py_lines.append(f"    main()")
        
        return "\n".join(py_lines)
    
    
java_code = """
public static boolean isDivisibleBy5(int num)
{
    return num >= 0 && num % 5 == 0;
}

public static void main(String[] args)
{
    int i = 0;
    int count = 0;
    while (i < 10)
    {
        if (i % 2 == 0)
        {
            System.out.print("Четное число");
            count++;
        }
        else
        {
            System.out.print("Нечетное число");
        }
        if (i < 0)
        {
            System.out.print("Отрицательное");
        }
        else
        {
            System.out.print("Неотрицательное");
        }
        i++;
    }
    System.out.print("Проверили:" + count + "чисел");

    boolean result = isDivisibleBy5(7);
    System.out.print("7 делится на 5? -" + result);
}
"""

translator = Translator()
py_code = translator.translate(java_code)
print(py_code)
print("\n\nРезультат работы программы:\n")
exec(py_code)  