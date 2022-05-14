package hypixel

import "os"

func Key() string {
	return os.Getenv("HYPIXEL_API_KEY")
}
