{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">メンバー追加</div>
        <div>
            <form action="/expense/complete?event_id={{.Id}}" method="POST">
                <labal for="name">名前：</labal>
                <input type="text" name="name">
                <labal for="price">合計金額：</labal>
                <input type="text" name="price">
                <labal for="remarks">備考：</labal>
                <input type="text" name="remarks">
                <div>※金額は各参加者に等分されます。</div>
                <input type="submit" value="追加">
            </form>
        </div>
    </div>
</div>

{{ template "footer"}}