package dbconnection

type DbInfo struct{
  Username string `json: "username"`
  Pass string `json: "pass"`
  Host string `json: "host"`
  Port string `json: "port"`
  Name string `json: "name"`
}
