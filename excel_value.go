package main
import (
    "fmt"
    "github.com/tealeg/xlsx"
    "log"
    "flag"
)

var xlsxPath = flag.String("o", "", "XLSX  file")
var colvalue = flag.Int("c",0 , "Value column number")
var operflag = flag.Int("v",0, "values")


func main() {
    count:=0
    flag.Parse()
    excelFileName := *xlsxPath
    xlFile, err := xlsx.OpenFile(excelFileName)

    if err != nil {
        log.Fatal(err)
    }
    for _, sheet := range xlFile.Sheets {
       for _, row := range sheet.Rows {
                for _, cell := range row.Cells {
                        
                      value,_ := row.Cells[*colvalue].Int()
                      str, _ := cell.String()

                      if value >= *operflag{
                         count=value 
                         fmt.Printf("%s ",str)
                      }
          }
          if count>= *operflag { 
           fmt.Printf("\n")
             count=-1
          }
     }
  }

}

//https://github.com/tealeg/xlsx