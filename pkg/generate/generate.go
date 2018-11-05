package generate

import (
	"io"
	"os"
	"strings"

	gen "github.com/dave/jennifer/jen"

	"github.com/golang/glog"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/api"
)

var DeprecatedWords = []string{"removed", "deprecated"}

const (
	PackageName string = "sonarqube"
)

func GenerateServiceDoc(apidoc *api.API) error {
	_ = os.Mkdir("test", os.ModeDir)
	for _, item := range apidoc.WebServices {
		name := item.Path[4:]
		contin := false
		for _, word := range DeprecatedWords {
			if strings.Contains(strings.ToLower(item.Description), word) {
				glog.V(0).Infof("Detected deprecated api:%s,source:%s\n", item.Path, item.Description)
				contin = true
				break
			}
		}
		if contin {
			continue
		}
		newFile, err := os.Create("test/" + strings.Title(name) + "Service.go")
		glog.V(2).Infof("Creaing go file %s", newFile.Name())
		if err != nil {
			return err
		}
		defer newFile.Close()
		if err := AddGoContent(&item, newFile); err != nil {
			return err
		}
	}
	return nil
}

func AddStaticFile() error {
	return nil
}
func AddGoContent(service *api.WebService, w io.Writer) error {
	f := gen.NewFile(PackageName)
	f.CanonicalPath = "github.com/magicsong/generate-go-for-sonarqube/pkg/" + PackageName
	f.PackageComment(service.Description)
	f.ImportName("github.com/google/glog", "glog")

	name := service.Path[4:]
	//Create Service Struct
	f.Type().Id(strings.Title(name) + "Service").Struct(
		gen.Id("client").Op("*").Id("Client"),
	)

	//Create Methods
	for _, item := range service.Actions {
		AddMethodOfAction(name, &item, f)
	}
	return f.Render(w)
}

func AddMethodOfAction(serviceName string, action *api.Action, f *gen.File) {
	//Create method options
	optionName := strings.Title(serviceName) + strings.Title(action.Key) + "Option"
	f.Type().Id(optionName).StructFunc(func(g *gen.Group) {
		for _, field := range action.Params {
			if strings.Contains(strings.ToLower(field.Description), "deprecated") {
				continue
			}
			g.Id(strings.Title(field.Key)).String().Commentf("Description:\"%s\",ExampleValue:\"%s\"", field.Description, field.ExampleValue)
		}
	})
	//create method
	f.Func().Params(gen.Id("s").Op("*").Id(strings.Title(serviceName)+"Service")).Id(strings.Title(action.Key)).Params(
		gen.Id("opt").Op("*").Id(optionName)).Params(
		gen.Id("Resp").Op("*").Interface(), gen.Err().Error()).Block(
		gen.Return(),
	)
}
