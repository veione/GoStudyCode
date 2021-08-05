package main

func main()  {
	trie :=  Constructor()
	trie.Insert("apple")
	trie.Search("apple")   // 返回 True
	trie.Search("app")     // 返回 False
	trie.StartsWith("app") // 返回 True
	trie.Insert("app")
	trie.Search("app")    // 返回 True


}
