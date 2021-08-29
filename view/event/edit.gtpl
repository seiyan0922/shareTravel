{{ template "header"}}
<div class="body">
    <div class="side-bar">
        <div class="icon-common">
            <a href="/member/add?event_id={{.Event.Id}}" class="member-add">
                <div class="member-icon"></div>
                <div class="add-member-text">メンバー追加</div>
            </a>
        </div>
        <div class="icon-common">
            <a href="/expense/add?event_id={{.Event..Id}}" class="member-add">
                <div class="expense-icon"></div>
                <div class="expense-text">会計追加</div>
            </a>
        </div>
        <div class="icon-common">
            <a href="/event/download?event_id={{.Event..Id}}" class="member-add">
                <div class="download-icon"></div>
                <div class="download-text">ダウンロード</div>
            </a>
        </div>
    </div>
    <div class="event-content .clearfix">
        <div class="event-header">
            <div class="event-icon"></div>
            <div class="event-info">
                <div class="event-name">{{.Event.Name}}</div>
                <div class="event-time">{{.Event.Date}}</div>
                <div class="event-key">認証ID：{{.Event.AuthKey}}</div>
            </div>

            <div class="event-top-tab">
                <ul>
                    <a class="tab-link" href="/event/show?event_id={{.Event.Id}}">
                        <li class="tab-item">会計履歴</li>
                    </a>
                    <a class="tab-link" href="/event/indexMember?event_id={{.Event.Id}}">
                        <li class="tab-item">参加者一覧</li>
                    </a>
                    <li class="tab-item under-line">設定</li>
                </ul>
            </div>
        </div>
        <div class="event-main">
            <div>
                <form action="/event/edit?event_id={{.Event.Id}}" method="POST">
                    <table class="">
                        <tr class="">
                            <td><label>認証ID</label></td>
                            <td>：<input type="text" name="auth_key" value="{{.Event.AuthKey}}"></td>
                        </tr>
                        <tr class="">
                            <td class="">イベント名</td>
                            <td>：<input type="text" name="name" value="{{.Event.Name}}"></td>
                        </tr>
                        <tr class="">
                            <td class="">日付</td>
                            <td>：<input type="date" name="date" value="{{.Event.Date}}"></td>
                        </tr>
                    </table>
                    <input type="submit" value="変更">
                </form>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}