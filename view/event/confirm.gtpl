{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="new-event-page">
        <div class="new-event-title">New Event</div>
        <div class="gray-line"></div>
        <div class="confirm-message">入力情報を確認し、間違いがなければ登録ボタンを押してください。</div>
        <div class="confirm-common">
            <table class="form-table">
                <tr>
                    <td>イベント名</td>
                    <td>：{{.Event.Name}}</td>
                </tr>
                <tr>
                    <td>日付</td>
                    <td>：{{.Event.Date}}</td>
                </tr>
            </table>
 
            <form action="/event/save" method="POST">
                <input type="hidden" value={{.Event.Date}} name="date">
                <input type="hidden" value={{.Event.Name}} name="name">
                <div class="common-submit-box">
                    <input class="submit-common" type="submit" value="登録">
                </div>  
            </form>
        </div>
        <div  class="back-common-box">
        <a href="/event/create" class="back-common">←back</a>
        </div>
    </div>
    </div>
</div>

{{ template "footer"}}