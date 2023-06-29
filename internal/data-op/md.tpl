{{ define "PlatformTemplate" }}
### {{.Plat.Name}}
{{.Plat.Id}}  
{{.Plat.Name}}  
{{.Plat.Site}}  
{{.Plat.Remark}}  
{{.Plat.CreateTime}}  
  
{{range $index,$item := .Accounts}}
#### {{$item.Username}}
{{$item.Id}} / {{$item.PlatformId}}  
{{$item.Username}} / {{$item.Password}}  
{{$item.Phone}} / {{$item.Email}}  
{{$item.Remark}}  
{{$item.CreateTime}}  
{{end}}
{{end}}