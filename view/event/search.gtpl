{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">イベント検索</div>
        <div>
            <form action="/event/search" method="POST">
                <div>
                    <label for="auth_key">イベントID</label>
                    <input type="text" name="auth_key">
                </div>
                <div>
                    <input type="submit">
                </div>
            </form>
        </div>
    </div>
</div>

{{ template "footer"}}