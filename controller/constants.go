package controller

func ReturnContantLanguageid() []string {
	language_id_str := []string{"CPP", "GoLang", "VisualBasic", "php", "CSharp", "Perl", "clojure", "python",
		"GCC", "Shell", "nodeJs", "ruby", "python3"}
	return language_id_str
}

func ReturnFilesList(language_id int) []string {
	data := map[int][]string{0: {"main.cpp", "input.txt", "errors.txt", "output.txt", "compile.sh"},
		1:  {"main.go", "input.txt", "output.txt", "errors.txt", "compile.sh"},
		2:  {"file.vb", "input.txt", "errors.txt", "output.txt", "compile.sh"},
		3:  {"file.php", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		4:  {"Solution.cs", "input.txt", "errors.txt", "output.txt", "compile.sh"},
		5:  {"file.pl", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		6:  {"file.clj", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		7:  {"file.py", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		8:  {"main.c", "input.txt", "errors.txt", "output.txt", "compile.sh"},
		9:  {"script.sh", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		10: {"script.js", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		11: {"file.rb", "input.txt", "output.txt", "compile.sh", "errors.txt"},
		12: {"file.py", "input.txt", "output.txt", "compile.sh", "errors.txt"},
	}

	return data[language_id]
}
