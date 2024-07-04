package loader

import (
	"encoding/json"
	"fmt"
	"os"
)

func SeedLoad(seed_path string) ([]string, error) {
	var seeds []string
	data, err := os.ReadFile(seed_path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &seeds)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal(): %s", err.Error())
	}
	if len(seeds) == 0 {
		return nil, fmt.Errorf("no seed in %s", seed_path)
	}
	return seeds, nil
}
