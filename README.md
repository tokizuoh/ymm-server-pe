# ymm-server-pe
  
[【新卒】サーバーサイドエンジニア応募者向けの模試 | ゆめみ](https://www.yumemi.co.jp/serverside_recruit) > コーディング試験の例 の回答コード。  
  
```bash
> docker-compose exec app go run main.go game_score_log.csv
rank,player_id,mean_score
1,8,7325
2,11,7290
3,10,6180
4,6,6000
5,14,5775
6,3,5728
7,12,5266
8,2,4950
9,4,4071
10,1,4054
```
## Docker
  
### Version

```bash
> docker --version
Docker version 19.03.12, build 48a66213fe

> docker-compose --version
docker-compose version 1.27.2, build 18f557f9
```
  
### Build
  
```bash
> docker-compose up --build -d
```
  
### Run
  
```bash
> docker-compose exec app go run main.go
```
  
---
  
## 振り返り
  
### 2021/08/05 (ブランチ名: [feature/2021_08_05](https://github.com/tokizuoh/ymm-server-pe/tree/feature/2021_08_05))
  
#### 評価観点チェックリスト
  
|  項目  | 自己評価 | 評価詳細 |
| ------------------------------------------------------------------ | ---- | ---- |
| 一度書いたコードは本人の手を離れる                                       |  ❌   | 本人の手を離れる = 他の人が読みやすいコードを書く必要があると解釈。 言語の命名規則の観点から命名が難しかった。改善の余地あり。 |
| 中長期に渡って使い続けられる可能性がある                                  |  ❌   | 中長期に渡って使い続けられる可能性がある = 変更しやすいコードを書く必要があると解釈。現在のコードは `main.go` の main関数内に処理が集約化されており、変更しにくいコードとなっているのでダメ。 |
| コードのメンテナンス時や、障害発生時に常に仕様書とセットで入手可能とは限らない  |  ❌   | コード内でわかりにくい箇所は適宜コメントを入れて可読性を上げる必要があると解釈。現在関数の処理などのコメントを付与できていないのでダメ。 |
| 要求仕様の一部が今後変更される可能性はある                                |  ❌   | 「中長期に渡って使い続けられる可能性がある」と同じ。 |
| 通常、プログラムは単体では利用されない                                   |  ❌   |　エラーのみmainとは別のモジュールを切っているが、ロジック部分を別モジュールに切り出せるので改善の余地あり。ロジックを別モジュールで切り出せば、main以外からも利用できる。 |
| 入力されるデータも他のプログラムにより生成される                           |  ❌   |　今の自分では項目の意味が理解できなかった。他のプログラムによって生成されるから、生成されるプログラムが変わる可能性がある→その変更に柔軟に対応できるようなコードを書く必要がある、ということなのだろうか。|
| データ形式が将来的に変わる可能性がある                                   |  ❌   |　「入力されるデータも他のプログラムにより生成される」と同じ。 |
| - | - | - |
| テストが書かれているか                   | ❌ | 本課題は明確にユニットテストが書ける。現在書けていないので改善の余地あり。 |
| 処理対象のCSVファイルは仕様に則っているか？ | ❌ | 「player_idの構成要素はアルファベットの大文字、小文字、および数字の0-9のみとなります。」現在この仕様に沿っていない。 |
| 処理パフォーマンスの検証はしているか？      | ❌ | 「対象のプレイログ全体は数千万行以上に肥大化することがあります」現在この仕様の検証が出来ていない。肥大化しても処理が耐えうるか検証できる環境を整備する必要がある。 |
| 例外処理は適切か？                       | ❌ | 処理中に例外が発生して、処理を終了させる時にユーザーに何が原因で終了したかを通知できる仕組みになっていない。 |
