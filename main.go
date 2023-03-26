package main

import (
	"fmt"
	"peacefulhack/goPexel/app/models"
	datastruct2 "peacefulhack/goPexel/app/shared/datastruct"
	"peacefulhack/goPexel/app/shared/utils"
	"strconv"
)

func main() {
	var PexAuth datastruct2.PexelsAuth
	env, err := utils.GetAllEnv(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	PexAuth.Api = env[0]
	PexAuth.Url = env[1]
	ran := utils.Randomizer(100, 1)
	models.GetImages(PexAuth, "dog", "1", "1920", "1080", strconv.Itoa(ran), "large2x")
}
