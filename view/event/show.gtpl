{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">イベント詳細</div>
        <div>
            <div>イベント名：{{.Event.Name}}</div>
            <div>日付：{{.Event.Date}}</div>
            <div>このイベントの認証IDは、{{.Event.AuthKey}}です</div>
        </div>
        <div>
            <a href="/member/add?event_id={{.Event.Id}}">メンバーの追加</a>
        </div>
        {{range .Members}}
            <div>name:{{.Name}} </div>
        {{end}}
    </div>
    <div>会計データの追加</div>
    <a href="/expense/add?event_id={{.Event.Id}}">追加する</a>
</div>

{{ template "footer"}}