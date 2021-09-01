{{ template "header"}}
<div class="body">
    <div class="content">
        {{template "sidebar" .}}
        <div class="new-event-page">
            
            <div class="new-event-title">New Member</div>
            <div class="gray-line"></div>
            {{range $err := .Errors}}
                <div class="error" style="color:red; text-align: center;">{{$err}}</div>
            {{end}}
                <div>
                    <div class="new-event-form-box">
                        <form action="/member/save?event_id={{.Event.Id}}" method="POST">
                            <table class="form-table">
                                <tr>
                                    <td><labal for="name">名前：</labal></td>
                                    <td><input type="text" name="name"></td>
                                </tr>
                            </table>
                            <div class="expense-info">※参加者の氏名を入力してください。</div>
                            <div class="common-submit-box">
                                <input class="submit-common" type="submit" value="追加">
                            </div>
                            </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}