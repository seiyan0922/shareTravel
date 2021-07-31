{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">新規イベント登録完了</div>
        <div>
            <div>イベント名：{{.Name}}</div>
            <div>日付：{{.Date}}</div>
            <div>このイベントの認証IDは、{{.AuthKey}}です</div>
            <a href="/event/show?auth_key={{.AuthKey}}">イベント詳細ページ</a>
        </div>
    </div>
</div>

{{ template "footer"}}