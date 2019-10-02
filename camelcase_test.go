package camelcase

import "fmt"

func ExampleSplit() {

	for _, c := range []string{
		"",
		"lowercase",
		"Class",
		"MyClass",
		"MyC",
		"HTML",
		"PDFLoader",
		"AString",
		"SimpleXMLParser",
		"vimRPCPlugin",
		"GL11Version",
		"99Bottles",
		"May5",
		"BFG9000",
		"BöseÜberraschung",
		"Two  spaces",
		"BadUTF8\xe2\xe2\xa1",
	} {
		fmt.Printf("%#v => %#v\n", c, Split(c))
	}

	// Output:
	// "" => []string{}
	// "lowercase" => []string{"lowercase"}
	// "Class" => []string{"Class"}
	// "MyClass" => []string{"My", "Class"}
	// "MyC" => []string{"My", "C"}
	// "HTML" => []string{"HTML"}
	// "PDFLoader" => []string{"PDF", "Loader"}
	// "AString" => []string{"A", "String"}
	// "SimpleXMLParser" => []string{"Simple", "XML", "Parser"}
	// "vimRPCPlugin" => []string{"vim", "RPC", "Plugin"}
	// "GL11Version" => []string{"GL", "11", "Version"}
	// "99Bottles" => []string{"99", "Bottles"}
	// "May5" => []string{"May", "5"}
	// "BFG9000" => []string{"BFG", "9000"}
	// "BöseÜberraschung" => []string{"Böse", "Überraschung"}
	// "Two  spaces" => []string{"Two", "  ", "spaces"}
	// "BadUTF8\xe2\xe2\xa1" => []string{"BadUTF8\xe2\xe2\xa1"}
}

func ExampleToMacroCase() {
	for _, s := range []string{
		"",
		"lowercase",
		"Class",
		"MyClass",
		"MyC",
		"HTML",
		"PDFLoader",
		"AString",
		"SimpleXMLParser",
		"vimRPCPlugin",
		"GL11Version",
		"99Bottles",
		"May5",
		"BFG9000",
		"BöseÜberraschung",
		"Two  spaces",
	} {
		fmt.Printf("%#v => %#v\n", s, ToMacroCase(s))
	}

	// OUTPUT:
	// "" => ""
	// "lowercase" => "LOWERCASE"
	// "Class" => "CLASS"
	// "MyClass" => "MY_CLASS"
	// "MyC" => "MY_C"
	// "HTML" => "HTML"
	// "PDFLoader" => "PDF_LOADER"
	// "AString" => "A_STRING"
	// "SimpleXMLParser" => "SIMPLE_XML_PARSER"
	// "vimRPCPlugin" => "VIM_RPC_PLUGIN"
	// "GL11Version" => "GL_11_VERSION"
	// "99Bottles" => "99_BOTTLES"
	// "May5" => "MAY_5"
	// "BFG9000" => "BFG_9000"
	// "BöseÜberraschung" => "BÖSE_ÜBERRASCHUNG"
	// "Two  spaces" => "TWO_SPACES"
}
