package hole

func emojify() ([]string, string) {
	return outputTests(shuffle([]test{
		{":-D", "😀"},
		{":-)", "🙂"},
		{":-|", "😐"},
		{":-(", "🙁"},
		{`:-\`, "😕"},
		{":-*", "😗"},
		{":-O", "😮"},
		{":-#", "🤐"},
		{"':-D", "😅"},
		{"':-(", "😓"},
		{":'-)", "😂"},
		{":'-(", "😢"},
		{":-P", "😛"},
		{";-P", "😜"},
		{"X-P", "😝"},
		{"X-)", "😆"},
		{"O:-)", "😇"},
		{";-)", "😉"},
		{":-$", "😳"},
		{":-", "😶"},
		{"B-)", "😎"},
		{":-J", "😏"},
		{"}:-)", "😈"},
		{"}:-(", "👿"},
		{":-@", "😡"},
	}))
}
