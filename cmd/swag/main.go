package main

import (
	"encoding/json"
	"fmt"
	"github.com/majintao/go-yapi"
	"github.com/majintao/swag"
	"github.com/majintao/swag/gen"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	searchDirFlag        = "dir"
	generalInfoFlag      = "generalInfo"
	propertyStrategyFlag = "propertyStrategy"
)

var parseFlag = []cli.Flag{
	&cli.StringFlag{
		Name:    generalInfoFlag,
		Aliases: []string{"g"},
		Value:   "main.go",
		Usage:   "需要解析的go文件",
	},
	&cli.StringFlag{
		Name:    searchDirFlag,
		Aliases: []string{"d"},
		Value:   "./",
		Usage:   "需要解析的目录",
	},
	&cli.StringFlag{
		Name:    propertyStrategyFlag,
		Aliases: []string{"p"},
		Value:   "camelcase",
		Usage:   "属性名解析规则：snakecase,camelcase,pascalcase",
	},
}

const (
	testInstanceURL = "http://yapi.iacorn.cn"
	testToken       = "8cde6e3bcbdb1c7bc0d8a2992fe0d79b0f07d186bfe7ebc5c6723b403e6830d8"
)

func listMenuAction(c *cli.Context) error {
	yapiClient, err1 := yapi.NewClient(nil, testInstanceURL, testToken)
	if err1 != nil {
		return fmt.Errorf("Client creation -> Got an error: %s", err1)
	}
	if yapiClient == nil {
		return fmt.Errorf("Expected a client. Got none")
	}

	projectResp, _, _ := yapiClient.Project.Get()
	result, _, _ := yapiClient.CatMenu.Get(projectResp.Data.ID)
	marshal, _ := json.Marshal(result)
	api := string(marshal)
	fmt.Println(api)
	return nil
}

func listApi(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case swag.CamelCase, swag.SnakeCase, swag.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	config := &gen.Config{
		SearchDir:          c.String(searchDirFlag),
		MainAPIFile:        c.String(generalInfoFlag),
		PropNamingStrategy: strategy,
		ParseDependency:    true, // 需要解析外部依赖
	}

	parser, err := gen.New().GetParser(config)
	if err != nil {
		return fmt.Errorf("generate error.")
	}

	for path, _ := range parser.GetSwagger().Paths.Paths {
		fmt.Println(path)
	}

	return nil
}

func UpdateApi(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case swag.CamelCase, swag.SnakeCase, swag.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	config := &gen.Config{
		SearchDir:          c.String(searchDirFlag),
		MainAPIFile:        c.String(generalInfoFlag),
		PropNamingStrategy: strategy,
		ParseDependency:    true, // 需要解析外部依赖
	}

	yapiClient, err1 := yapi.NewClient(nil, testInstanceURL, testToken)
	if err1 != nil {
		return fmt.Errorf("Client creation -> Got an error: %s", err1)
	}
	if yapiClient == nil {
		return fmt.Errorf("Expected a client. Got none")
	}

	swagger, _ := gen.New().BuildSwagger(config)
	swaggerJson := string(swagger)
	result, _, _ := yapiClient.Interface.UploadSwagger(&swaggerJson)
	marshal, _ := json.Marshal(result)
	api := string(marshal)
	fmt.Println(api)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = swag.Version
	app.Usage = "yapi"
	app.Commands = []*cli.Command{
		{
			Name:    "listApi",
			Aliases: []string{"la"},
			Usage:   "解析当前目录文件api",
			Action:  listApi,
			Flags:   parseFlag,
		},
		{
			Name:    "uploadApi",
			Aliases: []string{"ua"},
			Usage:   "解析当前目录文件api",
			Action:  UpdateApi,
			Flags:   parseFlag,
		},
		{
			Name:    "listYApiMenu",
			Aliases: []string{"lym"},
			Usage:   "列举Yapi的目录",
			Action:  listMenuAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
