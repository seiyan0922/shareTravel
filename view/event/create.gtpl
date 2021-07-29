{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="title">新規イベント作成画面</div>
        <div>
            <form action="/event/confirm" method="POST">
                <div>
                    <label for="name">イベント名</label>
                    <input type="text" name="name">
                </div>
                <div>
                    <label for=date>開催日</label>
                    <input type="date" name="datetime">
                </div>
                <div>
                    <input type="submit">
                </div>
            </form>
        </div>
    </div>
</div>

{{ template "footer"}}