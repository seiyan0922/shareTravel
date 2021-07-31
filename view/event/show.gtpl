{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">イベント詳細</div>
        <div>
            <div>イベント名：{{.Name}}</div>
            <div>日付：{{.Date}}</div>
            <div>このイベントの認証IDは、{{.AuthKey}}です</div>
        </div>
    </div>
</div>

{{ template "footer"}}