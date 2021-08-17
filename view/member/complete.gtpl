
{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="side-bar">
            <a href="/event/show?event_id={{.Id}}" class="member-add">
                <div class="icon-common">
                    <div class="jet-icon"></div>
                    <div class="add-member-text">イベントTOP</div>
                </div>
            </a>
            <div class="icon-common">
                <a href="" class="member-add">
                    <div class="member-icon"></div>
                    <div class="add-member-text">メンバー追加</div>
                </a>
            </div>
            <div class="icon-common">
                <a href="" class="member-add">
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
        <div class="new-event-page">
            <div class="new-event-title">Complete Create Member</div>
            <div class="gray-line"></div>
            <div class="confirm-message">参加者の新規作成が完了しました。</div>
            <div class="confirm-common">
                <table class="form-table">
                    <tr>
                        <td>参加者名</td>
                        <td>：{{.Name}}</td>
                    </tr>
                </table>
                <div class="addmember-to-top">
                    <div class="back-common-box">
                        <a href="/event/show" class="back-common">イベントTOPへ→</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}