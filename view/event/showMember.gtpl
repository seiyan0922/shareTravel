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
            <a href="/event/download?event_id={{.Event.Id}}" class="member-add">
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
                    <li class="tab-item  under-line">参加者一覧</li>
                    <a class="tab-link" href="/event/edit?event_id={{.Event.Id}}">
                        <li class="tab-item">設定</li>
                    </a>
                </ul>
            </div>
        </div>
        <div class="event-main">
            <div>
                {{if .Members}}
                <table class="event-top-table">
                    <tr class="event-top-table-top">
                        <th>Name</th>
                        <th>Total</th>
                        <th>Temporarily</th>
                    </tr>
                        {{range $member := .Members}}
                            <tr class="event-top-table-content">
                                <td class="human-name">
                                    <div class="icon-common">
                                        <div class="human-icon"></div>{{$member.Name}}
                                    </div>
                                </td>
                                <td class="item-name-total">{{$member.Total}}円</td>
                                <td class="item-name-total">{{$member.Temporarily}}円</td>
                            </tr>

                        {{end}}
                </table>
                {{else}}
                    <div>参加者がいません</div>
                {{end}}
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}