package main

func main() {
	var testMap = make(map[string]string)
	testMap["key1"] = "value1"
	testMap["key2"] = "value2"
	delete(testMap, "key1")
}
