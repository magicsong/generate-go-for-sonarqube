package generate

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	gen "github.com/dave/jennifer/jen"
	"github.com/golang/glog"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/api"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/util/strcase"
)

var (
	DeprecatedWords = []string{"removed", "deprecated"}
	validtion       *gen.File
)

const (
	PackageName string = "sonarqube"
	WorkingDir         = "test"
)

func init() {

}
func Build(apidoc *api.API) error {
	_ = os.Mkdir(WorkingDir, os.ModeDir)
	err := AddStaticFile()
	if err != nil {
		glog.Errorln("Import static files failed")
		return err
	}
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
		newFile, err := os.Create(WorkingDir + "/" + name + "_service.go")
		glog.V(3).Infof("Creaing go file %s", newFile.Name())
		if err != nil {
			return err
		}
		defer newFile.Close()
		if err := AddGoContent(&item, newFile); err != nil {
			return err
		}
	}
	//Write validation file
	validtion.Save(WorkingDir + "/validation.go")
	return nil
}

func AddStaticFile() error {
	err := ioutil.WriteFile(WorkingDir+"/sonarqube.go", []byte(SonarqubeConst), 0644)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(WorkingDir+"/web_client.go", []byte(WebClientConst), 0644)
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
	optionName := strcase.ToCamel(serviceName + action.Key + "Option")
	f.Type().Id(optionName).StructFunc(func(g *gen.Group) {
		for _, field := range action.Params {
			if strings.Contains(strings.ToLower(field.Description), "deprecated") {
				glog.V(0).Infof("Detected deprecated field <%s> in <action>:%s,description:%s\n", field.Key, action.Key, field.Description)
				continue
			}
			g.Id(strcase.ToCamel(field.Key)).String().Commentf("Description:\"%s\",ExampleValue:\"%s\"", field.Description, field.ExampleValue)
		}
	})
	//create valid method
	validtion.Func().Params(gen.Id("s").Op("*").Id(strings.Title(serviceName) + "Service")).Id("Validate" + strcase.ToCamel(action.Key) + "Opt").Params(
		gen.Id("opt").Op("*").Id(optionName)).Params(gen.Error()).Block(
		gen.Return(gen.Nil()),
	)
	//create method
	method := "GET"
	if action.Post {
		method = "POST"
	}
	f.Func().Params(gen.Id("s").Op("*").Id(strings.Title(serviceName)+"Service")).Id(strings.Title(action.Key)).Params(
		gen.Id("opt").Op("*").Id(optionName)).Params(
		gen.Id("Resp").Op("*").Interface(), gen.Err().Error()).BlockFunc(func(g *gen.Group) {
		g.List(gen.Id("req"), gen.Id("err")).Op(":=").Id("s").Dot("client").Dot("NewRequest").Call(gen.Lit(method), gen.Lit(serviceName+"/"+action.Key), gen.Id("Opt"))
		g.If(
			gen.Err().Op("!=").Nil(),
		).Block(
			gen.Return(gen.Err(), gen.Nil()),
		)
		g.Err().Op("=").Id("s").Dot("client").Dot("Do").Call(gen.Id("req"), gen.Id("Resp"))
		g.If(
			gen.Err().Op("!=").Nil(),
		).Block(
			gen.Return(),
		)
		g.Return()
	})
}
