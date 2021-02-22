// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dbconnection

// Represents the information stored at dbinfo.json
type DbInfo struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}
