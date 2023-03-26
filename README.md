# gopexel
Simple wrapper for Pexels API, all you need to do is create .env file, and define
```.env
PEXELSAPIKEY=<YOUR_API_KEY>
PEXELSAPIURL=https://api.pexels.com/v1/search
```
for more documentation about pexel, visit [pexel documentation site](https://www.pexels.com/api/documentation/)

to install get this library running, import
```go
"github.com/peacefulhack/goPexel/app/models"
```

make a struct from datastruct named PexelAuth
```go
import (
  "github.com/peacefulhack/goPexel/app/shared/datastruct"
)
...
func main() {
  ...
  var PexAuth datastruct.PexelsAuth
  ...
}
```

then i already created function to fetch .env file called GetAllEnv inside utils, or create on your own
to use my function do this:
```go
import (
  "github.com/peacefulhack/goPexel/app/shared/utils"
)
func main() {
  ...
  env, err := utils.GetAllEnv(2)
	if err != nil {
		fmt.Println(err)
		return
	}
  PexAuth.Api = env[0] //env 0 is PEXELSAPIKEY
	PexAuth.Url = env[1] //env 1 is PEXELSAPIURL
  ...
}
```
and then the last is call the function to get the image, for example like this
```go
// the request is the auth, the query of image or what we call keyword, image per-page, width, height, what page does the image will get: 1 for default, image size, and location of image will saved
models.GetImages(PexAuth, "dog", "1", "1920", "1080", "1", "large2x", "temp/")
```

for more example, please see my example folder~
if you have anything to ask ,ask me on discord peacefulhack#6524
