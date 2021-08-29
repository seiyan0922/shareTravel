{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="side-bar">
            <a href="/event/show?event_id={{.EventId}}" class="member-add">
                <div class="icon-common">
                    <div class="jet-icon"></div>
                    <div class="add-member-text">イベントTOP</div>
                </div>
            </a>
            <div class="icon-common">
                <a href="/member/add?event_id={{.EventId}}" class="member-add">
                    <div class="member-icon"></div>
                    <div class="add-member-text">メンバー追加</div>
                </a>
            </div>
            <a href="/expense/add?event_id={{.EventId}}" class="member-add">
                <div class="icon-common">
                    <div class="expense-icon"></div>
                    <div class="expense-text">会計追加</div>
                </div>
            </a>
            <div class="icon-common">
                <a href="/event/download?event_id={{.EventId}}" class="member-add">
                    <div class="download-icon"></div>
                    <div class="download-text">ダウンロード</div>
                </a>
            </div>
        </div>
        <div class="new-event-page">
            <div class="new-event-title">Complete Create Expense</div>
            <div class="gray-line"></div>
            <div class="confirm-message">会計の新規作成が完了しました。</div>
            <div class="confirm-common">
                <table class="form-table">
                    <tr>
                        <td>名称</td>
                        <td>：{{.Name}}</td>
                    </tr>
                    <tr>
                        <td>合計金額</td>
                        <td>：{{.Total}}</td>
                    </tr>
                </table>
                <div class="alert-authkey">
                    支払いの登録、各参加者の負担金の登録が完了しました。</br>
                    端数はイベントデータに保存し、イベント終了の際に生産します。
                </div>

                <div class="back-common-box">
                    <a href="/event/show?event_id={{.EventId}}" class="back-common">イベントTOPページへ→</a>
                </div>
               

        </div>
    </div>
</div>

{{ template "footer"}}