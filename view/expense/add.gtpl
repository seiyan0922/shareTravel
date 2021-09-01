{{ template "header"}}
<div class="body">
    <div class="content">
        {{template "sidebar"}}
        <div class="new-event-page">

            <div class="new-event-title">New Expense</div>
            <div class="gray-line"></div>
            {{range $err := .Errors}}
            <div class="error" style="color:red; text-align: center;">{{$err}}</div>
            {{end}}
            <div>
                <div class="new-event-form-box">
                    <form action="/expense/confirm?event_id={{.Event.Id}}" method="POST">
                        <table class="form-table">
                            <tr>
                                <td><labal for="name">名前</labal></td>
                                    {{if eq .Expense nil}}
                                    <td>：<input type="text" name="name" class="form-common"> </td>
                                    {{else}}
                                    <td>：<input type="text" name="name" class="form-common" value="{{.Expense.Name}}"> </td>
                                    {{end}}
                               
                            </tr>
                            <tr>
                                <td><labal for="price">合計金額</labal></td>
                                {{if eq .Expense nil}}
                                <td>：<input type="text" name="total" class="form-common"></td> 
                                {{else}}
                                    {{if eq .Expense.Total 0}}
                                    <td>：<input type="text" name="total" class="form-common"></td> 
                                    {{else}}
                                    <td>：<input type="text" name="total" class="form-common" value="{{.Expense.Total}}"></td> 
                                    {{end}}
                                {{end}}
                            </tr>
                            <tr>
                                <td><labal for="remarks">備考</labal></td>
                                {{if eq .Expense nil}}
                                <td>：<input type="text" name="remarks" class="form-common"></td>
                                {{else}}
                                <td>：<input type="text" name="remarks" class="form-common" value="{{.Expense.Remarks}}"></td>
                                {{end}}
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

    </div>
</div>

{{ template "footer"}}