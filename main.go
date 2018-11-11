package main

func main() {
	changelogItems, err := BuildChangeLog()
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(changelogItems); i++ {
			item := changelogItems[i]
			println(item.Card.Name)
			println(item.Date.String())
			println(item.Card.Desc)
			println()
		}
	}
}
