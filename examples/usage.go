package examples

import (
	"fmt"
	"peacefulhack/goPexel/app/models"
	datastruct2 "peacefulhack/goPexel/app/shared/datastruct"
	"peacefulhack/goPexel/app/shared/utils"
	"strconv"
)

func main() {
	var PexAuth datastruct2.PexelsAuth

	//dont forget to change your .env file on your local project directory
	env, err := utils.GetAllEnv(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	PexAuth.Api = env[0]
	PexAuth.Url = env[1]

	//this is for randomizer image so each request will be different
	ran := utils.Randomizer(100, 1)
	//for size, you will have from tiny to large2x, you can see all size inside https://www.pexels.com/api/documentation/
	models.GetImages(PexAuth, "dog", "1", "1920", "1080", strconv.Itoa(ran), "large2x", "temp/")
}
