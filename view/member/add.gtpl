{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="side-bar">
            <div class="icon-common">
                <a href="/member/add?event_id={{.Id}}" class="member-add">
                    <div class="member-icon"></div>
                    <div class="add-member-text">メンバー追加</div>
                </a>
            </div>
            <a href="/expense/add?event_id={{.Id}}" class="member-add">
                <div class="icon-common">
                    <div class="expense-icon"></div>
                    <div class="expense-text">会計追加</div>
                </div>
            </a>
            <div class="icon-common">
                <a href="" class="member-add">
                    <div class="download-icon"></div>
                    <div class="download-text">ダウンロード</div>
                </a>
            </div>
        </div>
        <div class="new-event-page">
            <div class="new-event-title">New Member</div>
            <div class="gray-line"></div>
                <div>
                    <div class="new-event-form-box">
                        <form action="/member/save?event_id={{.Id}}" method="POST">
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