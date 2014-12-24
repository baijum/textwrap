package textwrap

func Wrap(text string, width int) []string {
	var out []string
	old := 0
	for pos, _ := range text {
		if pos%width == 0 {
			if pos == 0 {
				continue
			}
			out = append(out, text[old:pos])
			old = pos
		}
	}
	out = append(out, text[old:len(text)])
	return out
}
