package main

import "html/template"

type adminPageData struct {
	Registrations []registration
	Counts        map[string]int
	Names         map[string]string
}

var adminTmpl = template.Must(template.New("admin").Funcs(template.FuncMap{
	"disciplineName": func(d string) string { return allowedDisciplines[d] },
}).Parse(`<!doctype html>
<html lang="ru">
<head>
<meta charset="utf-8">
<title>Регистрации — WCSF X Academy</title>
<style>
	body { font-family: system-ui, sans-serif; margin: 32px; color: #1a1a1a; }
	h1 { font-size: 20px; }
	.summary { display: flex; gap: 16px; margin: 16px 0 24px; }
	.summary div { background: #f2f2f2; border-radius: 8px; padding: 10px 16px; }
	table { border-collapse: collapse; width: 100%; }
	th, td { text-align: left; padding: 8px 12px; border-bottom: 1px solid #e2e2e2; font-size: 14px; }
	th { background: #fafafa; }
	a.export { display: inline-block; margin-bottom: 16px; }
</style>
</head>
<body>
	<h1>Регистрации на соревнования</h1>
	<div class="summary">
		<div>Всего: {{len .Registrations}}</div>
		{{range $key, $name := .Names}}<div>{{$name}}: {{index $.Counts $key}}</div>{{end}}
	</div>
	<a class="export" href="/admin/export.csv">Скачать CSV</a>
	<table>
		<thead>
			<tr><th>ID</th><th>Имя</th><th>Телефон</th><th>Дисциплина</th><th>Дата</th></tr>
		</thead>
		<tbody>
			{{range .Registrations}}
			<tr>
				<td>{{.ID}}</td>
				<td>{{.FullName}}</td>
				<td>{{.Phone}}</td>
				<td>{{disciplineName .Discipline}}</td>
				<td>{{.CreatedAt.Format "02.01.2006 15:04"}}</td>
			</tr>
			{{else}}
			<tr><td colspan="5">Пока никто не зарегистрировался.</td></tr>
			{{end}}
		</tbody>
	</table>
</body>
</html>
`))
