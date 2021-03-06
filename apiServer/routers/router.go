package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/ledger"
	//"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/channel"
)

func init() {
	ns := beego.NewNamespace("/fabric",

		/*
				beego.NSNamespace("/install",
					beego.NSInclude(
						&chaincode.InstallCCController{},
					),
				),
				beego.NSNamespace("/instantiate",
					beego.NSInclude(
						&chaincode.InstantiateController{},
					),
				),
				beego.NSNamespace("/invokeCC",
					beego.NSInclude(
						&chaincode.InvokeController{},
					),
				),
				beego.NSNamespace("/queryCC",
					beego.NSInclude(
						&chaincode.QueryController{},
					),
				),
				beego.NSNamespace("/chaincodeInfo",
					beego.NSInclude(
						&chaincode.ChaincodeInfoController{},
					),
				),

				beego.NSNamespace("/createChannel",
					beego.NSInclude(
						&channel.ChannelCreateController{},
					),
				),
				beego.NSNamespace("/joinChannel",
					beego.NSInclude(
						&channel.ChannelJoinController{},
					),
				),

				beego.NSNamespace("/queryInstalled",
					beego.NSInclude(
						&query.QueryInstalledController{},
					),
				),
				beego.NSNamespace("/queryBlock",
					beego.NSInclude(
						&query.QueryBlockController{},
					),
				),
				beego.NSNamespace("/queryChannels",
					beego.NSInclude(
						&query.QueryChannelsController{},
					),
				),
				beego.NSNamespace("/queryBlockchainInfo",
					beego.NSInclude(
						&query.QueryInfoController{},
					),
				),
			beego.NSNamespace("/user",
				beego.NSInclude(
					&user.UserController{},
				),
			),*/
		beego.NSNamespace("/ledger",
			beego.NSInclude(
				&ledger.LedgerController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&assetApp.UserManageController{},
			),
		),
		beego.NSNamespace("/initFabric",
			beego.NSInclude(
				&assetApp.InitializeController{},
			),
		),
		beego.NSNamespace("/model",
			beego.NSInclude(
				&assetApp.AssetController{},
			),
		),
		beego.NSNamespace("/cert",
			beego.NSInclude(
				&assetApp.CertificateController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
