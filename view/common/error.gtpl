{{ template "header"}}
{{range $err := .Errors}}
    <div class="error" style="color:red">{{$err}}</div>
{{end}}
<div>
    {{if eq .Event.Id 0}}
    <a href="/">TOPへ</a>
    {{else}}
    <a href="/event/show?event_id={{.Event.Id}}">イベントTOPへ</a>
    {{end}}
</div>
{{ template "footer"}}