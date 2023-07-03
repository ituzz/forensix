package cretutil

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

type OutHtml struct {
}

func NewOutHtml() Oututil {
	oututil := OutHtml{}
	return &oututil
}

func (o *OutHtml) out(hfs []HashFileStruct) error {
	tmpl, _ := template.New("webpage").Parse(htmlTpl)
	htmlFile, err := os.OpenFile(fmt.Sprintf("./%s.html", time.Now().Format("20060102")), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("无法生成HTML文件: %v", err)
		return fmt.Errorf("无法生成HTML文件")
	}
	defer htmlFile.Close()
	writer := bufio.NewWriter(htmlFile)
	tmpl.Execute(writer, struct {
		Data     []HashFileStruct
		HashType string
	}{hfs, hfs[0].HashType})
	writer.Flush()
	return nil
}

const htmlTpl = `

<html>
  <head>
    <meta charset="UTF-8">
    <style type="text/css">
      table {
        margin: 25px;
        min-width: 500px;
        border-collapse: separate;
        border-spacing: 0;
      }

      th,
      td {
        padding: 6px 15px;
      }

      th {
        background: #42444e;
        color: #fff;
        text-align: left;
      }

      tr:first-child th:first-child {
        border-top-left-radius: 6px;
      }

      tr:first-child th:last-child {
        border-top-right-radius: 6px;
      }

      td {
        border-right: 1px solid #c6c9cc;
        border-bottom: 1px solid #c6c9cc;
      }

      td:first-child {
        border-left: 1px solid #c6c9cc;
      }

      tr:nth-child(even) td {
        background: #eaeaed;
      }

      tr:last-child td:first-child {
        border-bottom-left-radius: 6px;
      }

      tr:last-child td:last-child {
        border-bottom-right-radius: 6px;
      }

      a {
        text-decoration: none;
      }
    </style>
  </head>

  <body>
    <table>
      <tr>
        <th>路径</th>
        <th>{{.HashType}}</th>
      </tr>
    {{range .Data}}
        <tr>
            <td><a href="{{.FilePath}}">{{.FilePath}}</a></td>
            <td>{{.HashStr}}</td>
        </tr>
    {{end}}
    </table>
  </body>
</html>
`

func init() {
	OutRegister("HTML", NewOutHtml)
}
