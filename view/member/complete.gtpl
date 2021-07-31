
{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">イベント詳細</div>
        <div>
            <div>名前：{{.Name}}</div>
        </div>
        <div>
            <a href="/member/add?event_id={{.EventId}}">メンバーの追加</a>
        </div>
    </div>
</div>

{{ template "footer"}}