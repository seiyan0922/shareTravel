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
            <a href="/expense/add?event_id={{.Event.Id}}" class="member-add">
                <div class="expense-icon"></div>
                <div class="expense-text">会計追加</div>
            </a>
        </div>
        <div class="icon-common">
            <a href="" class="member-add">
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
                    <li class="tab-item under-line">会計履歴</li>
                    <a class="tab-link" href="/members/index?event_id={{.Event.Id}}">
                        <li class="tab-item">参加者一覧</li>
                    </a>
                    <a class="tab-link" href="/event/edit?event_id={{.Event.Id}}">
                        <li class="tab-item">設定</li>
                    </a>
                </ul>
            </div>
        </div>
        <div class="event-main">
        </div>
    </div>
</div>

{{ template "footer"}}