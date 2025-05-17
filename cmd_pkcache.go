package main

func commandPkcache(cfg *config, params ...string) error {
	cfg.pokeapiClient.PrintCache()
	return nil
}
