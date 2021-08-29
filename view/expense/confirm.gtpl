{{ template "header"}}
<div class="body">
    <div class="content">
        <div class="side-bar">
            <a href="/event/show?event_id={{.Event.Id}}" class="member-add">
                <div class="icon-common">
                    <div class="jet-icon"></div>
                    <div class="add-member-text">イベントTOP</div>
                </div>
            </a>
            <a href="/member/add?event_id={{.Event.Id}}" class="member-add">
                <div class="icon-common">
                    <div class="member-icon"></div>
                    <div class="add-member-text">メンバー追加</div>
                </div>
            </a>
            <a href="/expense/add?event_id={{.Event.Id}}" class="member-add">
                <div class="icon-common">
                    <div class="expense-icon"></div>
                    <div class="expense-text">会計追加</div>
                </div>
            </a>
            <div class="icon-common">
                <a href="/event/download?event_id={{.Event.Id}}" class="member-add">
                    <div class="download-icon"></div>
                    <div class="download-text">ダウンロード</div>
                </a>
            </div>
        </div>
        <div class="new-event-page">
            <div class="new-event-title">Confirm Expense</div>
            <div class="gray-line"></div>
            {{range $err := .Errors}}
            <div class="error" style="color:red; text-align: center;">{{$err}}</div>
            {{end}}
            <div>
                <div class="new-event-form-box">
                    <table class="form-table">
                        <tr>
                            <td>名前</td>
                            <td>：{{.Expense.Name}}</td>
                        </tr>
                        <tr>
                            <td>合計金額</td>
                            <td>：{{.Expense.Total}}円</td> 
                        </tr>
                        <tr>
                            <td>備考</td>
                            <td>：{{.Expense.Remarks}}</td>
                        </tr>
                        <tr>
                            <td>端数</td>
                            <td>：{{.Expense.Pool}}円</td>
                        </tr>
                    </table>
                    <div class="expense-info">※金額は各参加者に等分されます。</div>
                    <div class="expense-info"></div>
                    <form action="/expense/complete?event_id={{.Event.Id}}" method="POST">
                        <table class="form-table">
                            {{range $member := .Members}}
                            <tr>
                                <td>{{$member.Name}}</td>
                                <td>：<input type="text" name="{{$member.Id}}" value="{{.Calculate}}">円</td>
                            </tr>
                            {{end}}
                        </table>
                        <div class="kirisute">
                            <input type="checkbox" name="slash" value="true" checked="checked">
                            <label for="slash">100円以下切捨て</label>
                        </div>
                        <input type="hidden" name="name" value="{{.Expense.Name}}">
                        <input type="hidden" name="total" value="{{.Expense.Total}}">
                        <input type="hidden" name="pool" value="{{.Expense.Pool}}">
                        <input type="hidden" name="remarks" value="{{.Expense.Remarks}}">
                        <input type="hidden" name="price" value="{{.Price}}">
                        <div class="saikeisan">
                            <button type="submit" formaction="/expense/calculate?event_id={{.Event.Id}}">再計算</button>
                        </div>
                        <div class="temporarily-box">
                            <div>一時負担者（立替者）</div>
                            <div class="temporarily-labels">
                                {{$t := .Temporarily}}
                                {{range $member := .Members}}
                                    {{if eq $t $member.Id}}
                                        <input type="radio" id="temporarily{{.Id}}" name="temporarily" value="{{.Id}}" checked="checked">
                                    {{else}}
                                        <input type="radio" id="temporarily{{.Id}}" name="temporarily" value="{{.Id}}">
                                    {{end}}
                                        <label for="temporarily{{.Id}}">{{.Name}}</label></br>
                                {{end}}
                            </div>
                        </div>
                        <div class="common-submit-box">
                            <input class="submit-common" type="submit" value="確定">
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

{{ template "footer"}}