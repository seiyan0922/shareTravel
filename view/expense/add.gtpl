{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="side-bar">
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
            <div class="new-event-title">New Expense</div>
            <div class="gray-line"></div>
                <div>
                    <div class="new-event-form-box">
                        <form action="/expense/complete?event_id={{.Id}}" method="POST">
                            <table class="form-table">
                                <tr>
                                    <td><labal for="name">名前</labal></td>
                                    <td>：<input type="text" name="name" class="form-common"></td>
                                </tr>
                                <tr>
                                    <td><labal for="price">合計金額</labal></td>
                                    <td>：<input type="text" name="price" class="form-common"></td> 
                                </tr>
                                <tr>
                                    <td><labal for="remarks">備考</labal></td>
                                    <td>：<input type="text" name="remarks" class="form-common"></td>
                                </tr>
                            </table>
                            <div class="expense-info">※金額は各参加者に等分されます。</div>
                            <div class="common-submit-box">
                                <input class="submit-common" type="submit" value="確認">
                            </div>         
                        </form>
                </div>
    </div>
</div>

{{ template "footer"}}