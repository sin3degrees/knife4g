# Knife4g
use knife4j-front to show the api documents for iris

# Usage
1. Add comments to your API source code
2. Download [Swag](https://github.com/swaggo/swag) for Go by using:
3. Run `swag init -ot json` in your project directory
4. Run `go get github.com/sin3degrees/knife4g`
5. Add router to your iris project
    ### example:
    ```go
    package main
    
    import (
        "github.com/kataras/iris/v12"
        "github.com/sin3degrees/knife4g"
    )
    
    func main() {
        engine := iris.Default()
        engine.Get("/doc/*any", knife4g.Handler(knife4g.Config{RelativePath: "/doc"}))
        engine.Run(":80")
    }
    ```
6. Visit http://localhost/doc/index

# Acknowledgement
Thanks to [knife4j](https://github.com/xiaoymin/swagger-bootstrap-ui)
Thanks to [knife4g](https://github.com/hononet639/knife4g)