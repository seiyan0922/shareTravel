{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">イベント詳細</div>
        <div>
            <div>イベント名：{{.Name}}</div>
            <div>日付：{{.Date}}</div>
            <div>このイベントの認証IDは、{{.AuthKey}}です</div>
        </div>
        <div>
            <a href="/member/add?event_id={{.Id}}">メンバーの追加</a>
        </div>
    </div>
</div>

{{ template "footer"}}