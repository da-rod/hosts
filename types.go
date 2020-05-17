package main

type source struct {
	Name    string   `json:"name"`
	URLs    []string `json:"urls"`
	Format  string   `json:"format"`
	Exclude []string `json:"exclude"`
}

type sources []source

type whitelist struct {
	Whitelist sources `json:"whitelist"`
}

type blacklist struct {
	Blacklist sources `json:"blacklist"`
}
