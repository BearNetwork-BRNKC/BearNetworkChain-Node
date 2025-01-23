## BearNetworkChain-Node

以太坊協定的Golang執行層實作。

[![API Reference](
https://pkg.go.dev/badge/github.com/ethereum/go-ethereum
)](https://pkg.go.dev/github.com/ethereum/go-ethereum?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethereum/go-ethereum)](https://goreportcard.com/report/github.com/ethereum/go-ethereum)
[![Travis](https://app.travis-ci.com/ethereum/go-ethereum.svg?branch=master)](https://app.travis-ci.com/github/ethereum/go-ethereum)
[![Discord](https://img.shields.io/badge/discord-join%20chat-blue.svg)](https://discord.gg/nthXNEv)

自動化建置可用於穩定版本和不穩定的主分支。二進位
檔案發佈於 https://geth.ethereum.org/downloads/.

## 建構原始碼

有關先決條件和詳細構建說明，請閱讀[安裝說明](https://geth.ethereum.org/docs/getting-started/installing-geth).

建置 `geth` 需要 Go（版本 1.19 或更高版本）和 C 編譯器。 您可以安裝
他們使用您最喜歡的套件管理器。安裝依賴項後，執行

```shell
make geth
```

或者，建立全套實用程式：

```shell
make all
```

## 執行檔

go-ethereum 專案附帶了在“cmd”中找到的幾個包裝器/可執行文件
目錄。

|  Command   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| :--------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`geth`** | 
我們主要的熊網鏈 CLI 客戶端。它是以太坊網路（主網、測試網或專用網路）的入口點，能夠作為完整節點（預設）、存檔節點（保留所有歷史狀態）或輕節點（即時檢索資料）運作。它可以被其他進程用作透過 HTTP、WebSocket 和/或 IPC 傳輸之上公開的 JSON RPC 端點進入熊網鏈網路的網關。 `geth --help` 和命令列選項的 [CLI 頁面](https://geth.ethereum.org/docs/fundamentals/command-line-options)。 |
|   `clef`   | 獨立簽名工具，可作為後端簽署者`geth`.                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
|  `devp2p`  | 與網路層上的節點互動的實用程序，無需運行完整的區塊鏈。                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|  `abigen`  | 
原始碼產生器，用於將以太坊合約定義轉換為易於使用、編譯時類型安全的 Go 套件。它在普通的[熊網鏈合約 ABI](https://docs.soliditylang.org/en/develop/abi-spec.html) 上運行，如果合約字節碼也可用，則具有擴展功能。然而，它也接受 Solidity 原始文件，使開發更加簡化。請參閱我們的 [Native DApps](https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings) 頁面以了解詳細資訊。                                  |
| `bootnode` | 我們的熊網鏈客戶端實現的精簡版本，僅參與網路節點發現協議，但不運行任何更高級別的應用程式協議。它可以用作輕量級引導節點，以幫助在專用網路中找到對等點。                                                                                                                                                                                                                                               |
|   `evm`    | EVM（以太坊虛擬機器）的開發實用程式版本，能夠在可設定的環境和執行模式中運行字節碼片段。其目的是允許對 EVM 操作碼進行隔離、細粒度的調試（例如. `evm --code 60ff60ff --debug run`).                                                                                                                                                                                                                                             |
| `rlpdump`  | 用於轉換二進位RLP（[遞歸長度前綴](https://ethereum.org/en/developers/docs/data-structs-and-encoding/rlp)）轉儲（以太坊協議和網路使用的資料編碼）的開發人員實用工具以及共識）到使用者友善的層次表示（例如 `rlpdump --hex CE0183FFFFFFC4C304050583616263`).                                                                                                                                                                               |

## Running `geth`

遍歷所有可能的命令列標誌超出了此處的範圍（請參閱我們的
[CLI 維基頁面](https://geth.ethereum.org/docs/fundamentals/command-line-options)),
但我們列舉了一些常見的參數組合，以幫助您快速上手
關於如何運行自己的“geth”實例。

### Hardware Requirements

最低限度：

* CPU with 2+ cores
* 4GB RAM
* 1TB free storage space to sync the Mainnet
* 8 MBit/sec download Internet service

建議使用的：

* Fast CPU with 4+ cores
* 16GB+ RAM
* High-performance SSD with at least 1TB of free space
* 25+ MBit/sec download Internet service

### Full node on the main Ethereum network

到目前為止，最常見的場景是人們想要簡單地與以太坊交互
網路：建立帳戶；轉移資金；部署合約並與之互動。為了這
特定的用例，用戶不關心多年的歷史數據，所以我們可以
快速同步到網路的目前狀態。為此：

```shell
$ geth console
```

該命令將：
 *以快照同步模式啟動`geth`（默認，可以使用`--syncmode`標誌更改），
   使其下載更多資料以換取避免處理整個歷史記錄
   以太坊網路的 CPU 密集型。
 *啟動內建的互動式[JavaScript控制台](https://geth.ethereum.org/docs/interacting-with-geth/javascript-console),
（透過尾隨的 `console` 子指令）您可以使用 [`web3` 方法](https://github.com/ChainSafe/web3.js/blob/0.20.7/DOCUMENTATION.md) 進行交互 
   （注意：「geth」中捆綁的「web3」版本非常舊，並且與官方文件不同步），
   以及`geth`自己的[管理API](https://geth.ethereum.org/docs/interacting-with-geth/rpc)。
   該工具是可選的，如果您省略它，您可以隨時將其附加到已經運行的
   `geth` instance with `geth attach`.

### A Full node on the Görli test network

如果您想嘗試創建熊網鏈，請轉向開發人員合同，您幾乎肯定希望在不涉及任何真實資金的情況下做到這一點，直到
您將掌握整個系統的竅門。換句話說，不是附加到主網絡，你想用你的節點加入**測試**網絡，這完全相當於
主網絡，但僅包含 play-Ether。

```shell
$ geth --goerli console
```

`console` 子命令與上面的意思相同，等同於在測試網上也很有用。

然而，指定 `--goerli` 標誌會稍微重新配置您的 `geth` 實例：

 * 客戶端將連接到 Görli，而不是連接到主以太坊網絡測試網絡，使用不同的P2P啟動節點、不同的網路ID和創世地址。
 * 不使用預設資料目錄（例如 Linux 上的“~/.ethereum”），而是使用“geth”會將自己嵌套到“goerli”子資料夾中更深一層（“~/.ethereum/goerli”
   Linux）。請注意，在 OSX 和 Linux 上，這也意味著連接到正在運行的測試網節點需要使用自訂端點，因為 `geth Attach` 將嘗試附加到
   預設生產節點端點，例如`geth Attach <datadir>/goerli/geth.ipc`。 Windows使用者不受這個指令。

*註：雖然一些內部保護措施阻止交易
主網和測試網之間的交叉，你應該始終
使用單獨的帳戶進行遊戲和真錢。除非你手動移動
帳戶，「geth」預設會正確分離兩個網絡，並且不會產生任何
他們之間可以使用帳戶。

### Configuration

作為將眾多標誌傳遞給“geth”二進位檔案的替代方法，您還可以傳遞
設定檔通過：

```shell
$ geth --config /path/to/your_config.toml
```

若要了解檔案的外觀，您可以使用「dumpconfig」子命令
匯出您現有的配置：

```shell
$ geth --your-favourite-flags dumpconfig
```

*注意：這僅適用於 `geth` v1.6.0 及更高版本。

#### Docker quick start

在您的電腦上啟動並運行以太坊的最快方法之一是使用
Docker:

```shell
docker run -d --name ethereum-node -v /Users/alice/ethereum:/root \
           -p 8545:8545 -p 30303:30303 \
           ethereum/client-go
```

這將以快照同步模式啟動“geth”，資料庫記憶體限額為 1GB，因為
上面的命令確實如此。  它還會在您的主目錄中建立一個持久性卷
保存您的區塊鏈並映射預設連接埠。還有一個「alpine」標籤
可用於影像的瘦身版本。

如果你想從其他容器存取 RPC，不要忘記 `--http.addr 0.0.0.0`
和/或主機。預設情況下，`geth` 綁定到本地接口，RPC 端點不綁定
從外部可存取。

### Programmatically interfacing `geth` nodes

As a developer, sooner rather than later you'll want to start interacting with `geth` and the
Ethereum network via your own programs and not manually through the console. To aid
this, `geth` has built-in support for a JSON-RPC based APIs ([standard APIs](https://ethereum.github.io/execution-apis/api-documentation/)
and [`geth` specific APIs](https://geth.ethereum.org/docs/interacting-with-geth/rpc)).
These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based
platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by `geth`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` 啟用 HTTP-RPC 伺服器
  * `--http.addr` HTTP-RPC伺服器監聽介面 (default: `localhost`)
  * `--http.port` HTTP-RPC伺服器監聽連接埠 (default: `8545`)
  * `--http.api` 透過 HTTP-RPC 介面提供的 API(default: `eth,net,web3`)
  * `--http.corsdomain` 逗號分隔的接受跨來源請求的網域列表 (browser enforced)
  * `--ws` 啟用 WS-RPC 伺服器
  * `--ws.addr` WS-RPC伺服器監聽介面 (default: `localhost`)
  * `--ws.port` WS-RPC伺服器監聽端口(default: `8546`)
  * `--ws.api` 透過 WS-RPC 介面提供的 API(default: `eth,net,web3`)
  * `--ws.origins` 接受 WebSocket 請求的來源
  * `--ipcdisable` 停用 IPC-RPC 伺服器
  * `--ipcapi` 透過 IPC-RPC 介面提供的 API(default: `admin,debug,eth,miner,net,personal,txpool,web3`)
  * `--ipcpath` 資料目錄中 IPC 套接字/管道的檔案名(explicit paths escape it)

您需要使用自己的程式設計環境的功能（庫、工具等）來
透過 HTTP、WS 或 IPC 連線到配置上述標誌的 `geth` 節點，您將
需要在所有傳輸上使用 [JSON-RPC](https://www.jsonrpc.org/specation)。你
可以為多個請求重複使用同一個連線！

**注意：請了解開放基於 HTTP/WS 的安全隱患
運送之前這樣做！網路上的駭客正積極嘗試顛覆
具有公開 API 的以太坊節點！此外，所有瀏覽器選項卡都可以本地訪問
運行網頁伺服器，因此惡意網頁可能會嘗試破壞本地可用
蜜蜂！

### Operating a private network

維護您自己的專用網路更加複雜，因為需要進行大量配置
官方網路授予的權限需要手動設定。

#### Defining the private genesis state

首先，您需要建立網路的創世狀態，所有節點都需要
知曉並同意。這由一個小的 JSON 檔案組成（例如，將其稱為“genesis.json”）：

```json
{
  "config": {
    "chainId": <arbitrary positive integer>,
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0,
    "byzantiumBlock": 0,
    "constantinopleBlock": 0,
    "petersburgBlock": 0,
    "istanbulBlock": 0,
    "berlinBlock": 0,
    "londonBlock": 0
  },
  "alloc": {},
  "coinbase": "0x0000000000000000000000000000000000000000",
  "difficulty": "0x20000",
  "extraData": "",
  "gasLimit": "0x2fefd8",
  "nonce": "0x0000000000000042",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```

儘管我們建議更改，但上述欄位應該適合大多數用途
將「nonce」設定為某個隨機值，以便防止未知的遠端節點能夠
與您聯繫。如果您想為一些帳戶預先提供資金以方便測試，請建立
帳戶並用其地址填充“alloc”字段。

```json
"alloc": {
  "0x0000000000000000000000000000000000000001": {
    "balance": "111111111"
  },
  "0x0000000000000000000000000000000000000002": {
    "balance": "222222222"
  }
}
```

使用上面 JSON 檔案中定義的創世狀態，您需要初始化 **every**
在啟動之前使用“geth”節點以確保所有區塊鏈參數正確
放：

```shell
$ geth init path/to/genesis.json
```

#### Creating the rendezvous point

將要運行的所有節點初始化為所需的創世狀態後，您需要
啟動一個引導節點，其他人可以使用該節點在您的網路和/或網路上找到彼此
網際網路.乾淨的方法是配置並運行專用的引導節點：

```shell
$ bootnode --genkey=boot.key
$ bootnode --nodekey=boot.key
```

當 bootnode 線上時，它將顯示一個 [`enode` URL](https://ethereum.org/en/developers/docs/networking-layer/network-addresses/#enode)
其他節點可以用來連接到它並交換對等資訊。確保
將顯示的 IP 位址資訊（很可能是 `[::]`）替換為您的外部 IP 位址資訊
可存取的 IP 來取得實際的「enode」 URL。

*注意：您也可以使用成熟的“geth”節點作為引導節點，但它的作用較小
推薦方式。

#### Starting up your member nodes

隨著引導節點的運行和外部可存取（您可以嘗試
`telnet <ip> <port>` 以確保它確實可以存取），啟動每個後續的 `geth`
節點透過「--bootnodes」標誌指向引導節點以進行對等發現。它將
可能還需要將專用網路的資料目錄分開，所以
也請指定自訂“--datadir”標誌。

```shell
$ geth --datadir=path/to/custom/data/folder --bootnodes=<bootnode-enode-url-from-above>
```

*注意：由於您的網路將與主網路和測試網路完全切斷，您將
還需要配置一個礦工來處理交易並為您建立新區塊。

#### Running a private miner


在專用網路設定中，單一 CPU 礦工實例足以滿足
實用目的，因為它可以以正確的間隔產生穩定的塊流
不需要大量資源（考慮在單一執行緒上運行，不需要多個
也有）。要啟動“geth”實例進行挖掘，請使用所有常用標誌運行它，擴展
經過：

```shell
$ geth <usual-flags> --mine --miner.threads=1 --miner.etherbase=0x0000000000000000000000000000000000000000
```

這將在單一 CPU 執行緒上開始挖掘區塊和交易，將所有
處理到「--miner.etherbase」指定的帳戶。您可以進一步調整挖礦
透過更改預設氣體限制塊收斂到（`--miner.targetgaslimit`）和價格
交易在（`--miner.gasprice`）接受。

## Contribution

感謝您考慮幫助提供原始碼！我們歡迎貢獻
來自互聯網上的任何人，並感謝即使是最小的修復！

如果您想為 go-ethereum 做出貢獻，請分叉、修復、提交並發送拉取請求
供維護人員審查並合併到主程式碼庫中。如果您想提交
不過，如果有更複雜的更改，請先在 [我們的 Discord 伺服器](https://discord.gg/invite/nthXNEv) 上與核心開發人員核實
確保這些變更符合項目的整體理念和/或得到
一些早期回饋可以讓您的工作以及我們的審核變得更加輕鬆
合併程序快速而簡單。

請確保您的貢獻符合我們的編碼指南：

 *程式碼必須遵循Go官方[格式](https://golang.org/doc/effective_go.html#formatting)
   指南（即使用 [gofmt](https://golang.org/cmd/gofmt/)）。
 *程式碼必須依照官方 Go [註釋] 進行記錄(https://golang.org/doc/effective_go.html#commentary)
   指導方針。
 *拉取請求需要基於“master”分支並針對“master”分支打開。
*提交訊息應該以它們修改的包為前綴。
   *例如“eth，rpc：使追蹤配置可選”

請參閱[開發者指南](https://geth.ethereum.org/docs/developers/geth-developer/dev-guide)
有關配置環境、管理專案依賴項的更多詳細信息，以及
測試程序。

### 為 geth.ethereum.org 做出貢獻
如需對 [go-ethereum 網站](https://geth.ethereum.org) 做出貢獻，請查看「website」分支並提出拉取請求。
有關更詳細的說明，請參閱 `website` 分支 [README](https://github.com/ethereum/go-ethereum/tree/website#readme) 或 
網站的[貢獻](https://geth.ethereum.org/docs/developers/geth-developer/contributing)頁面。

＃＃ 執照

go-ethereum 庫（即「cmd」目錄之外的所有程式碼）已獲得許可
[GNU 較寬鬆通用公共授權 v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
也包含在我們的儲存庫中的「COPYING.LESSER」檔案中。

go-ethereum 二進位檔案（即「cmd」目錄內的所有程式碼）均已獲得許可
[GNU 通用公共授權 v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)，還有
包含在我們儲存庫的“COPYING”檔案中。
