package main

type source struct {
	Name    string   `json:"name"`
	URLs    []string `json:"urls"`
	Format  string   `json:"format"`
	Exclude []string `json:"exclude"`
}

type sources []source

type safelist struct {
	List sources `json:"safelist"`
}

type blocklist struct {
	List sources `json:"blocklist"`
}
