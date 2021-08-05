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
  