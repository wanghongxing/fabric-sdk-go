package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:BlockController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:BlockController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:BlockController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:BlockController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:number`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:LedgerController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:LedgerController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
