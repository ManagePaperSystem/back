package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["APIProject/controllers:QuesController"] = append(beego.GlobalControllerRouter["APIProject/controllers:QuesController"],
        beego.ControllerComments{
            Method: "Check",
            Router: "/check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:QuesController"] = append(beego.GlobalControllerRouter["APIProject/controllers:QuesController"],
        beego.ControllerComments{
            Method: "CheckList",
            Router: "/check/list",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:QuesController"] = append(beego.GlobalControllerRouter["APIProject/controllers:QuesController"],
        beego.ControllerComments{
            Method: "CheckPaper",
            Router: "/check/paper",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:QuesController"] = append(beego.GlobalControllerRouter["APIProject/controllers:QuesController"],
        beego.ControllerComments{
            Method: "DeletePaper",
            Router: "/delete",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:QuesController"] = append(beego.GlobalControllerRouter["APIProject/controllers:QuesController"],
        beego.ControllerComments{
            Method: "Generate",
            Router: "/gen",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:UserController"] = append(beego.GlobalControllerRouter["APIProject/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetCode",
            Router: "/getcode",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:UserController"] = append(beego.GlobalControllerRouter["APIProject/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:UserController"] = append(beego.GlobalControllerRouter["APIProject/controllers:UserController"],
        beego.ControllerComments{
            Method: "ModifyPassword",
            Router: "/modify",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:UserController"] = append(beego.GlobalControllerRouter["APIProject/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: "/reg",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APIProject/controllers:UserController"] = append(beego.GlobalControllerRouter["APIProject/controllers:UserController"],
        beego.ControllerComments{
            Method: "SetAccount",
            Router: "/set",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
