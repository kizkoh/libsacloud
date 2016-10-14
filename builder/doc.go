// Package builder is the High Level APIs for creating resources on SakuraCloud.
//
// さくらのクラウドでのリソース作成用の高レベルAPIです。
// サーバー/ディスク作成時の手順を単純化します。
//
// Building resources
//
// リソースの作成は以下のように行います。
//	import (
//		"github.com/yamamoto-febc/libsacloud/api"
//		"github.com/yamamoto-febc/libsacloud/builder"
//		"github.com/yamamoto-febc/libsacloud/sacloud/ostype"
//	)
//
//	func main() {
//
//		// APIクライアントの作成
//		client := api.NewClient("PUT-YOUR-TOKEN", "PUT-YOUR-SECRET", "tk1a")
//
//		// パブリックアーカイブ(CentOS)から作成するビルダー、共有セグメントに接続、以外はデフォルト値で作成
//		res , err := builder.FromPublicArchiveUnix(client, ostype.CentOS, "ServerName", "Password").AddPublicNWConnectedNIC().Build()
//
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("%v" , res.Server)
//	}
// 1) 作成したいサーバーのディスク/ソースアーカイブの種類ごとにビルダーを作成します。
//
// 2) 必要に応じてNICやディスク、サーバースペックなどをビルダーのメソッドで設定します。
//
// 3) Buildメソッドを呼び出すことでサーバーが作成されます。
//
//
// Server builder types
//
// ビルダーはディスク構成やソースアーカイブ/ディスクにより以下のような種類に分かれています。
//
// Linux(Unix)系パブリックアーカイブ
//	func FromPublicArchiveUnix(client *api.Client, os ostype.ArchiveOSTypes, name string, password string) *PublicArchiveUnixServerBuilder
//
// Windows系パブリックアーカイブ
//	func FromPublicArchiveWindows(client *api.Client, name string, archiveID int64) *PublicArchiveWindowsServerBuilder
//
// 汎用(アーカイブから)
//	func FromArchive(client *api.Client, name string, sourceArchiveID int64) *CommonServerBuilder
//
// 汎用(ディスクから)
//	func FromDisk(client *api.Client, name string, sourceDiskID int64) *CommonServerBuilder
//
// ディスクレス
//	func FromDiskless(client *api.Client, name string) *DisklessServerBuilder
//
// 空のディスク
//	func FromBlankDisk(client *api.Client, name string) *BlankDiskServerBuilder
//
//
//
// Fluent API
//
// ビルダーではFluent APIを採用しています。
//
// 以下の例は同じ内容の処理を通常の書き方/Fluent APIでの書き方両方で記載しています。
//
//	// 通常
//	b := builder.FromPublicArchiveUnix(client, ostype.CentOS, "ServerName", "Password")
//	b.AddPublicNWConnectedNIC()
//	res , err := b.Build()
//
//	// Fluent APIの場合
//	res , err := builder.FromPublicArchiveUnix(client, ostype.CentOS, "ServerName", "Password").AddPublicNWConnectedNIC().Build()
//
//
// Event handling
//
// ビルダーでは各タイミングごとにイベントハンドラ(コールバック)が利用可能です。
//	func main() {
//		// APIクライアントの作成
//		client := api.NewClient("PUT-YOUR-TOKEN", "PUT-YOUR-SECRET", "tk1a")
//
//		// ディスクレスビルダー、イベントハンドラ(ServerBuildOnComplete)を登録
//		builder.FromDiskless(client, "example").
//			SetEventHandler(builder.ServerBuildOnComplete, callbackFunc).
//			Build()
//	}
//
//	func callbackFunc(value *builder.ServerBuildValue, result *builder.ServerBuildResult) {
//		// Do someting here
//	}
//
package builder
