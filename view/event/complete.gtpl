{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="new-event-page">
            <div class="new-event-title">Complete Create Event</div>
            <div class="gray-line"></div>
            <div class="confirm-message">イベントの新規作成が完了しました。</div>
            <div class="confirm-common">
                <table class="form-table">
                    <tr>
                        <td>イベント名</td>
                        <td>：{{.Name}}</td>
                    </tr>
                    <tr>
                        <td>日付</td>
                        <td>：{{.Date}}</td>
                    </tr>
                </table>
                <div class="authkey-confirm">このイベントの認証IDは、
                    <span class="red">{{.AuthKey}}</span>です
                </div>
                <div class="alert-authkey">
                    ※こちらの認証IDは必ず控えて置くようにしてください。
                </div>

                <div class="back-common-box">
                    <a href="/event/show?auth_key={{.AuthKey}}" class="back-common">{{.Name}}の詳細ページへ→</a>
                </div>
               

        </div>
    </div>
</div>

{{ template "footer"}}