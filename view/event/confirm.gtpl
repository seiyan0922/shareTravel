{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">新規イベント確認画面</div>
        <div>
            <div>イベント名：{{.Name}}</div>
            <div>日付：{{.Date}}</div>
        </div>
        <form action="/event/save" method="POST">
            <input type="hidden" value={{.Date}} name="date">
            <input type="hidden" value={{.Name}} name="name">
            <input type="submit" value="OK">
            
        </form>
    </div>
</div>

{{ template "footer"}}