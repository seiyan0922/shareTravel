{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">メンバー追加</div>
        <div>
            <form action="/member/save?event_id={{.Id}}" method="POST">
                <labal for="name">名前：</labal>
                <input type="text" name="name">
                <input type="submit" value="追加">
            </form>
        </div>
    </div>
</div>

{{ template "footer"}}