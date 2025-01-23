## BearNetworkChain-Node

BearNetworkChain-Node 是熊網鏈（BRNKC）生態系統中的一個關鍵組件，旨在提供區塊鏈網路的去中心化運行環境。每個節點都是網路的一部分，負責區塊的生成、交易驗證與同步。BearNetworkChain-Node 基於 GETH (Go-Ethereum)，並使用原生幣作為交易媒介，符合 ERC-20 標準，確保了高效、低成本的交易處理。

這些節點透過 Docker 進行部署，提供穩定且高效的運行環境。創世節點負責產生區塊並提供對外 RPC 支援，而其他同步節點則作為創世節點的代理，負責與創世節點同步區塊資訊並提供外部交易查詢服務。節點之間的同步機制保證了網路的一致性和可擴展性，並支持大量交易量的承載，確保在高交易需求下仍能保持極低的 Gas 費用。

BearNetworkChain-Node 的設計專注於提升區塊鏈運行的速度與安全性，並且能夠自動調整節點的運行狀態，適應網路負載的變化。這使得 BearNetworkChain 成為一個穩定且具彈性的區塊鏈平台，適用於各種分散式應用（DApp）和智能合約的運行。

整體而言，BearNetworkChain-Node 的目標是提供一個高效能、低成本且易於管理的區塊鏈節點解決方案，為開發者和使用者創建可靠且可擴展的區塊鏈基礎設施。

#### 熊網鏈唯一推薦使用Docker進行佈署節點

在您的Ubuntu上啟動並運行以太坊的最快方法之一是使用Docker:

1.先創建一個資料夾，依後續指令會用到的資料夾示範為例 :

```shell
創建名稱為: backup-node 的資料夾。
```

2.進到backup-node的資料夾下載官方快速安裝指令文件 setup-node.sh : 

```shell
sudo wget -q https://raw.githubusercontent.com/BearNetwork-BRNKC/genesis/main/setup-node.sh -O
```

3.在/home所在位置執行指令:
```shell
sudo chmod +x setup-node.sh && sudo ./setup-node.sh
```

4.拉取映像檔並且創建容器，其中 -v /home/brnkc/backup-node:/node 的部份要改成你現在backup-node資料夾正確的路徑。(本機資料夾路徑:Docker路徑，這是本機與Docker資料夾映射關係)
```shell
sudo docker run -d -it --restart unless-stopped --name backup-node --network brnkc --ip 172.20.0.5 -v /home/brnkc/backup-node:/node -p 8545:8545 -p 30303:30303 -p 55555:55555 --entrypoint /bin/sh bearnetworkchain/brnkc-node:v1.13.15
```

5.佈署熊網鏈創世文件 :
```shell
sudo docker exec -it backup-node /bin/sh -c "cd /node && geth --datadir brnkc01 init genesis.json"
```

6.啟動節點 :
```shell
sudo docker exec -it backup-node /bin/sh -c "cd /node && geth --datadir brnkc01 init genesis.json && geth --config config.toml --identity \"bearnetwork\" --datadir brnkc01 --http --http.addr 172.20.0.5 --port 30303 --http.corsdomain \"*\" --http.port 8545 --networkid 641230 --nat any --http.api debug,web3,eth,txpool,personal,clique,miner,net --ws --ws.port 55555 --ws.addr 172.20.0.5 --ws.origins \"*\" --ws.api web3,eth --syncmode full --gcmode=archive --nodiscover --http.vhosts=\"*\" --allow-insecure-unlock console"
```

7. 完成。

整個過程都是複製貼上就完成，唯一要調整的部份就是你的主機用戶名稱會所有不同(/home/用戶名稱/backup-node)，因此只要在第四步那個部份依照你的路徑修改一下 -v 路逕指令內容，後續步驟都是複製貼上就可以。



### 熊網鏈節點setup-node.sh內容 (此內容是公開展示的，與下載的setup-node.sh內容相同)
```shell
#!/bin/sh

# 2.. 下載 genesis.json 和 config.toml
echo "下載 genesis.json 和 config.toml..."
wget -q https://raw.githubusercontent.com/BearNetwork-BRNKC/genesis/main/genesis.json
wget -q https://raw.githubusercontent.com/BearNetwork-BRNKC/genesis/main/config.toml

# 2. 設置防火牆端口
echo "設置防火牆端口..."
sudo ufw allow 8545/tcp
sudo ufw allow 30303/tcp
sudo ufw allow 55555/tcp
sudo ufw --force enable

# 3. 創建 Docker 網路（如果已存在則忽略錯誤）
echo "創建 Docker 網路..."
sudo docker network create -d bridge --subnet=172.20.0.0/16 brnkc || true

```

### 熊網鏈節點指令

作為一名開發人員，您遲早會想要開始使用自己的方式與熊網鏈網路透過您自己的方式進行，而不是透過官方配置進行。
`geth` 內建了對基於 JSON-RPC 的 API 的支援（[標準 API](https://ethereum.github.io/execution-apis/api-documentation/)
和 [`geth` 特定 API](https://geth.ethereum.org/docs/interacting-with-geth/rpc))。
這些可以透過 HTTP、WebSockets 和 IPC（基於 UNIX 的 UNIX 套接字）公開平台和 Windows 上的命名管道）。

HTTP 和 WS 介面需要手動啟用，localhost(IP請盡可能使用Docker內網IP，不要隨意使用宿主機或公網IP)

基於 HTTP 的 JSON-RPC API 選項：

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

### 營運熊網鏈專用網絡

維護您自己的熊網鏈專用網路，如有額外需求可以自行進行手動配置啟動設定，配置的指令適用GETH指令。

感謝來自互聯網上的任何人幫助提供熊網鏈節點貢獻。


＃＃ 許可證

熊網鏈庫（即Docker映像之外的所有程式碼）已獲得許可
[GNU 較寬鬆通用公共授權 v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
也包含在我們的儲存庫中的「COPYING.LESSER」檔案中。

熊網鏈二進位檔案（即Docker映像的所有程式碼）均已獲得許可
[GNU 通用公共授權 v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html)，還有
包含在我們儲存庫的“COPYING”檔案中。
