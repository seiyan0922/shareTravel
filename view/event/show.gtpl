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
                    <li class="tab-item under-line">会計履歴</li>
                    <a class="tab-link" href="/event/indexMember?event_id={{.Event.Id}}">
                        <li class="tab-item">参加者一覧</li>
                    </a>
                    <a class="tab-link" href="/event/edit?event_id={{.Event.Id}}">
                        <li class="tab-item">設定</li>
                    </a>
                </ul>
            </div>
        </div>
        <div class="event-main">
            <div>
                {{if .Expenses}}
                    <table class="event-top-table">
                        <tr class="event-top-table-top">
                            <th>Name</th><th>Total</th><th>DateTime</th><th>Remarks</th><th></th>
                        </tr>
                            {{range $expense := .Expenses}}
                                <tr class="event-top-table-content">
                                    <td class="item-name">{{$expense.Name}}</td>
                                    <td class="item-name-total">{{$expense.Total}}円</td>
                                    <td class="item-time">{{$expense.CreateTime}}</td>
                                    <td class="item-remarks">{{$expense.Remarks}}</td>
                                    <td class="item-link">

                                        <a class="link" href="/expense/edit?expense_id={{$expense.Id}}">
                                            <div class="botton">編集する
                                            </div>
                                        </a>
                                    </td>
                                </tr>

                            {{end}}
                        
                    </table>
                {{else}}
                    <div>会計データがありません</div>
                {{end}}
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}