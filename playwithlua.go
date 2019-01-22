package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Datarows struct {
	filename string
	sheetname string
	xlsx *excelize.File
}

const luaDatarowsTypeName = "datarows"

func registerDatarowsType(L *lua.LState) {
    mt := L.NewTypeMetatable(luaDatarowsTypeName)
    L.SetGlobal("datarows", mt)
    // static attributes
    L.SetField(mt, "new", L.NewFunction(newDatarows))
    // methods
    L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), DatarowsMethods))
}

// Constructor
func newDatarows(L *lua.LState) int {
	filename := L.CheckString(1)
	sheetname := L.CheckString(2)
	xlsx := readExcel(filename,sheetname)
    person := &Datarows{filename,sheetname,xlsx}
    ud := L.NewUserData()
    ud.Value = person
    L.SetMetatable(ud, L.GetTypeMetatable(luaDatarowsTypeName))
    L.Push(ud)
    return 1
}

// Checks whether the first lua argument is a *LUserData with *Person and returns this *Datarows.
func checkDatarows(L *lua.LState) *Datarows {
    ud := L.CheckUserData(1)
    if v, ok := ud.Value.(*Datarows); ok {
        return v
    }
    L.ArgError(1, "Datarows expected")
    return nil
}

var DatarowsMethods = map[string]lua.LGFunction{
    "get": datarowsGet,
    "set": datarowsSet,
    "save": datarowsSave,
}

// Getter for the Datarows row, col 
func datarowsGet(L *lua.LState) int {
	p := checkDatarows(L)
	axis := L.CheckString(2)
	value := p.xlsx.GetCellValue(p.sheetname, axis)
    L.Push(lua.LString(value))
    return 1
}

// Setter for the Datarows row, col 
func datarowsSet(L *lua.LState) int {
	p := checkDatarows(L)
	axis := L.CheckString(2)
	value := L.CheckString(3)
	p.xlsx.SetCellValue(p.sheetname,axis,value)
    return 1
}

// Setter for the Datarows row, col 
func datarowsSave(L *lua.LState) int {
	p := checkDatarows(L)
	if L.GetTop() == 2 {
		name := L.CheckString(2)
		p.xlsx.SaveAs(name)
        return 1
    }
	p.xlsx.Save()
    return 1
}

func readExcel(filename , sheetname string) (*excelize.File){
	xlsx, err := excelize.OpenFile(filename)
    if err != nil {
        fmt.Println(err)
        return nil
    }
	return xlsx
}

func main() {
	L := lua.NewState()
	defer L.Close()

	registerDatarowsType(L)	

	if err := L.DoFile(`business.lua`); err != nil {
    	panic(err)
	}
	if err := L.DoString(`computefile("","平台展示")`); err != nil {
    	panic(err)
	}
}
