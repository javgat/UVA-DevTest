// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

package dbconnection

import(
  "testing"
)

// Testing getDbInfo:

func TestGetDbInfoMissingFile(t *testing.T){
  dbinfo, err := getDbInfo("test/missing_file.json")
  if err == nil{
    t.Log("error should not be nil")
    t.Fail()
  }else if dbinfo != nil{
    t.Log("dbinfo should be nil")
    t.Fail()
  }
}

func TestGetDbInfoEmptyString(t *testing.T){
  dbinfo, err := getDbInfo("")
  if err == nil{
    t.Log("error should not be nil")
    t.Fail()
  }else if dbinfo != nil{
    t.Log("dbinfo should be nil")
    t.Fail()
  }
}

func TestGetDbInfoEmptyFile(t *testing.T){
  dbinfo, err := getDbInfo("test/empty.json")
  if err == nil{
    t.Log("error should not be nil")
    t.Fail()
  }else if dbinfo != nil{
    t.Log("dbinfo should be nil")
    t.Fail()
  }
}

func TestGetDbInfoMissingElements(t *testing.T){
  dbinfo, err := getDbInfo("test/missing.json")
  if err != nil{
    t.Log("error should be nil", err)
    t.Fail()
  }else if dbinfo == nil{
    t.Log("dbinfo should not be nil")
    t.Fail()
  }
}

func TestGetDbInfoExtraElements(t *testing.T){
  dbinfo, err := getDbInfo("test/extra.json")
  if err != nil{
    t.Log("error should be nil", err)
    t.Fail()
  }else if dbinfo == nil{
    t.Log("dbinfo should not be nil")
    t.Fail()
  }
}

func TestGetDbInfoExactElements(t *testing.T){
  dbinfo, err := getDbInfo("test/correct.json")
  if err != nil{
    t.Log("error should be nil", err)
    t.Fail()
  }else if dbinfo == nil{
    t.Log("dbinfo should not be nil")
    t.Fail()
  }
}
