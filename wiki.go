package main

import "os"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() (err error) {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}
