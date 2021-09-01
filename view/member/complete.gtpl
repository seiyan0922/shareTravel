
{{ template "header"}}
<div class="body">
    <div class="content">
        {{template "sidebar" .}}
        <div class="new-event-page">
            <div class="new-event-title">Complete Create Member</div>
            <div class="gray-line"></div>
            <div class="confirm-message">参加者の新規作成が完了しました。</div>
            <div class="confirm-common">
                <table class="form-table">
                    <tr>
                        <td>参加者名</td>
                        <td>：{{.Member.Name}}</td>
                    </tr>
                </table>
                <div class="addmember-to-top">
                    <div class="back-common-box">
                        <a href="/event/show?event_id={{.Event.Id}}" class="back-common">イベントTOPへ→</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}